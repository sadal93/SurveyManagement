package main

import (
	pb "SurveyManagement/api"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
)

func createSurvey(c pb.SurveyClient, survey pb.SurveyData)  {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateSurvey(ctx, &pb.SurveyData{Description: survey.Description, Questions: survey.Questions})
	if err != nil {
		log.Fatalf("could not insert: %v", err)
	}

	log.Print("Survey = ", r)
}

func createQuestion(c pb.SurveyClient, question pb.Question)  {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateQuestion(ctx, &pb.Question{Language: question.Language, Text:question.Text, Type: question.Type, AnswerOptions: question.AnswerOptions})
	if err != nil {
		log.Fatalf("could not insert: %v", err)
	}

	log.Print("Question = ", r)
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSurveyClient(conn)

	language := "english"
	text := "This is question 1"
	qType := "multi"

	matrixText1 := "How happy are you?"
	options1 := []string{"happy", "normal", "sad"}
	matrixOption1 := &pb.Question_MatrixOptions{Text: matrixText1, Options: options1}

	matrixText2 := "How are you feeling?"
	options2 := []string{"good", "normal", "bad"}
	matrixOption2 := &pb.Question_MatrixOptions{Text: matrixText2, Options: options2}

	var matrixOptions []*pb.Question_MatrixOptions
	matrixOptions = append(matrixOptions, matrixOption1, matrixOption2)

	multipleOptions := []string{"a", "b", "c"}
	selectedOption := "a"
	next := "12345"

	nextQuestion := &pb.Question_NextQuestion{SelectedOption: selectedOption, Question: next}
	var nextQuestions []*pb.Question_NextQuestion
	nextQuestions = append(nextQuestions, nextQuestion)

	answerOptions1 := &pb.Question_AnswerOptions{MatrixOptions: matrixOptions}
	answerOptions2 := &pb.Question_AnswerOptions{MultipleOptions: multipleOptions, NextQuestion: nextQuestions}

	var question1 = &pb.Question{Language: language, Text: text, Type: qType, AnswerOptions: answerOptions1}
	var question2  = &pb.Question{Language: language, Text: text, Type: qType, AnswerOptions: answerOptions2}

	var questions []*pb.Question
	questions = append(questions, question1, question2)

	var survey = pb.SurveyData{ Description: "This is a Survey", Questions: questions}
	//createQuestion(c, question)
	createSurvey(c, survey)
}