package main

import (
	pb "SurveyManagement/api"
	sm "StudyManagement/api"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"labix.org/v2/mgo/bson"
	"log"
	"net"
	"time"
)

const (
	address     = "localhost:50051"
)

const (
	port = ":50052"
)

var clientOptions  = options.Client().ApplyURI("mongodb://localhost:27017")
var client, err = mongo.Connect(context.TODO(), clientOptions)

type server struct{}

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

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := sm.NewStudyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetAllStudies(ctx, &sm.Empty{})
	if err != nil {
		log.Fatalf("could not retrieve: %v", err)
	}
	fmt.Println("Studies: " , r)

	survey := Survey{primitive.NewObjectID(), surveyData.Description, surveyData.Questions}
	createSurveyDocument(survey)

	log.Printf("Survey Created: %v", survey)
	return &pb.SurveyData{Id: surveyData.Id, Description: surveyData.Description, Questions: surveyData.Questions}, nil
}

func (s *server)  CreateQuestion(ctx context.Context, questionData *pb.Question) (*pb.Question, error) {

	question := Question{primitive.NewObjectID(), questionData.Type, questionData.QuestionWithLanguage}
	createQuestionDocument(question)

	log.Printf("Question Created: %v", question)
	return &pb.Question{Id: questionData.Id, Type: questionData.Type, QuestionWithLanguage: questionData.QuestionWithLanguage}, nil
}

func (s *server) DeleteQuestion(ctx context.Context, question *pb.QuestionID) (*pb.EmptySurvey, error) {

	objectID, err := primitive.ObjectIDFromHex(question.QuestionID)
	filter := bson.M{"_id": objectID}
	deleteResult, err := questionCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the questions collection\n", deleteResult.DeletedCount)

	return &pb.EmptySurvey{}, nil
}

func (s *server) GetQuestion(ctx context.Context, question *pb.QuestionID) (*pb.Question, error) {
	var result Question
	objectID, err := primitive.ObjectIDFromHex(question.QuestionID)
	filter := bson.M{"_id": objectID}

	err = questionCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.Question{Id: result.ID.Hex(), Type: result.Type, QuestionWithLanguage: result.QuestionsWithLanguage }, nil

}

func (s *server) GetAllQuestions(ctx context.Context, empty *pb.EmptySurvey) (*pb.QuestionArray, error) {

	var questions []*pb.Question
	documents := getAllQuestions()
	for _, document := range documents{
		var question *pb.Question = new(pb.Question)
		question.Id = document.ID.Hex()
		question.Type = document.Type
		question.QuestionWithLanguage = document.QuestionsWithLanguage


		questions = append(questions, question)
	}
	return &pb.QuestionArray{Questions: questions}, nil
}
