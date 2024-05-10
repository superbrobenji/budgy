#!/usr/bin/env node
import * as dotenv from "dotenv";
import { ApiStack } from "../lib/apiStack";
import { DatabaseStack } from "../lib/databaseStack";
import { AuthStack } from "../lib/authStack";
import { App } from "aws-cdk-lib";

dotenv.config();
const app = new App();
if (!process.env.VERSION) {
  throw new Error("VERSION environment variable is required");
}
const version =
  process.env.VERSION.charAt(0).toUpperCase() + process.env.VERSION.slice(1);
const dashedVersion = version.replace(".", "-");
const env = {
  account: process.env.AWS_ACCOUNT as string,
  region: process.env.AWS_REGION as string,
};

const databaseStack = new DatabaseStack(
  app,
  `BudgyDatabaseStack${dashedVersion}`,
  {
    env,
  },
);

const authStack = new AuthStack(app, `BudgyAuthStack${dashedVersion}`, {
  userpoolConstructName: `BudgyUserpool${dashedVersion}`,
  hasCognitoGroups: false,
  identitypoolConstructName: `BudgyIdentityPool${dashedVersion}`,
  tables: databaseStack.tables,
  env,
});

new ApiStack(app, `BudgyApiStack${dashedVersion}`, {
  env: {
    ...env,
    version: process.env.VERSION,
    routeApiEndpoint: process.env.ROUTE_API_ENDPOINT as string,
  },
  tables: databaseStack.tables,
  userPool: authStack.userpool,
  userPoolClient: authStack.userpoolClient,
});
