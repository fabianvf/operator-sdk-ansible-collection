package mod_helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ParseArgs will parse the arguments and set them in the moduleArgs
func ParseArgs(args []string, moduleArgs interface{}) error {
	if len(args) != 2 {
		return fmt.Errorf("No argument file provided")
	}

	argsFile := args[1]

	text, err := ioutil.ReadFile(argsFile)
	if err != nil {
		return fmt.Errorf("Configuration file not valid JSON: %v", argsFile)
	}

	err = json.Unmarshal(text, moduleArgs)
	if err != nil {
		return fmt.Errorf("Configuration file not valid JSON: %v", argsFile)
	}
	return nil
}

type BaseResponse struct {
	Msg     string `json:"msg"`
	Changed bool   `json:"changed"`
	Failed  bool   `json:"failed"`
}

// ExitJSON will convert response to json and exit correctly based on failure
func ExitJSON(responseBody interface{}, failed bool) {
	var response []byte
	var err error
	response, err = json.Marshal(responseBody)
	if err != nil {
		response, _ = json.Marshal(BaseResponse{Msg: "Invalid response object"})
	}
	fmt.Println(string(response))
	if failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
