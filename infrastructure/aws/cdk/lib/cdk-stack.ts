import { CfnResource, Stack, StackProps } from 'aws-cdk-lib/core';
import * as dotenv from 'dotenv'
import { Construct } from 'constructs';
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import * as go from '@aws-cdk/aws-lambda-go-alpha';
import getFileNamesAndDirectories from '../utils/generateLambdaFunctions';
import * as routeDefs from '../../../transport/http/routeDefs.json'
import path = require('path');
import { CfnIntegration, CfnRoute, HttpApi, HttpMethod } from 'aws-cdk-lib/aws-apigatewayv2';
import * as ec2 from 'aws-cdk-lib/aws-ec2';
import { HttpLambdaIntegration } from 'aws-cdk-lib/aws-apigatewayv2-integrations';
//import * as ecs from 'aws-cdk-lib/aws-ecs';
//import * as ecs_patterns from 'aws-cdk-lib/aws-ecs-patterns';
//import { DockerImageAsset } from 'aws-cdk-lib/aws-ecr-assets';
//import { join } from 'path';
//import { CfnIntegration, CfnRoute, HttpApi } from 'aws-cdk-lib/aws-apigatewayv2';

export class BudgyStack extends Stack {
    constructor(scope: Construct, id: string, props?: StackProps) {
        super(scope, id, props);
        dotenv.config()

        //DynamoDB tables
        const categoriesTable = new dynamodb.Table(this, 'dynamodbCategoriesStack', {
            partitionKey: { name: 'Id', type: dynamodb.AttributeType.STRING },
            billingMode: dynamodb.BillingMode.PAY_PER_REQUEST, // Use on-demand billing
            tableName: 'categories'
        });

        const itemsTable = new dynamodb.Table(this, 'dynamodbItemsStack', {
            partitionKey: { name: 'Id', type: dynamodb.AttributeType.STRING },
            billingMode: dynamodb.BillingMode.PAY_PER_REQUEST, // Use on-demand billing
            tableName: 'items'
        });
        const transactionsTable = new dynamodb.Table(this, 'dynamodbTransactionsStack', {
            partitionKey: { name: 'Id', type: dynamodb.AttributeType.STRING },
            billingMode: dynamodb.BillingMode.PAY_PER_REQUEST, // Use on-demand billing
            tableName: 'transactions'
        });

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
        const baseApiPath = process.env.ROUTE_API_ENDPOINT + "/" + process.env.VERSION;
        const lambdaPath = path.join(__dirname, '..', '..', '..', 'transport', 'http', 'lambdaHandlers');
        const { fileNames, directories } = getFileNamesAndDirectories(lambdaPath);
        for (let i = 0; i < fileNames.length; i++) {
            const lambdaFunc = new go.GoFunction(this, fileNames[i], {
                entry: path.join(directories[i]),
            });
            //TODO create dynamic table access for each lambda function
            itemsTable.grantReadWriteData(lambdaFunc);
            transactionsTable.grantReadWriteData(lambdaFunc);
            categoriesTable.grantReadWriteData(lambdaFunc);
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
                    });
                }
            }
        }


        //        //ECS Fargate
        //        const image = new DockerImageAsset(this, "BackendImage", {
        //            directory: join(__dirname, "..", "..", "..", ".."),
        //            file: "Dockerfile.multistage",
        //            target: "release-stage",
        //        });
        //
        //        const vpc = new ec2.Vpc(this, "MyVpc", {
        //            maxAzs: 3 // Default is all AZs in region
        //        });
        //
        //        const cluster = new ecs.Cluster(this, "MyCluster", {
        //            vpc: vpc
        //        });
        //
        //        // Create a load-balanced Fargate service and make it public
        //        const fargate = new ecs_patterns.ApplicationLoadBalancedFargateService(this, "MyFargateService", {
        //            cluster: cluster, // Required
        //            taskImageOptions: { image: ecs.ContainerImage.fromDockerImageAsset(image) },
        //        });
        //
        //        const httpVpcLink = new cdk.CfnResource(this, 'HttpVpcLink', {
        //            type: 'AWS::ApiGatewayV2::VpcLink',
        //            properties: {
        //                Name: 'V2 VPC Link',
        //                SubnetIds: vpc.privateSubnets.map(m => m.subnetId)
        //            }
        //        });
        //
        //        const api = new HttpApi(this, 'HttpApiGateway', {
        //            apiName: 'ApigwFargate',
        //            description: 'Integration between apigw and Application Load-Balanced Fargate Service',
        //        });
        //
        //        const integration = new CfnIntegration(this, 'HttpApiGatewayIntegration', {
        //            apiId: api.httpApiId,
        //            connectionId: httpVpcLink.ref,
        //            connectionType: 'VPC_LINK',
        //            description: 'API Integration with AWS Fargate Service',
        //            integrationMethod: 'GET', // for GET and POST, use ANY
        //            integrationType: 'HTTP_PROXY',
        //            integrationUri: fargate.listener.listenerArn,
        //            payloadFormatVersion: '1.0', // supported values for Lambda proxy integrations are 1.0 and 2.0. For all other integrations, 1.0 is the only supported value
        //        });
        //
        //        new CfnRoute(this, 'Route', {
        //            apiId: api.httpApiId,
        //            routeKey: 'GET /',  // for something more general use 'ANY /{proxy+}'
        //            target: `integrations/${integration.ref}`,
        //        })
        //
        //        new cdk.CfnOutput(this, 'APIGatewayUrl', {
        //            description: 'API Gateway URL to access the GET endpoint',
        //            value: api.url!
        //        })
    }
}

