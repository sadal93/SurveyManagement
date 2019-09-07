package main

import (
	pb "SurveyManagement/api"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"testing"
)


func TestCreateOpenQuestion(t *testing.T)  {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial localhost: %v", err)
	}
	defer conn.Close()
	client := pb.NewSurveyClient(conn)

	qType := "open"

	textEn := "How old are you?"
	textDe := "Wie alt sind Sie?."
	text := make(map[string]string)
	text["english"] = textEn
	text["german"] = textDe

	openOptions := &pb.Question_OpenOptions{DataType: "integer", Limit: 3, OpenAnswer: ""}

	answerOptions := &pb.Question_AnswerOptions{OpenOptions: openOptions}

	questionWithLanguage := &pb.Question_QuestionWithLanguage{Text: text, AnswerOptions : answerOptions}

	resp, err := client.CreateQuestion(ctx, &pb.Question{Type: qType, QuestionWithLanguage: questionWithLanguage})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	fmt.Print(resp)
}

func TestCreateSingleChoiceQuestion(t *testing.T)  {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial localhost: %v", err)
	}
	defer conn.Close()
	client := pb.NewSurveyClient(conn)

	qType := "single"

	textEn := "How are you feeling today?"
	textDe := "Wie geht es Ihnen heute?"
	text := make(map[string]string)
	text["english"] = textEn
	text["german"] = textDe

	optionsEn := pb.Question_Options{OptionText: []string{"good", "normal", "bad"}}
	optionsDe := pb.Question_Options{OptionText: []string{"gut", "normal", "schlecht"}}
	options := make(map[string] *pb.Question_Options)
	options["english"] = &optionsEn
	options["german"] = &optionsDe

	singleOption := &pb.Question_SingleOption{Options: options, SelectedOption: ""}

	answerOptions := &pb.Question_AnswerOptions{SingleOption: singleOption}

	questionWithLanguage := &pb.Question_QuestionWithLanguage{Text: text, AnswerOptions : answerOptions}

	resp, err := client.CreateQuestion(ctx, &pb.Question{Type: qType, QuestionWithLanguage: questionWithLanguage})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	fmt.Print(resp)
}

func TestCreateMultipleChoiceQuestion(t *testing.T)  {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial localhost: %v", err)
	}
	defer conn.Close()
	client := pb.NewSurveyClient(conn)

	qType := "multiple"

	textEn := "What are your activities for today?"
	textDe := "Was machen Sie heute?"
	text := make(map[string]string)
	text["english"] = textEn
	text["german"] = textDe

	optionsEn := pb.Question_Options{OptionText: []string{"watch TV", "play tennis", "sleep", "cook dinner"}}
	optionsDe := pb.Question_Options{OptionText: []string{"fernsehen", "Tennis spielen", "schlafen", "Abendessen kochen"}}
	options := make(map[string] *pb.Question_Options)
	options["english"] = &optionsEn
	options["german"] = &optionsDe

	multipleOptions := &pb.Question_MultipleOptions{Options: options, SelectedOptions: []string{}}

	answerOptions := &pb.Question_AnswerOptions{MultipleOptions: multipleOptions}

	questionWithLanguage := &pb.Question_QuestionWithLanguage{Text: text, AnswerOptions : answerOptions}

	resp, err := client.CreateQuestion(ctx, &pb.Question{Type: qType, QuestionWithLanguage: questionWithLanguage})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	fmt.Print(resp)
}

func TestCreateSurvey(t *testing.T)  {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial localhost: %v", err)
	}
	defer conn.Close()
	client := pb.NewSurveyClient(conn)

	qType1 := "open"

	textEn1 := "How old are you?"
	textDe1 := "Wie alt sind Sie?."
	text1 := make(map[string]string)
	text1["english"] = textEn1
	text1["german"] = textDe1

	openOptions := &pb.Question_OpenOptions{DataType: "integer", Limit: 3, OpenAnswer: ""}

	answerOptions1 := &pb.Question_AnswerOptions{OpenOptions: openOptions}

	questionWithLanguage1 := &pb.Question_QuestionWithLanguage{Text: text1, AnswerOptions : answerOptions1}

	question1 := &pb.Question{Id: "5d72069c61dd85453b0dc775", Type: qType1, QuestionWithLanguage: questionWithLanguage1}



	qType2 := "single"

	textEn2 := "How are you feeling today?"
	textDe2 := "Wie geht es Ihnen heute?"
	text2 := make(map[string]string)
	text2["english"] = textEn2
	text2["german"] = textDe2

	optionsEn := pb.Question_Options{OptionText: []string{"good", "normal", "bad"}}
	optionsDe := pb.Question_Options{OptionText: []string{"gut", "normal", "schlecht"}}
	options := make(map[string] *pb.Question_Options)
	options["english"] = &optionsEn
	options["german"] = &optionsDe

	singleOption := &pb.Question_SingleOption{Options: options, SelectedOption: ""}

	answerOptions2 := &pb.Question_AnswerOptions{SingleOption: singleOption}

	questionWithLanguage2 := &pb.Question_QuestionWithLanguage{Text: text2, AnswerOptions : answerOptions2}

	question2 := &pb.Question{Id: "5d72074161dd85453b0dc776", Type: qType2, QuestionWithLanguage: questionWithLanguage2}



	qType3 := "multiple"

	textEn3 := "What are your activities for today?"
	textDe3 := "Was machen Sie heute?"
	text3 := make(map[string]string)
	text3["english"] = textEn3
	text3["german"] = textDe3

	optionsEn2 := pb.Question_Options{OptionText: []string{"watch TV", "play tennis", "sleep", "cook dinner"}}
	optionsDe2 := pb.Question_Options{OptionText: []string{"fernsehen", "Tennis spielen", "schlafen", "Abendessen kochen"}}
	options2 := make(map[string] *pb.Question_Options)
	options2["english"] = &optionsEn2
	options2["german"] = &optionsDe2

	multipleOptions := &pb.Question_MultipleOptions{Options: options, SelectedOptions: []string{}}

	answerOptions3 := &pb.Question_AnswerOptions{MultipleOptions: multipleOptions}

	questionWithLanguage3 := &pb.Question_QuestionWithLanguage{Text: text3, AnswerOptions : answerOptions3}

	question3 := &pb.Question{Id: "5d7207bb61dd85453b0dc777", Type: qType3, QuestionWithLanguage: questionWithLanguage3}



	var questions []*pb.Question
	questions = append(questions, question1, question2, question3)


	resp, err := client.CreateSurvey(ctx, &pb.SurveyData{Description: "This is a test Survey", Type: "triggered", Questions: questions})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	fmt.Print(resp)
}

func TestGetSurvey(t *testing.T){
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial localhost: %v", err)
	}
	defer conn.Close()
	client := pb.NewSurveyClient(conn)

	resp, err := client.GetQuestion(ctx, &pb.QuestionID{QuestionID: ""});
	if err != nil {
		log.Fatalf("Error on Add: %v", err)
	}
	fmt.Println(resp)
}