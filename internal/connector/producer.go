package connector

import (
	"cap-logging-service/internal/util"
	//	"crypto/tls"
	//"crypto/x509"
	"fmt"
	//"io/ioutil"
	"log"

	"cloud.google.com/go/pubsub"

	syslog "github.com/RackSec/srslog"
)

type BalabitClient struct {
	config *BalabitConfig
	writer *syslog.Writer
	stats  *util.Stats
	debug  bool
}

type LogMessage struct {
	PubSubMessage *pubsub.Message
	Message       string
	Priority      syslog.Priority
}

func NewBalabitClient(config *BalabitConfig, debug bool, stats *util.Stats) (*BalabitClient, error) {
	var writer *syslog.Writer
	log.Printf("Connecting worker to server : %s", config.EndpointAddress)
	return &BalabitClient{
		config: config,
		writer: writer,
		stats:  stats,
		debug:  debug,
	}, nil
}

func (client *BalabitClient) PushEvents(events chan LogMessage) {
	for event := range events {
		client.stats.Inc(util.Produce)
		client.PushMessage(event)
	}
}

func (client *BalabitClient) PushMessage(message LogMessage) {
	fmt.Println(message.Message)
	message.PubSubMessage.Ack()
}
