//go:build wireinject
// +build wireinject

package main

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/wire"
)

func InitializeService() *dynamodb.Client {
    wire.Build(
		initializeAwsConfig,
		initializeDynamoDB,
	)
    return &dynamodb.Client{}
}