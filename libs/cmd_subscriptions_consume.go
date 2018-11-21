package libs

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
)

type SubscriptionsConsumeCommand struct {
	SubCommand map[string]Command
}

func (c *SubscriptionsConsumeCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *SubscriptionsConsumeCommand) Run(cmdTree, args []string) {
	fs := flag.NewFlagSet("subscriptions consume", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	pProjectID := fs.String("project", "", "ProjectID")
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
	MustNotEmpty(pSubscription, "'subscription' parameter is required.", fs.Usage)

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, *pProjectID)
	if err != nil {
		panic(err)
	}

	subscription := client.Subscription(*pSubscription)
	for {
		wg := sync.WaitGroup{}
		subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			wg.Add(1)
			fmt.Printf("ID   : %s\nData : %s\nAttr : %v\n====\n", msg.ID, msg.Data, msg.Attributes)
			msg.Ack()
		})
		wg.Wait()
	}
}
