package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/jamesroutley/unum/unumpb"
)

func (s Server) SendSMS(ctx context.Context, req *unumpb.SendSMSRequest) (*unumpb.SendSMSResponse, error) {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + s.Config.TwilioAccountSID + "/Messages.json"

	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", s.Config.PhoneNumber)
	msgData.Set("From", "+18507903914")
	msgData.Set("Body", req.Body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	twilioReq, err := http.NewRequest("POST", urlStr, &msgDataReader)
	if err != nil {
		return nil, err
	}
	twilioReq.SetBasicAuth(s.Config.TwilioAccountSID, s.Config.TwilioAuthToken)
	twilioReq.Header.Add("Accept", "application/json")
	twilioReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(twilioReq)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data)
		}
		fmt.Println(resp)
	}

	return &unumpb.SendSMSResponse{}, nil
}
