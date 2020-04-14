package cron

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/jamesroutley/unum/unumpb"
)

var (
	slotAvailablePreviously = false
)

func checkShiptonQueue(ctx context.Context) error {
	url := "https://www.shipton-mill.com/queue"
	rsp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	body := string(bodyBytes)

	slotAvailable := !strings.Contains(body, "Sorry, we don't have any delivery slots available at this time, please come back later.")

	msg := fmt.Sprintf("Slot available: %v", slotAvailable)
	log.Println(msg)

	if slotAvailable == slotAvailablePreviously {
		return nil
	}

	client := unumpb.NewUnumProtobufClient("http://localhost:8080", &http.Client{})
	_, err = client.SendSlackMessage(ctx, &unumpb.SendSlackMessageRequest{
		Text: msg,
	})
	if err != nil {
		return err
	}

	slotAvailablePreviously = slotAvailable

	return nil
}
