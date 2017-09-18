package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetAPIURL(t *testing.T) { // Test
	var inp = []string{"https://github.com/klyngen/Mcopy", "https://github.com/klyngen/pgnToFen", "https://github.com/"}
	var out = []string{"https://api.github.com/repos/klyngen/Mcopy", "https://api.github.com/repos/klyngen/pgnToFen", ""}

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
