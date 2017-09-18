package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// OwnerT - struct containing ownerinfo
type OwnerT struct {
	Login string
}

// OwnerContainer - Contains an owner
type OwnerContainer struct {
	Owner OwnerT
}

// TopComitter - struct for extracting the top comitter
type TopComitter struct {
	Login         string
	Contributions int
}

// Payload Struct to contain information to be encoded to JSON
type Payload struct {
	Name      string   `json:"name"`
	Owner     string   `json:"owner"`
	Committer string   `json:"committer"`
	Committs  int      `json:"commits"`
	Languages []string `json:"languages"`
}

/*
Creates a URL to the github-api
output should be api.github.com/repos/:owner/:repo
*/

func getAPIURL(uri string) (url string) {
	url = "https://api.github.com/repos/"
	var part = strings.TrimLeft(uri, "/projectinfo/v1/https://github.com/")

	if len(part) != 0 {
		url += part
	} else {
		url = ""
	}

	return
}

func requestData(url string) (data []byte) {
	r, error := http.Get(url) // Makes request

	if error != nil {
		fmt.Print(error)
		return nil
	}

	data, err := ioutil.ReadAll(r.Body) // Reads into a type byte[]
	if err != nil {                     // If error occurred
		fmt.Println(err.Error()) // print error occuring during read
		return nil
	}

	defer r.Body.Close()

	return
}

func generatePayload(url string) (payload Payload) {
	owner := new(OwnerContainer)
	langs := new(map[string]interface{})
	contrib := new([1]TopComitter)

	// Gets data from api
	mainResponse := requestData(url)                         // Gets root of api
	languageResponse := requestData(url + "/languages")      // Gets language from api
	contributeResponse := requestData(url + "/contributors") // Gets contributor data

	collectData(mainResponse, &payload)      // Reads repo name
	collectData(mainResponse, owner)         // Reads owner name
	collectData(languageResponse, langs)     // Reads languages
	collectData(contributeResponse, contrib) // Reads contribution info

	// Assign values to struct
	payload.Committer = contrib[0].Login
	payload.Committs = contrib[0].Contributions
	payload.Owner = owner.Owner.Login

	for r := range *langs {
		payload.Languages = append(payload.Languages, r)
	}

	return
}

// Parses the request-body and puts it in the referenced variable
func collectData(data []byte, payload interface{}) {
	json.Unmarshal(data, &payload)
}

func handler(w http.ResponseWriter, r *http.Request) {
	payload := generatePayload(getAPIURL(r.URL.RequestURI()))

	output, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", output)

}

func main() {
	http.HandleFunc("/projectinfo/v1/", handler) // Setter en spesifik handler for root
	http.ListenAndServe(":8081", nil)

}