import { Stack, StackProps } from 'aws-cdk-lib/core';
import { Construct } from 'constructs';
import * as go from '@aws-cdk/aws-lambda-go-alpha';
import getFileNamesAndDirectories from '../utils/generateLambdaFunctions';
import * as routeDefs from '../../../transport/http/routeDefs.json'
import path = require('path');
import { HttpApi, HttpAuthorizer, HttpAuthorizerType, HttpMethod } from 'aws-cdk-lib/aws-apigatewayv2';
import { HttpLambdaIntegration } from 'aws-cdk-lib/aws-apigatewayv2-integrations';
import { Table } from 'aws-cdk-lib/aws-dynamodb';
import { UserPool } from 'aws-cdk-lib/aws-cognito';
import { Authorizer } from 'aws-cdk-lib/aws-apigateway';
import { HttpUserPoolAuthorizer } from 'aws-cdk-lib/aws-apigatewayv2-authorizers';

interface ApiStackProps extends StackProps {
    env: {account: string, region: string, version: string, routeApiEndpoint: string},
    tables: {[key: string]: Table},
    userPool: UserPool,
}
 
export class ApiStack extends Stack {
    constructor(scope: Construct, id: string, props: ApiStackProps) {
        super(scope, id, props);
        // API Gateway
        const httpApi = new HttpApi(this, 'BudgyHttpApi', {
            apiName: 'BudgyApi',
            description: 'This is the Budgy API',
        });

        // Lambda functions
        const methodFactory = (method: string) => {
            switch (method) {
                case 'GET':
                    return HttpMethod.GET;
                case 'POST':
                    return HttpMethod.POST;
                case 'PUT':
                    return HttpMethod.PUT;
                case 'DELETE':
                    return HttpMethod.DELETE;
                default:
                    return HttpMethod.ANY;
            }
        }
        //lambda functions factory
        const userPoolAuthorizer = new HttpUserPoolAuthorizer('userPoolAuthorizer', props.userPool)

        const baseApiPath = process.env.ROUTE_API_ENDPOINT + "/" + process.env.VERSION;
        const lambdaPath = path.join(__dirname, '..', '..', '..', 'transport', 'http', 'lambdaHandlers');
        const { fileNames, directories } = getFileNamesAndDirectories(lambdaPath);
        for (let i = 0; i < fileNames.length; i++) {
            const lambdaFunc = new go.GoFunction(this, fileNames[i], {
                entry: path.join(directories[i]),
            });
            //TODO create dynamic table access for each lambda function
            props.tables.itemsTable.grantReadWriteData(lambdaFunc);
            props.tables.transactionsTable.grantReadWriteData(lambdaFunc);
            props.tables.categoriesTable.grantReadWriteData(lambdaFunc);
            // API Gateway route factory
            for (const routeDef of routeDefs) {
                const serviceName = path.basename(directories[i]);
                if (routeDef.handler === serviceName) {
                    const route = routeDef.route;
                    const method = methodFactory(routeDef.method)
                    const lambdaInegration = new HttpLambdaIntegration(fileNames[i], lambdaFunc);
                    httpApi.addRoutes({
                        path: baseApiPath + route,
                        methods: [method],
                        integration: lambdaInegration,
                        authorizer: userPoolAuthorizer
                    });
                }
            }
        }
    }
}

