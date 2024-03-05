import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import * as ec2 from 'aws-cdk-lib/aws-ec2';
import * as ecs from 'aws-cdk-lib/aws-ecs';
import * as ecs_patterns from 'aws-cdk-lib/aws-ecs-patterns';
import { DockerImageAsset } from 'aws-cdk-lib/aws-ecr-assets';
import { join } from 'path';

export class BudgyStack extends cdk.Stack {
    public readonly table: dynamodb.Table;
    constructor(scope: Construct, id: string, props?: cdk.StackProps) {
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

        //ECS Fargate
        const image = new DockerImageAsset(this, "BackendImage", {
            directory: join(__dirname, "..", "..", "..", ".."),
            file: "Dockerfile.multistage",
            target: "release-stage"
        });
        const vpc = new ec2.Vpc(this, "MyVpc", {
            maxAzs: 3 // Default is all AZs in region
        });

        const cluster = new ecs.Cluster(this, "MyCluster", {
            vpc: vpc
        });

        // Create a load-balanced Fargate service and make it public
        new ecs_patterns.ApplicationLoadBalancedFargateService(this, "MyFargateService", {
            cluster: cluster, // Required
            cpu: 512, // Default is 256
            desiredCount: 1, // Default is 1
            taskImageOptions: { image: ecs.ContainerImage.fromDockerImageAsset(image) },
            memoryLimitMiB: 512, // Default is 512
            publicLoadBalancer: true // Default is true
        });
    }
}