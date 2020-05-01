package // TODO package name

import (
	"os"

	helpers "github.com/operator-framework/operator-sdk-ansible-collection/pkg/ansible/mod_helpers"
)

type ModuleArgs struct {
	// TODO: Define the arguments your module takes here
}

type Response struct {
	helpers.BaseResponse
	// TODO: Add additional response fields here
}

func main() {
	var response Response
	var moduleArgs ModuleArgs

	err := helpers.ParseArgs(os.Args, &moduleArgs)
	if err != nil {
		response.Msg = err.Error()
		helpers.ExitJSON(response, true)
	}

	// TODO: Implement business logic here
	response.Changed = false
	respose.Msg = "Did business logic"

	helpers.ExitJSON(response, false)
}
