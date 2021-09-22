package main

import (
	"context"
	"log"
	"logging-service/internal/connector"
	"logging-service/internal/util"
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	version = "1.1.9"
)

var (
	cfg          = kingpin.Flag("cfg", "Configuration file path").Envar("CONFIGURATION_FILE_PATH").String()
	subscriber   = kingpin.Flag("sub-id", "Subscriber ID").Envar("SUB_ID").String()
	messageLimit = kingpin.Flag("message-limit", "Number of Messages").Default("10").Envar("NUMBER_OF_MESSAGES").Int()
	maxWorkers   = kingpin.Flag("max-workers", "Max workers").Default("3").Envar("MAX_PRODUCER_WORKERS").Int()
	keyfile      = kingpin.Flag("keyfile", "path of access token").Default("testpath").Envar("KEY_FILE").String()
)

func main() {
	kingpin.Version(version)
	kingpin.Parse()
	log.SetOutput(os.Stdout)

	pubsubConfig := connector.PubSubConfig{GcpProjectName: *project,
		SubscriberId:     *subscriber,
		NumberOfMessages: *messageLimit,
		KeyfileLocation:  *keyfile,
	}

	log.Println("Loading the config...")

	config, err := connector.NewConfig(*cfg)

	if err != nil {
		log.Fatalf("Error loading the config: %v", err)
	}

	stats := util.NewStats()
	go stats.PerSec()
	go stats.Monitor(10 * time.Second)

	message_chain := make(chan connector.LogMessage)

	go createWorkers(config, stats, message_chain)

	ctx := context.Background()
	/*ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(100 * time.Second)
		log.Println("sleep done!")
		cancel()
	}()

			go func() {
				for now := range time.Tick(time.Minute) {
					log.Printf("tick reach %v", now)
				}
			}()


		select {
		case <-ctx.Done():
			log.Println("context end")
		case <-time.Tick(200 * time.Second):
			log.Println("tick reach")
		}
	*/
	log.Println("Initializing Pub Sub Connector....")

	err = connector.NewPubSubProcessor(ctx, &pubsubConfig, config, stats, message_chain)
	if err != nil {
		log.Printf("Failed NewPubSubprocessor with error: %s", err)
		panic(err)
	}

}

func createWorkers(config *connector.Config, stats *util.Stats, message chan connector.LogMessage) {
	log.Println("Initializing connector....")
	for i := 0; i < *maxWorkers; i++ {
		producer, err := connector.NewBalabitClient(config.Balabit, config.Debug, stats)
		if err != nil {
			log.Printf("Failed Client with error: %s", err)
			panic(err)
		}
		go producer.PushEvents(message)
	}
	log.Println(" connector Up")
}
