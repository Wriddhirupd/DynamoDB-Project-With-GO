package main

import (
	// "encoding/json"
	"fmt"
	// "os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	svc := dynamodb.New(sess)

	type MovieInfo struct {
		Plot   string  `json:"plot"`
		Rating float64 `json:"rating"`
	}

	type Movie struct {
		Year  int       `json:"year"`
		Title string    `json:"title"`
		Info  MovieInfo `json:"info"`
	}

	movie := Movie{
		Year:  2015,
		Title: "The Big New Movie",
		Info: MovieInfo{
			Plot:   "Nothing happens at all.",
			Rating: 0.0,
		},
	}

	av, err := dynamodbattribute.MarshalMap(movie)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Movies"),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("We have inserted a new item!\n")

}
