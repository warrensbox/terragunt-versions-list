package main

import (
	"encoding/json"
	"io/ioutil"

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

	// fileList := "list"
	// if !lib.CheckFileExist(fileList) {
	// 	err := ioutil.WriteFile(fileList, []byte("Hello"), 0600)
	// 	if err != nil {
	// 		fmt.Printf("Unable to write file: %v", err)
	// 	}
	// }

	var list List
	list.Versions = allVersions

	file, _ := json.MarshalIndent(list, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}

type List struct {
	Versions []string `json:"versions"`
}
