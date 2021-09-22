package connector

import (
	"context"
	"fmt"
	"log"
	"logging-service/internal/util"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

var (
	converters = []PubSubEventConverter{RegexConverter{}.ConvertUAAEvent}
)

func PublishEvents(event LogMessage, message chan LogMessage) {
	message <- event
}

func NewPubSubProcessor(ctx context.Context, config *PubSubConfig, appConfig *Config, stats *util.Stats, message chan LogMessage) error {
	projectID := config.GcpProjectName
	subscriptionID := config.SubscriberId
	NumberOfMessages := config.NumberOfMessages
	//ctx := context.Background()
	//var jsonPath string = "/Users/jtsang/logging/build/packages/testuserKeyfile.json" // config.KeyfileLocation
	var jsonPath string = config.KeyfileLocation
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(jsonPath))
	//client, err := pubsub.NewClient(ctx, projectID)

	if err != nil {
		log.Printf("Could not create pubsub Client: %v", err)
		return err
	}

	subName := fmt.Sprintf("projects/%s/subscriptions/%s", projectID, subscriptionID)
	log.Printf("Sub %q\n", subName)
	sub := client.Subscription(subscriptionID)

	status, err := sub.Exists(ctx)

	if err != nil {
		log.Printf("Could not get subscription status: %v", err)
		return err
	}

	log.Printf("sub created: %v", status)

	//log.Printf("Sub created")
	sub.ReceiveSettings.MaxExtension = 30 * time.Second
	sub.ReceiveSettings.MaxOutstandingMessages = NumberOfMessages
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		//log.Println("it reach to the cycle")
		stats.Inc(util.Consume)
		for _, converter := range converters {
			event, err := converter(string(m.Data))
			if err == nil && event.Message != "" {
				event.PubSubMessage = m
				go PublishEvents(event, message)
			} else {
				m.Ack()
				if err != nil {
					log.Printf("ERROR: Log Message is unprocessable. Failed to convert message with error: %s\nLog Content Length: %d", err, len(string(m.Data)))
				}
			}
		}
	})
	log.Println("cycle finished")
	if err != context.Canceled {
		log.Printf("Could not create processor with error : %v", err)
		return err
	}

	return err

}
