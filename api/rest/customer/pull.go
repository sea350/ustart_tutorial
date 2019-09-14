package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

// Pull wraps backend/customer/pull.go
func (rapi *RESTAPI) Pull(w http.ResponseWriter, req *http.Request) {

	//NOTE: this method of retrieving data from a REST request should only be used for GET requests
	//later on you will be shown the difference between what is and when you should use GET or POST
	req.ParseForm()
	uuid := req.Form.Get("uuid")

	lookReq := &customerpb.PullRequest{
		UUID: uuid,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.prof.Pull(req.Context(), lookReq)
	if resp != nil {
		ret["response"] = resp
	} else {
		ret["response"] = ""
	}
	if err != nil {
		ret["error"] = err.Error()
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
