package cron

import (
	"context"
	"log"
	"time"
)

func Init() error {
	register(5*time.Minute, "check shipton queue", checkShiptonQueue)
	return nil
}

func register(
	duration time.Duration, name string, f func(context.Context) error,
) {
	go func() {
		ticker := time.NewTicker(duration)
		for {
			<-ticker.C
			log.Printf("Running cron '%s'", name)
			err := f(context.Background())
			if err != nil {
				log.Println(err)
			}
		}
	}()
}
