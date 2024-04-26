package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fetchLatestWorkspaceGradlePlugin()
	buidReleaseFiles()
}

func fetchLatestWorkspaceGradlePlugin() {
	httpClient := getHttpClient()
	resp, err := httpClient.Get("https://search.maven.org/solrsearch/select?q=a:com.liferay.gradle.plugins.workspace&rows=1&wt=json")

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

func buidReleaseFiles() {
	start := time.Now()
	httpClient := getHttpClient()
	fmt.Print("Get releases.json")
	resp, err := httpClient.Get("https://releases.liferay.com/releases.json")

	if err != nil {
		fmt.Printf(" ❌ (%.2f s)\n", time.Since(start).Seconds())
		panic(err)
	}

	fmt.Printf(" ✅ (%.2f s)\n", time.Since(start).Seconds())
	start = time.Now()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var releases []Release
	json.Unmarshal(body, &releases)

	var dxp74, dxp73, dxp72, dxp71, dxp70 []Release
	var portal74, portal73, portal72, portal71, portal70 []Release

	for _, release := range releases {
		release.FetchProperties()
		if release.Product == "dxp" {
			dxp74, dxp73, dxp72, dxp71, dxp70 = updateDXPReleases(release, dxp74, dxp73, dxp72, dxp71, dxp70)
		}
		if release.Product == "portal" {
			portal74, portal73, portal72, portal71, portal70 = updatePortalReleases(release, portal74, portal73, portal72, portal71, portal70)
		}
	}

	fmt.Printf("\nDXP 7.4 (%v releases)\n", len(dxp74))
	fmt.Printf("DXP 7.3 (%v releases)\n", len(dxp73))
	fmt.Printf("DXP 7.2 (%v releases)\n", len(dxp72))
	fmt.Printf("DXP 7.1 (%v releases)\n", len(dxp71))
	fmt.Printf("DXP 7.0 (%v releases)\n\n", len(dxp70))

	fmt.Printf("Portal 7.4 (%v releases)\n", len(portal74))
	fmt.Printf("Portal 7.3 (%v releases)\n", len(portal73))
	fmt.Printf("Portal 7.2 (%v releases)\n", len(portal72))
	fmt.Printf("Portal 7.1 (%v releases)\n", len(portal71))
	fmt.Printf("Portal 7.0 (%v releases)\n", len(portal70))

	var dxpReleases, portalReleases []Release

	dxpReleases = append(dxpReleases, dxp74...)
	dxpReleases = append(dxpReleases, dxp73...)
	dxpReleases = append(dxpReleases, dxp72...)
	dxpReleases = append(dxpReleases, dxp71...)
	dxpReleases = append(dxpReleases, dxp70...)

	portalReleases = append(portalReleases, portal74...)
	portalReleases = append(portalReleases, portal73...)
	portalReleases = append(portalReleases, portal72...)
	portalReleases = append(portalReleases, portal71...)
	portalReleases = append(portalReleases, portal70...)

	writeReleaseFile("dxp", "", dxpReleases)
	writeReleaseFile("portal", "", portalReleases)

	writeReleaseFile("dxp", "7.4", dxp74)
	writeReleaseFile("dxp", "7.3", dxp73)
	writeReleaseFile("dxp", "7.2", dxp72)
	writeReleaseFile("dxp", "7.1", dxp71)
	writeReleaseFile("dxp", "7.0", dxp70)

	writeReleaseFile("portal", "7.4", portal74)
	writeReleaseFile("portal", "7.3", portal73)
	writeReleaseFile("portal", "7.2", portal72)
	writeReleaseFile("portal", "7.1", portal71)
	writeReleaseFile("portal", "7.0", portal70)
}

func writeReleaseFile(edition, version string, releases []Release) {
	var pathBuilder strings.Builder

	pathBuilder.WriteString("releases/")
	pathBuilder.WriteString(edition)
	if version != "" {
		pathBuilder.WriteString("_")
		pathBuilder.WriteString(strings.ReplaceAll(version, ".", ""))
	}
	pathBuilder.WriteString("_releases.json")

	fileContent, _ := json.MarshalIndent(releases, "", " ")

	err := os.WriteFile(pathBuilder.String(), fileContent, 0644)

	if err != nil {
		panic(err)
	}
}

func updateDXPReleases(release Release, dxp74, dxp73, dxp72, dxp71, dxp70 []Release) ([]Release, []Release, []Release, []Release, []Release) {
	if release.ProductGroupVersion == "7.4" || strings.Contains(release.ProductGroupVersion, ".q") {
		dxp74 = append(dxp74, release)
	}
	if release.ProductGroupVersion == "7.3" {
		dxp73 = append(dxp73, release)
	}
	if release.ProductGroupVersion == "7.2" {
		dxp72 = append(dxp72, release)
	}
	if release.ProductGroupVersion == "7.1" {
		dxp71 = append(dxp71, release)
	}
	if release.ProductGroupVersion == "7.0" {
		dxp70 = append(dxp70, release)
	}
	return dxp74, dxp73, dxp72, dxp71, dxp70
}

func updatePortalReleases(release Release, portal74, portal73, portal72, portal71, portal70 []Release) ([]Release, []Release, []Release, []Release, []Release) {
	if release.ProductGroupVersion == "7.4" {
		portal74 = append(portal74, release)
	}
	if release.ProductGroupVersion == "7.3" {
		portal73 = append(portal73, release)
	}
	if release.ProductGroupVersion == "7.2" {
		portal72 = append(portal72, release)
	}
	if release.ProductGroupVersion == "7.1" {
		portal71 = append(portal71, release)
	}
	if release.ProductGroupVersion == "7.0" {
		portal70 = append(portal70, release)
	}
	return portal74, portal73, portal72, portal71, portal70
}

func (release *Release) FetchProperties() {
	httpClient := getHttpClient()
	start := time.Now()
	releasePropertiesURL := release.URL + "/release.properties"
	fmt.Print("Get " + releasePropertiesURL)
	resp, err := httpClient.Get(releasePropertiesURL)

	if err != nil {
		fmt.Printf(" ❌ (%.2f s)\n", time.Since(start).Seconds())
		panic(err)
	}

	fmt.Printf(" ✅ (%.2f s)\n", time.Since(start).Seconds())

	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	releasePropertiesDirPath := getPathFromURL(releasePropertiesURL)
	releasePropertiesPath := filepath.Join(releasePropertiesDirPath, "release.properties")

	err = os.MkdirAll(releasePropertiesDirPath, os.ModePerm)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(releasePropertiesPath, bodyBytes, 0644)

	if err != nil {
		panic(err)
	}

	config, err := ReadPropertiesFile(resp.Body)

	release.ReleaseProperties = ReleaseProperties{
		AppServerTomcatVersion: config["app.server.tomcat.version"],
		BuildTimestamp:         config["build.timestamp"],
		BundleChecksumSha512:   config["bundle.checksum.sha512"],
		BundleURL:              config["bundle.url"],
		GitHashLiferayDocker:   config["git.hash.liferay-docker"],
		GitHasLiferayPortalEE:  config["git.hash.liferay-portal-ee"],
		LiferayDockerImage:     config["liferay.docker.image"],
		LiferayDockerTags:      config["liferay.docker.tags"],
		LiferayProductVersion:  config["liferay.product.version"],
		ReleaseDate:            config["release.date"],
		TargetPlatformVersion:  config["target.platform.version"],
	}

	if err != nil {
		panic(err)
	}
}

func getPathFromURL(url string) string {
	var pathBuilder strings.Builder
	urlArray := strings.Split(url, "/")
	urlArray = urlArray[3 : len(urlArray)-1]

	pathBuilder.WriteString("releases")
	for _, part := range urlArray {
		pathBuilder.WriteString("/")
		pathBuilder.WriteString(part)
	}

	return pathBuilder.String()
}

func getHttpClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 60 * time.Second,
		},
		Timeout: 240 * time.Second,
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

func ReadPropertiesFile(filecontent io.Reader) (AppConfigProperties, error) {
	config := AppConfigProperties{}

	scanner := bufio.NewScanner(filecontent)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}

type Release struct {
	Product               string            `json:"product"`
	ProductGroupVersion   string            `json:"productGroupVersion"`
	ProductVersion        string            `json:"productVersion"`
	Promoted              string            `json:"promoted"`
	ReleaseKey            string            `json:"releaseKey"`
	TargetPlatformVersion string            `json:"targetPlatformVersion"`
	URL                   string            `json:"url"`
	ReleaseProperties     ReleaseProperties `json:"releaseProperties"`
}

type ReleaseProperties struct {
	AppServerTomcatVersion string `json:"appServerTomcatVersion"`
	BuildTimestamp         string `json:"buildTimestamp"`
	BundleChecksumSha512   string `json:"bundleChecksumSha512"`
	BundleURL              string `json:"bundleURL"`
	GitHashLiferayDocker   string `json:"gitHashLiferayDocker"`
	GitHasLiferayPortalEE  string `json:"gitHashLiferayPortalEE"`
	LiferayDockerImage     string `json:"liferayDockerImage"`
	LiferayDockerTags      string `json:"liferayDockerTags"`
	LiferayProductVersion  string `json:"liferayProductVersion"`
	ReleaseDate            string `json:"releaseDate"`
	TargetPlatformVersion  string `json:"targetPlatformVersion"`
}

type AppConfigProperties map[string]string
