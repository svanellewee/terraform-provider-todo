package main

import (
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/svanellewee/todoclient/client"
	"github.com/svanellewee/todoclient/client/operations"
	"github.com/svanellewee/todoclient/models"
	"log"
)

func main() {
	fmt.Println("Hiya")
	postParams := operations.NewPostTaskParams()
	description := "hello"
	postParams.Body = &models.Task{
		Description: &description,
	}
	transport := httptransport.New("127.0.0.1:8080", "", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)

	result, err := client.Operations.PostTask(postParams)
	if err != nil {
		log.Fatal("err", err)
	}
	fmt.Println(result)
	getParams := operations.NewGetTaskParams()
	getTaskOK, err := client.Operations.GetTask(getParams)
	if err != nil {
		log.Fatal("err", err)
	}
	for _, task := range getTaskOK.Payload {
		fmt.Printf("%p.. bla %v\n", task, task)
	}

	getTaskID := operations.NewGetTaskIDParams()
	getTaskID.ID = 1
	currentTask, err := client.Operations.GetTaskID(getTaskID)
	if err != nil {
		log.Fatal("ERRR", err)
	}
	fmt.Println(currentTask)

	currentTask.Payload.Completed = true

	updateParams := operations.NewUpdateOneParams()
	updateParams.ID = 1
	updateParams.Body = &models.Task{}
	updateParams.Body.Description = currentTask.Payload.Description
	updateParams.Body.Completed = true
	updateResult, err := client.Operations.UpdateOne(updateParams)
	if err != nil {
		log.Fatalf("err update %v\n", err)
	}
	fmt.Println(updateResult)
}
