package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetAPIURL(t *testing.T) { // Test
	var inp = []string{"https://github.com/klyngen/Mcopy", "https://github.com/klyngen/pgnToFen", "https://github.com/"}
	var out = []string{"https://api.github.com/repos/klyngen/Mcopy", "https://api.github.com/repos/klyngen/pgnToFen", "j"}

	for i := range inp {
		if getAPIURL(inp[i]) != out[i] {
			t.Fatalf("ERROR expected: %s but got: %s", out[i], getAPIURL(inp[i]))
		}
	}
}

func TestCollectData(t *testing.T) {
	var payload Payload

	data, error := ioutil.ReadFile("./root.json")
	//fmt.Println(data)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		collectData(data, &payload)
	}

	if payload.Name != "EIS" {
		t.Fatalf("ERROR expected EIS but got: %s", payload.Name)
	}
}

func TestPayloadGeneration(t *testing.T) {
	//												CORRECT									WRONG		  REALLY WRONG
	adresses := []string{"github.com/klyngen/eis", "github.com/", "vg.no"}
	outputsReponame := []string{"EIS", "", ""}
	outputsOwner := []string{"klyngen", "", ""}

	for x := range adresses {
		payload := generatePayload(getAPIURL(adresses[x]))

		if payload.Name != outputsReponame[x] {
			t.Fatalf("ERROR expected %s but got %s", outputsReponame[x], payload.Name)
		}

		if payload.Owner != outputsOwner[x] {
			t.Fatalf("ERROR expected %s but got %s", outputsOwner[x], payload.Owner)
		}

	}
}

func TestVeridyPayload(t *testing.T) {

	var test1, test2, test3, test4 Payload
	// Test 1 should pass the verification
	test1.Committer = "Some comitter"
	test1.Committs = 250
	test1.Name = "Some Name"
	test1.Owner = "Some owner"

	// Test2
	test2 = test1
	test2.Committs = 0
	test2.Committer = ""

	test3 = test2
	test3.Name = ""

	payloads := []Payload{test1, test2, test3, test4}
	output := []bool{true, true, false, false}

	for r := range output {
		if verifyPayload(payloads[r]) != output[r] {
			t.Fatalf("Testing failed in condition %d", r)
		}
	}

}
