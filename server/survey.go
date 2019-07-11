package main

import (
	pb "SurveyManagement/api"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)
type Survey struct {
	ID primitive.ObjectID  `bson:"_id,omitempty"`
	Description string
	Questions[]* pb.Question
}

var surveyCollection = client.Database("test").Collection("surveys")

func createSurveyDocument(doc Survey) {
	insertResult, err := surveyCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(insertResult)
}