#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import * as dotenv from 'dotenv'
import { ApiStack } from '../lib/apiStack';
import { DatabaseStack } from '../lib/databaseStack';
import { AuthStack } from '../lib/authStack';

dotenv.config()
const app = new cdk.App();
if (!process.env.VERSION) {
    throw new Error('VERSION environment variable is required')
}
const version = process.env.VERSION.charAt(0).toUpperCase() + process.env.VERSION.slice(1)
const dashedVersion = version.replace('.', '-')

const authStack = new AuthStack(app, `BudgyAuthStack${dashedVersion}`, {
    userpoolConstructName: `BudgyUserpool${dashedVersion}`,
    hasCognitoGroups: false,
    identitypoolConstructName: `BudgyIdentityPool${dashedVersion}`,
    env: {
        account: process.env.AWS_ACCOUNT as string,
        region: process.env.AWS_REGION as string,
    },
});

const databaseStack = new DatabaseStack(app, `BudgyDatabaseStack${dashedVersion}`, {
    env: {
        account: process.env.AWS_ACCOUNT as string,
        region: process.env.AWS_REGION as string,
    },
});


new ApiStack(app, `BudgyApiStack${dashedVersion}`, {
    env: {
        account: process.env.AWS_ACCOUNT as string,
        region: process.env.AWS_REGION as string,
        version: process.env.VERSION,
        routeApiEndpoint: process.env.ROUTE_API_ENDPOINT as string
    },
    tables: databaseStack.tables,
    userPool: authStack.userpool,
});
