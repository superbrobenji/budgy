package sdk

import (
	"context"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfigMod "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

var awsConfig aws.Config
var onceAwsConfig sync.Once

var cognitoClient *cognitoidentityprovider.Client
var onceAuthClient sync.Once

func getAwsConfig() aws.Config {
	onceAwsConfig.Do(func() {
		var err error
		awsConfig, err = awsConfigMod.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic(err)
		}
	})

	return awsConfig
}

func GetCognitoClient() *cognitoidentityprovider.Client {
	onceAuthClient.Do(func() {
		awsConfig = getAwsConfig()

		region := os.Getenv("AWS_REGION")

		cognitoClient = cognitoidentityprovider.NewFromConfig(awsConfig, func(opt *cognitoidentityprovider.Options) {
			opt.Region = region
		})
	})

	return cognitoClient
}
