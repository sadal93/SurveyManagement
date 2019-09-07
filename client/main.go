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

	textEn := "Tell us about your week"
	textDe := "Erzählen Sie uns von Ihrer Woche."
	text := make(map[string]string)
	text["english"] = textEn
	text["german"] = textDe

	/*fileEn, err := ioutil.ReadFile("barn.jpg")
	fileDe, err := ioutil.ReadFile("Meteora.jpg")

	if err != nil {
		panic(err.Error())
	}
	image := make(map[string] []byte)
	image["english"] = fileEn
	image["german"] = fileDe*/

	matrixText1En := "How active have you been?"
	matrixText1De := "Wie aktiv waren Sie?"
	matrixText1 := make(map[string]string)
	matrixText1["english"] = matrixText1En
	matrixText1["german"] = matrixText1De

	options1En := pb.Question_Options{OptionText: []string{"very active", "active", "not active at all"}}
	options1De := pb.Question_Options{OptionText: []string{"sehr aktiv", "aktiv", "nicht aktiv"}}
	options1 := make(map[string] *pb.Question_Options)
	options1["english"] = &options1En
	options1["german"] = &options1De

	matrixOption1 := &pb.Question_MatrixOptions{Text: matrixText1, Options: options1, SelectedOption: ""}

	matrixText2En := "How are you feeling now?"
	matrixText2De := "Wie geht es Ihnen heute?"
	matrixText2 := make(map[string]string)
	matrixText2["english"] = matrixText2En
	matrixText2["german"] = matrixText2De

	options2En := pb.Question_Options{OptionText: []string{"good", "normal", "bad"}}
	options2De := pb.Question_Options{OptionText: []string{"gut", "normal", "schlecht"}}
	options2 := make(map[string] *pb.Question_Options)
	options2["english"] = &options2En
	options2["german"] = &options2De

	matrixOption2 := &pb.Question_MatrixOptions{Text: matrixText2, Options: options2, SelectedOption: ""}

	var matrixOptions []*pb.Question_MatrixOptions
	matrixOptions = append(matrixOptions, matrixOption1, matrixOption2)



	/*multipleOptions := []string{"a", "b", "c"}
	selectedOption := "a"
	next := "12345"

	nextQuestion := &pb.Question_HideQuestion{SelectedOption: selectedOption, Question: next}
	var nextQuestions []*pb.Question_HideQuestion
	nextQuestions = append(nextQuestions, nextQuestion)
	answerOptions2 := &pb.Question_AnswerOptions{MultipleOptions: multipleOptions, NextQuestion: nextQuestions}*/

	answerOptions := &pb.Question_AnswerOptions{MatrixOptions: matrixOptions}

	questionWithLanguage := &pb.Question_QuestionWithLanguage{Text: text, AnswerOptions : answerOptions}







	var question1 = &pb.Question{Id: "5d6928bee278a85b3f2c96d6",Type: qType, QuestionWithLanguage: questionWithLanguage}
	//var question2  = &pb.Question{Type: qType, AnswerOptions: answerOptions2}

	var questions []*pb.Question
	questions = append(questions, question1)

	var survey = pb.SurveyData{ Description: "This is a Survey", Type: "timely", Questions: questions}
	//createQuestion(c, question1)
	createSurvey(c, survey)

}