package main

import (
	"testing"
	"fmt"
)


func TestParseSonarQubeTaskResult(t *testing.T) {
	success := parseSonarQubeTaskResult("http", "localhost", "9000", "AVrSaxLAo2-alXUM2_ac")
	if success {
		fmt.Print("Success")
	} else {
		fmt.Print("Error")
	}
}
