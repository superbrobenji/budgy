#!/usr/bin/env node
import * as dotenv from "dotenv";
import { ApiStack } from "../lib/apiStack";
import { DatabaseStack } from "../lib/databaseStack";
import { AuthStack } from "../lib/authStack";
import { App } from "aws-cdk-lib";

dotenv.config();
const app = new App();
if (!process.env.AWS_ACCOUNT || !process.env.AWS_REGION) {
  throw new Error("AWS environment variable is required");
}

const _routeApiEndpoint = "/api";

const env = {
  account: process.env.AWS_ACCOUNT as string,
  region: process.env.AWS_REGION as string,
};

const databaseStack = new DatabaseStack(app, "BudgyDatabaseStack", {
  env,
});

const authStack = new AuthStack(app, "BudgyAuthStack", {
  userpoolConstructName: "BudgyUserpool",
  hasCognitoGroups: false,
  identitypoolConstructName: "BudgyIdentityPool",
  tables: databaseStack.tables,
  env,
});

new ApiStack(app, "BudgyApiStack", {
  env: {
    ...env,
    routeApiEndpoint: _routeApiEndpoint,
  },
  tables: databaseStack.tables,
  userPool: authStack.userpool,
  userPoolClient: authStack.userpoolClient,
});
