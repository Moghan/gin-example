package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func initializeAwsConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }
	return cfg
}

func initializeDynamoDB(cfg aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(cfg)
}

func getAllFeatures(c *gin.Context) {

	cfg := initializeAwsConfig()

    // Using the Config value, create the DynamoDB client
    svc := initializeDynamoDB(cfg)

    // Build the request with its input parameters
    resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
        Limit: aws.Int32(15),
    })
    if err != nil {
        log.Fatalf("failed to list tables, %v", err)
    }

    fmt.Println("Tables:")
    for _, tableName := range resp.TableNames {
        fmt.Println(tableName)
    }
}

func getFeatureById(c *gin.Context) {
	id := c.Query("id")
	c.String(http.StatusOK, "TODO: getFeatureById, %s", id)
}

func StartServer(c *cli.Context) {
	router := gin.Default()

	v0 := router.Group("/v0")
	
	v0.GET("features", getAllFeatures)
	v0.GET("features/:id", getFeatureById)
	

	router.Run()
}

func Execute() {
	app := &cli.App{
		Name: "boom",
		Usage: "boom boom challange",
		Action: StartServer,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func main() {
	Execute()
}