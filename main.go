package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/warrensbox/terragrunt-versions-list/lib"
	"github.com/warrensbox/terragrunt-versions-list/modal"
)

const (
	terragruntURL = "https://api.github.com/repos/gruntwork-io/terragrunt/releases?"
)

var version = "0.2.0\n"

var CLIENT_ID = "xxx"
var CLIENT_SECRET = "xxx"

func main() {

	var client modal.Client

	client.ClientID = CLIENT_ID
	client.ClientSecret = CLIENT_SECRET

	allVersions, _ := lib.GetAppList(terragruntURL, &client)
	var list List
	list.Versions = allVersions

	currentTime := time.Now()
	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))
	lastUpdate := currentTime.Format("2006.01.02 15:04:05")
	list.LastUpdated = fmt.Sprintf("%s UTC", lastUpdate)
	file, _ := json.MarshalIndent(list, "", " ")

	_ = ioutil.WriteFile("index.json", file, 0644)
}

type List struct {
	LastUpdated string   `json:"lastUpdated"`
	Versions    []string `json:"versions"`
}
