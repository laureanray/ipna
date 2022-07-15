package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// TODO: Map json result to this
type GithubResult struct {
	TotalCount int `json:"total_count"`
	Items      []Item
}

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Owner       Owner
}

type Owner struct {
	Login string `json:"login"`
}

const GithubApi = "https://api.github.com"
const GITHUB_USERNAME = "github_username"
const GITHUB_TOKEN = "github_token"

func SearchGithubRepos(query string) {
	log.Println("Searching Github Repos")
	url := fmt.Sprintf(GithubApi+"/search/repositories?q=%s", query)
	log.Println(url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("%s:%s", viper.GetString(GITHUB_USERNAME), viper.GetString(GITHUB_TOKEN)))
	req.Header.Set("Accept", "application/vnd.github.text-match+json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("No response")
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result GithubResult

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cannot unmarshall JSON")
	}

	log.Print(result)

}
