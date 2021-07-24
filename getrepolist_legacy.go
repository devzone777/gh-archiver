package main
 
import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)
 
type GithubRepos struct {
	GithubRepos []Repos `json:"repositories"`
}
 
type Repos struct {
	Username string `json: "username"`
	Name    string `json: "name"`
	Url    string   `json: "url"`
}
 
func main() {
	file, _ := ioutil.ReadFile("gh_full_repo_list.json")
 
	data := GithubRepos{}
 
	_ = json.Unmarshal([]byte(file), &data)

	fileout, fileErr := os.Create("repolist.txt")
	if fileErr != nil {
            fmt.Println(fileErr)
            return
	}
 
	for i := 0; i < len(data.GithubRepos); i++ {
		fmt.Fprintln(fileout, "Username: ", data.GithubRepos[i].Username)
		fmt.Fprintln(fileout, "Name: ", data.GithubRepos[i].Name)
		fmt.Fprintln(fileout, "Url: ", data.GithubRepos[i].Url, ",")
	}
 
}


