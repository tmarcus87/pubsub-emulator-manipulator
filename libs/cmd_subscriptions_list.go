package libs

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/pubsub"
)

type SubscriptionsListCommand struct {
	SubCommand map[string]Command
}

func (c *SubscriptionsListCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *SubscriptionsListCommand) Run(cmdTree, args []string) {
	fs := flag.NewFlagSet("subscriptions list", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	pProjectID := fs.String("project", "", "ProjectID")
	pHelp := fs.Bool("help", false, "Display help")
	fs.Parse(args)
	fs.SetOutput(os.Stderr)

	if *pHelp {
		fs.SetOutput(os.Stdout)
		fs.Usage()
		return
	}

	MustNotEmpty(pProjectID, "'project' parameter is required.", fs.Usage)

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, *pProjectID)
	if err != nil {
		panic(err)
	}

	it := client.Subscriptions(ctx)
	for {
		subscription, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			panic(err)
		}
		fmt.Printf("* %s (%s)\n", subscription.ID(), subscription.String())
	}
}
