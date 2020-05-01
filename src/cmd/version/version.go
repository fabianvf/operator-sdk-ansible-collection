package main

import (
	"fmt"

	helpers "github.com/operator-framework/operator-sdk-ansible-collection/pkg/ansible/mod_helpers"
	ver "github.com/operator-framework/operator-sdk/version"
)

type ModuleArgs struct{}

type Response struct {
	helpers.BaseResponse
	OperatorSDKVersion string `json:"operator_sdk_version,omitempty"`
	GitCommit          string `json:"git_commit,omitempty"`
	KubernetesVersion  string `json:"kubernetes_version,omitempty"`
	GoVersion          string `json:"go_version,omitempty"`
}

func main() {
	var response Response

	run(response)
}

func run(response Response) {
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

	helpers.ExitJSON(response, false)
}
