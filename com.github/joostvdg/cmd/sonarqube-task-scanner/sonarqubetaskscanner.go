package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"
	"io/ioutil"
	"os"
)

type SonarTask struct {
	Task struct {
		Id                 string //`json:"id"`
		TaskType           string `json:"type"`
		ComponentId        string //`//json:"componentId"`
		ComponentKey       string //`json:"componentKey"`
		ComponentName      string //`json:"componentName"`
		ComponentQualifier string //`json:"componentQualifier"`
		AnalysisId         string //`json:"analysisId"`
		Status             string //`json:"status"`
		SubmittedAt        string //`json:"submittedAt"`
		StartedAt          string //`json:"number"`
		ExecutedAt         string //`json:"executedAt"`
		ExecutionTimeMs    int    //`json:"executionTimeMs"`
		Logs               bool   //`json:"logs"`
		HasScannerContext  bool   //`json:"hasScannerContext"`
	}
}

func main() {
	if len (os.Args) > 3 {
		parseSonarQubeTaskResult(os.Args[0], os.Args[1], os.Args[2], os.Args[4])
	} else {
		// TODO: add panic
	}

}

func parseSonarQubeTaskResult(protocol string, host string, port string, taskId string) (bool){

	url := fmt.Sprintf("%v://%v:%v/api/ce/task?id=%v", protocol, host, port, taskId)
	sonarClient  := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}


	res, getErr := sonarClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	fmt.Printf("Response: %v", res)

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	sonarTask := SonarTask{}
	masrshallErr := json.Unmarshal(body, &sonarTask)
	if masrshallErr != nil {
		fmt.Println(err)
		return false
	}

	fmt.Printf("SonarTask=%v\n", sonarTask)
	fmt.Printf("Task=%v\n", sonarTask.Task)
	fmt.Printf("Key=%v\n", sonarTask.Task.ComponentKey)
	fmt.Printf("Name=%v\n", sonarTask.Task.ComponentName)
	fmt.Printf("Status=%v\n", sonarTask.Task.Status)

	return sonarTask.Task.Status == "SUCCESS"
}
