package main

import (
	pb "StudyManagement/api"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestServer_CreateSurvey(t *testing.T)  {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial localhost: %v", err)
	}
	defer conn.Close()
	client := pb.NewSurveyClient(conn)
	resp, err := client.CreateSurvey(ctx, &pb.SurveyData{Description: ""})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	fmt.Print(resp)
}

func TestServer_GetSurvey(t *testing.T){
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