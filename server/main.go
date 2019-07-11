package main

import (
	pb "SurveyManagement/api"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

var clientOptions  = options.Client().ApplyURI("mongodb://localhost:27017")
var client, err = mongo.Connect(context.TODO(), clientOptions)

var surveyCollection = client.Database("test").Collection("surveys")
var questionCollection = client.Database("test").Collection("questions")

type Survey struct {
	ID primitive.ObjectID  `bson:"_id,omitempty"`
	Description string
	Questions[]* pb.Question
}

type Question struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Language string
	Text string
	Type string
	AnswerOptions* pb.Question_AnswerOptions
}

type server struct{}

func createSurveyDocument(doc Survey) {
	insertResult, err := surveyCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(insertResult)
}

func createQuestionDocument(doc Question) {
	insertResult, err := questionCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(insertResult)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSurveyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err )
	}
}

func (s *server)  CreateSurvey(ctx context.Context, surveyData *pb.SurveyData) (*pb.SurveyData, error) {

	survey := Survey{primitive.NewObjectID(), surveyData.Description, surveyData.Questions}
	createSurveyDocument(survey)

	log.Printf("Survey Created: %v", survey)
	return &pb.SurveyData{Id: surveyData.Id, Description: surveyData.Description, Questions: surveyData.Questions}, nil
}

func (s *server)  CreateQuestion(ctx context.Context, questionData *pb.Question) (*pb.Question, error) {

	question := Question{primitive.NewObjectID(), questionData.Language, questionData.Text, questionData.Type, questionData.AnswerOptions}
	createQuestionDocument(question)

	log.Printf("Question Created: %v", question)
	return &pb.Question{Id: questionData.Id, Language: questionData.Language, Type: questionData.Type, AnswerOptions: questionData.AnswerOptions}, nil
}