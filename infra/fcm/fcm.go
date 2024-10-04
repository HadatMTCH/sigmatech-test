package fcm

import (
	"base-api/config"
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	objects "base-api/objects/notification"
)

type fcm struct {
	cfg *config.FCM
}

type FCMInterface interface {
	SendNotification(data objects.SendNotification)
}

func NewFCM(cfg *config.FCM) FCMInterface {
	return &fcm{
		cfg: cfg,
	}
}

func send(ctx context.Context, client *messaging.Client, data objects.SendNotification) {
	message := messaging.MulticastMessage{
		Tokens: data.TargetTokens,
		Notification: &messaging.Notification{
			Title: data.Title,
			Body:  data.Body,
		},
	}

	br, err := client.SendMulticast(ctx, &message)
	if err != nil {
		log.Errorln(err)
		return
	}

	if br.FailureCount > 0 {
		log.Error(fmt.Sprintf("failed sending %d message", br.FailureCount))
		for _, response := range br.Responses {
			log.Errorln(response)
			return
		}
	}
}

func (f fcm) SendNotification(data objects.SendNotification) {
	ctx := context.Background()

	opt := option.WithCredentialsFile(f.cfg.KeyFileDir)
	config := &firebase.Config{
		ProjectID: f.cfg.ProjectID,
	}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Errorf("error initializing app: %s\n", err.Error())
		return
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Errorf("error getting Messaging client: %s\n", err.Error())
		return
	}

	send(ctx, client, data)
}
