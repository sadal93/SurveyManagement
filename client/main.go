package main

import (
	pb "SurveyManagement/api"
	"context"
	"google.golang.org/grpc"
	"io/ioutil"
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
	textEn := "Tell us about your week"

	matrixText1En := "How active have you been?"
	options1En := []string{"very active", "active", "not active at all"}
	matrixOption1En := &pb.Question_MatrixOptions{Text: matrixText1En, Options: options1En}

	matrixText2En := "How are you feeling now?"
	options2En := []string{"good", "normal", "bad"}
	matrixOption2En := &pb.Question_MatrixOptions{Text: matrixText2En, Options: options2En}

	var matrixOptionsEn []*pb.Question_MatrixOptions
	matrixOptionsEn = append(matrixOptionsEn, matrixOption1En, matrixOption2En)

	languageDe := "german"
	textDe := "Erzählen Sie uns von Ihrer Woche."

	matrixText1De := "Wie aktiv waren Sie?"
	options1De := []string{"sehr aktiv", "aktiv", "nicht aktiv"}
	matrixOption1De := &pb.Question_MatrixOptions{Text: matrixText1De, Options: options1De}

	matrixText2De := "Wie geht es Ihnen heute?"
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



	file, err := ioutil.ReadFile("barn.jpg")

	if err != nil {
		panic(err.Error())
	}



	var question1 = &pb.Question{Id: "5d667f707af26867861650d2", Type: qType, Image:file, QuestionWithLanguage: questionsWithLanguage}
	//var question2  = &pb.Question{Type: qType, AnswerOptions: answerOptions2}

	var questions []*pb.Question
	questions = append(questions, question1)

	var survey = pb.SurveyData{ Description: "This is a Survey", Type: "timely", Questions: questions}
	//createQuestion(c, question1)
	createSurvey(c, survey)

}