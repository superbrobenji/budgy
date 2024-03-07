import { Stack, StackProps } from 'aws-cdk-lib/core';
import { Construct } from 'constructs';
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import * as go from '@aws-cdk/aws-lambda-go-alpha';
import getFileNamesAndDirectories from '../utils/generateLambdaFunctions';
import path = require('path');
//import * as ec2 from 'aws-cdk-lib/aws-ec2';
//import * as ecs from 'aws-cdk-lib/aws-ecs';
//import * as ecs_patterns from 'aws-cdk-lib/aws-ecs-patterns';
//import { DockerImageAsset } from 'aws-cdk-lib/aws-ecr-assets';
//import { join } from 'path';
//import { CfnIntegration, CfnRoute, HttpApi } from 'aws-cdk-lib/aws-apigatewayv2';

export class BudgyStack extends Stack {
    public readonly table: dynamodb.Table;
    constructor(scope: Construct, id: string, props?: StackProps) {
        super(scope, id, props);

        //DynamoDB tables
        this.table = new dynamodb.Table(this, 'dynamodbCategoriesStack', {
            partitionKey: { name: 'Id', type: dynamodb.AttributeType.STRING },
            billingMode: dynamodb.BillingMode.PAY_PER_REQUEST, // Use on-demand billing
            tableName: 'categories'
        });

        this.table = new dynamodb.Table(this, 'dynamodbItemsStack', {
            partitionKey: { name: 'Id', type: dynamodb.AttributeType.STRING },
            billingMode: dynamodb.BillingMode.PAY_PER_REQUEST, // Use on-demand billing
            tableName: 'items'
        });
        this.table = new dynamodb.Table(this, 'dynamodbTransactionsStack', {
            partitionKey: { name: 'Id', type: dynamodb.AttributeType.STRING },
            billingMode: dynamodb.BillingMode.PAY_PER_REQUEST, // Use on-demand billing
            tableName: 'transactions'
        });

        const lambdaPath = path.join(__dirname, '..', '..', '..', '..', 'core', 'service', 'lambda');
        const { fileNames, directories } = getFileNamesAndDirectories(lambdaPath);
        for (let i = 0; i < fileNames.length; i++) {
            const lambdaFunc = new go.GoFunction(this, fileNames[i], {
                entry: path.join(directories[i]),
                bundling: {
                    environment: {
                        TABLE_NAME: this.table.tableName
                    }
                }
            });
            this.table.grantReadWriteData(lambdaFunc);
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

