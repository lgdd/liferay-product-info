package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fetchLatestWorkspaceGradlePlugin()
}

func fetchLatestWorkspaceGradlePlugin() {
	resp, err := http.Get("https://search.maven.org/solrsearch/select?q=a:com.liferay.gradle.plugins.workspace&rows=1&wt=json")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var searchResponse MavenCentralSearchResponse
	json.Unmarshal(body, &searchResponse)

	if searchResponse.Body.NumFound == 0 {
		fmt.Println("could not find com.liferay.gradle.plugins.workspace in maven central")
		return
	}

	latestVersion := searchResponse.Body.Results[0].LatestVersion

	fmt.Printf("com.liferay.gradle.plugins.workspace=%s\n", latestVersion)

	err = os.WriteFile("com.liferay.gradle.plugins.workspace", []byte(latestVersion), 0644)
	if err != nil {
		panic(err)
	}
}

type MavenCentralSearchResponse struct {
	Body struct {
		NumFound int `json:"numFound"`
		Results  []struct {
			ID            string `json:"id"`
			Group         string `json:"g"`
			Artifact      string `json:"a"`
			LatestVersion string `json:"latestVersion"`
			Packaging     string `json:"p"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"docs"`
	} `json:"response"`
	Spellcheck struct {
		Suggestions []any `json:"suggestions"`
	} `json:"spellcheck"`
}
