package main

import (
	pb "SurveyManagement/api"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:50052"
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
	r, err := c.CreateQuestion(ctx, &pb.Question{Type: question.Type, QuestionWithLanguage: question.QuestionWithLanguage})
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

	qType := "matrix"

	languageEn := "english"
	textEn := "This is question 1"

	matrixText1En := "How happy are you?"
	options1En := []string{"happy", "normal", "sad"}
	matrixOption1En := &pb.Question_MatrixOptions{Text: matrixText1En, Options: options1En}

	matrixText2En := "How are you feeling?"
	options2En := []string{"good", "normal", "bad"}
	matrixOption2En := &pb.Question_MatrixOptions{Text: matrixText2En, Options: options2En}

	var matrixOptionsEn []*pb.Question_MatrixOptions
	matrixOptionsEn = append(matrixOptionsEn, matrixOption1En, matrixOption2En)

	languageDe := "german"
	textDe := "Das ist Frage 1"

	matrixText1De := "Wie glucklisch sind Sie?"
	options1De := []string{"sehr", "normal", "traurig"}
	matrixOption1De := &pb.Question_MatrixOptions{Text: matrixText1De, Options: options1De}

	matrixText2De := "Wie geht es Ihnen?"
	options2De := []string{"gut", "normal", "schlecht"}
	matrixOption2De := &pb.Question_MatrixOptions{Text: matrixText2De, Options: options2De}

	var matrixOptionsDe []*pb.Question_MatrixOptions
	matrixOptionsDe = append(matrixOptionsDe, matrixOption1De, matrixOption2De)



	/*multipleOptions := []string{"a", "b", "c"}
	selectedOption := "a"
	next := "12345"

	nextQuestion := &pb.Question_HideQuestion{SelectedOption: selectedOption, Question: next}
	var nextQuestions []*pb.Question_HideQuestion
	nextQuestions = append(nextQuestions, nextQuestion)
	answerOptions2 := &pb.Question_AnswerOptions{MultipleOptions: multipleOptions, NextQuestion: nextQuestions}*/

	answerOptions1En := &pb.Question_AnswerOptions{MatrixOptions: matrixOptionsEn}
	answerOptions1De := &pb.Question_AnswerOptions{MatrixOptions: matrixOptionsDe}

	questionWithLanguageEn := &pb.Question_QuestionWithLanguage{Language: languageEn, Text:	textEn, AnswerOptions:answerOptions1En}
	questionWithLanguageDe := &pb.Question_QuestionWithLanguage{Language: languageDe, Text:	textDe, AnswerOptions:answerOptions1De}

	var questionsWithLanguage []*pb.Question_QuestionWithLanguage
	questionsWithLanguage = append(questionsWithLanguage, questionWithLanguageEn, questionWithLanguageDe)







	var question1 = &pb.Question{Id: "5d38cad2fc0b1b1fd7444274",Type: qType, QuestionWithLanguage: questionsWithLanguage}
	//var question2  = &pb.Question{Type: qType, AnswerOptions: answerOptions2}

	var questions []*pb.Question
	questions = append(questions, question1)

	var survey = pb.SurveyData{ Description: "This is a Survey", Questions: questions}
	//createQuestion(c, question1)
	createSurvey(c, survey)

}