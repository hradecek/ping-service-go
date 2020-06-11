package dynamodb

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
	"testing"
	"time"
)

// DynamoDB erroneous mock
type mockDynamoDBClientBroken struct {
	dynamodbiface.DynamoDBAPI
}

func (m mockDynamoDBClientBroken) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, errors.New("cannot put item into dynamodb")
}

// Given uptime item has been created,
// When uptime item is stored into DynamoDB
//      and DynamoDB PutItem operation fails,
// Then error is returned.
func TestStoreItemPutItemFailure(t *testing.T) {
	// Given
	requestID, _ := uuid.NewRandom()
	uptimeID, _ := uuid.NewRandom()
	uptimeItem := UptimeItem{
		RequestID:  requestID.String(),
		UptimeID:   uptimeID.String(),
		RunAt:      time.Now().Unix(),
		Host:       "https://www.hopefuly-not-existing-host.com",
		StatusCode: 200,
		TTFB:       100,
	}

	// When
	err := StoreUptime(uptimeItem, "SampleTable", mockDynamoDBClientBroken{})

	// Then
	if err == nil {
		t.Error("Error was expected to be returned")
	}
}
