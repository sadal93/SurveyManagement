package main

import (
	pb "SurveyManagement/api"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Question struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Language string
	Text string
	Type string
	AnswerOptions* pb.Question_AnswerOptions
}

var questionCollection = client.Database("test").Collection("questions")
var questionResults []*Question

func getAllQuestions() []*Question{
	if err != nil{
		log.Fatal(err)
	}

	findOptions := options.Find()
	cur, err := questionCollection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem Question
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		questionResults = append(questionResults, &elem)
	}
	return questionResults
}

func createQuestionDocument(doc Question) {
	insertResult, err := questionCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(insertResult)
}