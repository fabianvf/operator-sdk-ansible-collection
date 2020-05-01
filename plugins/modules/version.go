package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	ver "github.com/operator-framework/operator-sdk/version"
)

type ModuleArgs struct{}

type Response struct {
	Msg                string `json:"msg"`
	Changed            bool   `json:"changed"`
	Failed             bool   `json:"failed"`
	OperatorSDKVersion string `json:"operator_sdk_version,omitempty"`
	GitCommit          string `json:"git_commit,omitempty"`
	KubernetesVersion  string `json:"kubernetes_version,omitempty"`
	GoVersion          string `json:"go_version,omitempty"`
}

func ExitJson(responseBody Response) {
	returnResponse(responseBody)
}

func FailJson(responseBody Response) {
	responseBody.Failed = true
	returnResponse(responseBody)
}

func returnResponse(responseBody Response) {
	var response []byte
	var err error
	response, err = json.Marshal(responseBody)
	if err != nil {
		response, _ = json.Marshal(Response{Msg: "Invalid response object"})
	}
	fmt.Println(string(response))
	if responseBody.Failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func main() {
	var response Response

	if len(os.Args) != 2 {
		response.Msg = "No argument file provided"
		FailJson(response)
	}

	argsFile := os.Args[1]

	text, err := ioutil.ReadFile(argsFile)
	if err != nil {
		response.Msg = "Could not read configuration file: " + argsFile
		FailJson(response)
	}

	var moduleArgs ModuleArgs
	err = json.Unmarshal(text, &moduleArgs)
	if err != nil {
		response.Msg = "Configuration file not valid JSON: " + argsFile
		FailJson(response)
	}

	run(moduleArgs, response)
}

func run(args ModuleArgs, response Response) {
	version := ver.GitVersion
	if version == "unknown" {
		version = ver.Version
	}

	response.OperatorSDKVersion = version
	response.GitCommit = ver.GitCommit
	response.KubernetesVersion = ver.KubernetesVersion
	response.GoVersion = ver.GoVersion

	response.Msg = fmt.Sprintf("operator-sdk version: %s, commit: %s, kubernetes version: %s, go version: %s",
		version, ver.GitCommit, ver.KubernetesVersion, ver.GoVersion)

	ExitJson(response)
}
