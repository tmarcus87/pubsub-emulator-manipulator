package libs

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/pubsub"
)

type TopicsCreateCommand struct {
	SubCommand map[string]Command
}

func (c *TopicsCreateCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *TopicsCreateCommand) Run(cmdTree, args []string) {
	fs := flag.NewFlagSet("topics create", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	pProjectID := fs.String("project", "", "ProjectID")
	pTopicName := fs.String("topic", "", "Topic name")
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

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, *pProjectID)
	if err != nil {
		panic(err)
	}
	topic, err := client.CreateTopic(ctx, *pTopicName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Topic(%s) is created.\n", topic.ID())
}
