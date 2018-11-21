package libs

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/pubsub"
)

type TopicsSubscriptionCommand struct {
	SubCommand map[string]Command
}

func (c *TopicsSubscriptionCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *TopicsSubscriptionCommand) Run(cmdTree, args []string) {
	fs := flag.NewFlagSet("topics create", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	pProjectID := fs.String("project", "", "ProjectID")
	pTopicName := fs.String("topic", "", "Topic name")
	pSubscription := fs.String("subscription", "", "Subscription name")
	pHelp := fs.Bool("help", false, "Display help")
	fs.Parse(args)
	fs.SetOutput(os.Stderr)

	if *pHelp {
		fs.SetOutput(os.Stdout)
		fs.Usage()
		return
	}

	MustNotEmpty(pProjectID, "'project' parameter is required.", fs.Usage)
	MustNotEmpty(pTopicName, "'topic' parameter is required.", fs.Usage)
	MustNotEmpty(pSubscription, "'subscription' parameter is required.", fs.Usage)

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, *pProjectID)
	if err != nil {
		panic(err)
	}
	topic := client.Topic(*pTopicName)
	subscription, err :=
		client.CreateSubscription(ctx, *pSubscription, pubsub.SubscriptionConfig{Topic: topic})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Subscription(%s) for %s is created.\n", subscription.ID(), topic.ID())
}
