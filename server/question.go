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
	Type string
	QuestionsWithLanguage pb.Question_QuestionWithLanguage
}

var questionCollection = client.Database("test").Collection("questions")


func getAllQuestions() []*Question{
	var questionResults []*Question
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

func getQuestionsByLanguage(language string) []*Question{
	var questionResults []*Question
	findOptions := options.Find()
	filter := bson.M{"language": language}

	cur, err := questionCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
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