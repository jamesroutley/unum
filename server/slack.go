package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jamesroutley/unum/unumpb"
)

func (s Server) SendSlackMessage(ctx context.Context, req *unumpb.SendSlackMessageRequest) (*unumpb.SendSlackMessageResponse, error) {

	body := map[string]interface{}{
		"text": req.Text,
	}
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	log.Println(string(bodyStr))

	slackRequest, err := http.NewRequest("POST", s.Config.SlackURL, bytes.NewBuffer(bodyStr))
	if err != nil {
		return nil, err
	}
	slackRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(slackRequest)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	log.Println("response Body:", string(rspBody))

	return &unumpb.SendSlackMessageResponse{}, nil
}
