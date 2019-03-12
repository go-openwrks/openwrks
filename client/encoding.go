package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-openwrks/openwrks/response"
)

func TransformResponse(src *http.Response, dst interface{}) error {
	defer src.Body.Close()

	d, err := ioutil.ReadAll(src.Body)
	if err != nil {
		return err
	}

	// If the status code is that of an error, return an error.
	if src.StatusCode > 399 {
		e := &response.Error{}
		json.Unmarshal(d, e)
		return e
	}

	return json.Unmarshal(d, dst)
}
