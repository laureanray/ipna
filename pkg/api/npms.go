package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const NpmsApi = "https//api.npms.io/v2"

type NpmsResult struct {
	Total   int32 `json:"total"`
	Results []Result
}

type Result struct {
	Package     Package
	SearchScore float64 `json:"searchScore"`
}

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func SearchNpmsRepos(query string) {
	url := fmt.Sprintf(NpmsApi+"/search?q=%s", query)

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)

	if err != nil {
		log.Println("NO response")
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result NpmsResult

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshall JSON")
	}

	log.Println(result)
}
