import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
// import * as sqs from 'aws-cdk-lib/aws-sqs';
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';

export class BudgyStack extends cdk.Stack {
    public readonly table: dynamodb.Table;
    constructor(scope: Construct, id: string, props?: cdk.StackProps) {
        super(scope, id, props);

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
        // example resource
        // const queue = new sqs.Queue(this, 'CdkQueue', {
        //   visibilityTimeout: cdk.Duration.seconds(300)
        // });
    }
}
