package main

import (
	pb "SurveyManagement/api"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
type Survey struct {
	ID primitive.ObjectID  `bson:"_id,omitempty"`
	Description string
	Questions[]* pb.Question
}

var surveyCollection = client.Database("test").Collection("surveys")

func getAllSurveys() []*Survey{
	var surveyResults []*Survey
	if err != nil{
		log.Fatal(err)
	}

	findOptions := options.Find()
	cur, err := surveyCollection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem Survey
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		surveyResults = append(surveyResults, &elem)
	}
	return surveyResults
}

func createSurveyDocument(doc Survey) {
	insertResult, err := surveyCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(insertResult)
}