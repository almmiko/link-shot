package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type URL struct {
	OriginalUrl string `json:"original_url"`
	ShortUrl    string `json:"short_url"`
}

type ValidationError struct {
	Err     error  `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	TABLE = "Links"
)

var dbConn *dynamodb.DynamoDB

// handlers
func CreateShortUrl(w http.ResponseWriter, r *http.Request) {

	url := &URL{}

	err := json.NewDecoder(r.Body).Decode(url)
	if err != nil {
		JSON(w, http.StatusUnprocessableEntity, &ValidationError{
			Err: err,
		})
		return
	}

	if url.OriginalUrl == "" {
		JSON(w, http.StatusUnprocessableEntity, &ValidationError{
			Message: "expected 'original_url'",
		})
		return
	}

	data := &URL{
		OriginalUrl: url.OriginalUrl,
		ShortUrl:    GetShortID(),
	}

	av, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		log.Fatalf("Error marshalling map: %v", err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TABLE),
	}

	_, err = dbConn.PutItem(input)
	if err != nil {
		log.Fatalf("Error calling PutItem: %v", err.Error())

	}
	JSON(w, http.StatusOK, data)
}

func NewDatabaseConnection() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)

	return dynamodb.New(sess), err
}

func init() {
	conn, err := NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Failed database connection %v", err)
	}

	dbConn = conn
}
