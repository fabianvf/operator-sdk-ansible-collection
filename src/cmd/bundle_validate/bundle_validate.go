package main

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/operator-framework/api/pkg/manifests"
	"github.com/operator-framework/operator-registry/pkg/lib/bundle"
	helpers "github.com/operator-framework/operator-sdk-ansible-collection/pkg/ansible/mod_helpers"
)

type ModuleArgs struct {
	// TODO: Define the arguments your module takes here
	FilePath string `json:"file_path"`
	Image    string `json:"image"`
	Verbose  string `json:"verbose"`
}

type Response struct {
	helpers.BaseResponse
	VerboseOutput string   `json:"verbose_output"`
	Errors        []string `json:"errors"`
	Warnings      []string `json:"warnings"`
}

func main() {
	response := Response{}
	var moduleArgs ModuleArgs

	err := helpers.ParseArgs(os.Args, &moduleArgs)
	if err != nil {
		response.BaseResponse.Msg = err.Error()
		response.BaseResponse.Failed = true
		helpers.ExitJSON(response, true)
	}

	if moduleArgs.FilePath == "" && moduleArgs.Image == "" {
		response.BaseResponse.Msg = "File path or image is a required"
		response.BaseResponse.Failed = true
		helpers.ExitJSON(response, true)
	}

	if moduleArgs.FilePath != "" {
		if isExist(moduleArgs.FilePath) {
			response.BaseResponse.Msg = "File Path must point to existing file"
			response.BaseResponse.Failed = true
			helpers.ExitJSON(response, true)
		}
		buf := &bytes.Buffer{}
		logger := logrus.WithField("name", "bundle-test")
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(buf)
		val := bundle.NewImageValidator("", logger)

		if err := val.ValidateBundleFormat(dir); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Failed = true
		}

		// Validate bundle content.
		manifestsDir := filepath.Join(moduleArgs.FilePath, bundle.ManifestsDir)
		_, _, validationResults := manifests.GetManifestsDir(dir)
		for _, result := range validationResults {
			for _, e := range result.Errors {
				response.Errors = append(response.Errors, e.Error())
				response.Failed = true
			}
			for _, w := range result.Warnings {
				response.Warnings = append(response.Warnings, w.Error())
			}
		}
		if err := val.ValidateBundleContent(manifestsDir); err != nil {
			response.Errors = append(response.Errors, err.Error())
			response.Failed = true
		}
		response.VerboseOutput = buf.String()
	}

	if response.Failed {
		response.BaseResponse.Msg = "validation errors found"
		helpers.ExitJSON(response, true)
	}
	response.BaseResponse.Msg = "validation passed"
	helpers.ExitJSON(response, false)
}

// isExist returns true if path exists.
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
