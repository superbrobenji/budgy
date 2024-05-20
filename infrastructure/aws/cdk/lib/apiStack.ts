import { Stack, StackProps } from "aws-cdk-lib/core";
import { Construct } from "constructs";
import getFileNamesAndDirectories from "../utils/generateLambdaFunctions";
import * as routeDefs from "../../../transport/http/routeDefs.json";
import path = require("path");
import {
  AddRoutesOptions,
  HttpApi,
  HttpMethod,
} from "aws-cdk-lib/aws-apigatewayv2";
import { HttpLambdaIntegration } from "aws-cdk-lib/aws-apigatewayv2-integrations";
import { TableV2 } from "aws-cdk-lib/aws-dynamodb";
import { UserPool, UserPoolClient } from "aws-cdk-lib/aws-cognito";
import { HttpUserPoolAuthorizer } from "aws-cdk-lib/aws-apigatewayv2-authorizers";
import { GoFunction } from "@aws-cdk/aws-lambda-go-alpha/lib/function";

interface ApiStackProps extends StackProps {
  env: {
    account: string;
    region: string;
    routeApiEndpoint: string;
  };
  tables: { [key: string]: TableV2 };
  userPool: UserPool;
  userPoolClient: UserPoolClient;
}

type routeOptions = {
  path: string;
  methods: string[];
  integration: HttpLambdaIntegration;
  authorizer?: HttpUserPoolAuthorizer;
};

export class ApiStack extends Stack {
  constructor(scope: Construct, id: string, props: ApiStackProps) {
    super(scope, id, props);
    // API Gateway
    const httpApi = new HttpApi(this, "BudgyHttpApi", {
      apiName: "BudgyApi",
      description: "This is the Budgy API",
    });

    // Lambda functions
    const methodFactory = (method: string) => {
      switch (method) {
        case "GET":
          return HttpMethod.GET;
        case "POST":
          return HttpMethod.POST;
        case "PUT":
          return HttpMethod.PUT;
        case "DELETE":
          return HttpMethod.DELETE;
        default:
          return HttpMethod.ANY;
      }
    };
    //lambda functions factory
    const userPoolAuthorizer = new HttpUserPoolAuthorizer(
      "userPoolAuthorizer",
      props.userPool,
      { userPoolClients: [props.userPoolClient] },
    );

    const baseApiPath = props.env.routeApiEndpoint;
    const lambdaPath = path.join(
      __dirname,
      "..",
      "..",
      "..",
      "transport",
      "http",
      "lambdaHandlers",
    );
    const { fileNames, directories } = getFileNamesAndDirectories(lambdaPath);
    for (let i = 0; i < fileNames.length; i++) {
      const lambdaFunc = new GoFunction(this, fileNames[i], {
        entry: path.join(directories[i]),
        functionName: fileNames[i],
      });
      // API Gateway route factory
      for (const routeDef of routeDefs) {
        const serviceName = path.basename(directories[i]);
        if (routeDef.handler === serviceName) {
          if (routeDef.environments) {
            for (const environment of routeDef.environments) {
              const env = process.env[environment];
              if (!env) {
                throw new Error(`env ${environment} does not exist`);
              }
              lambdaFunc.addEnvironment(environment, env);
            }
          }
          if (routeDef.tables) {
            for (const table of routeDef.tables) {
              const dbTable = props.tables[table];
              if (!dbTable) {
                throw new Error(`table ${table} does not exist`);
              }
              dbTable.grantReadWriteData(lambdaFunc);
            }
          }
          const route = routeDef.route;
          const method = methodFactory(routeDef.method);
          const lambdaInegration = new HttpLambdaIntegration(
            fileNames[i],
            lambdaFunc,
          );
          let routeOptions: routeOptions = {
            path: baseApiPath + route,
            methods: [method],
            integration: lambdaInegration,
          };
          if (routeDef.authorizer) {
            routeOptions.authorizer = userPoolAuthorizer;
          }
          httpApi.addRoutes(routeOptions as AddRoutesOptions);
        }
      }
    }
  }
}
