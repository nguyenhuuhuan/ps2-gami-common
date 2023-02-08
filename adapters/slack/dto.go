package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Request request data
type Request struct {
	Text string `json:"text"`
}

// Response response data
type Response struct {
	Text string `json:"text"`
}

func encodeSlackRequest(_ context.Context, r *http.Request, req interface{}) error {
	request, _ := req.(*Request)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeSlackRequest(ctx context.Context, r *http.Response) (interface{}, error) {
	return r.Body, nil
}
