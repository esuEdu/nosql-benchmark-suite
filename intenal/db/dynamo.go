package db

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBBench struct {
	Client    *dynamodb.Client
	TableName string
}

func NewDynamo(tableName string) (*DynamoDBBench, error) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				if service == dynamodb.ServiceID {
					return aws.Endpoint{
						URL: "http://localhost:8000", // DynamoDB Local 🚀
					}, nil
				}
				return aws.Endpoint{}, fmt.Errorf("unknown endpoint: %s", service)
			}),
		),
	)
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBBench{Client: client, TableName: tableName}, nil
}

func (d *DynamoDBBench) Name() string { return "dynamo" }

func (d *DynamoDBBench) WriteTest(n int) (time.Duration, error) {
	ctx := context.Background()
	start := time.Now()
	for i := 0; i < n; i++ {
		_, err := d.Client.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: &d.TableName,
			Item: map[string]types.AttributeValue{
				"id":    &types.AttributeValueMemberS{Value: strconv.Itoa(i)},
				"value": &types.AttributeValueMemberS{Value: fmt.Sprintf("value-%d", i)},
				"ts":    &types.AttributeValueMemberS{Value: strconv.FormatInt(time.Now().UnixNano(), 10)},
			},
		})
		if err != nil {
			return 0, err
		}
	}
	return time.Since(start), nil
}

func (d *DynamoDBBench) ReadTest(n int) (time.Duration, error) {
	ctx := context.Background()
	start := time.Now()
	for i := 0; i < n; i++ {
		id := strconv.Itoa(i)
		_, _ = d.Client.GetItem(ctx, &dynamodb.GetItemInput{
			TableName: &d.TableName,
			Key: map[string]types.AttributeValue{
				"id": &types.AttributeValueMemberS{Value: id},
			},
		})
		// ignoring content, measuring latency
	}
	return time.Since(start), nil
}
