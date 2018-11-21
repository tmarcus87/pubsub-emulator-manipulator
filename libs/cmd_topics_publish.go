package libs

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
)

type TopicsPublishCommand struct {
	SubCommand map[string]Command
}

func (c *TopicsPublishCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *TopicsPublishCommand) Run(cmdTree, args []string) {
	fs := flag.NewFlagSet("topics publish", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	pProjectID := fs.String("project", "", "ProjectID")
	pTopicName := fs.String("topic", "", "Topic name")
	pData := fs.String("data", "", "Message data")
	pAttribute := fs.String("attribute", "", "Message attribute(KEY1=VAL1,KEY2=VAL2")
	pHelp := fs.Bool("help", false, "Display help")
	fs.Parse(args)
	fs.SetOutput(os.Stderr)

	if *pHelp {
		fs.SetOutput(os.Stdout)
		fs.Usage()
		return
	}

	attributes := make(map[string]string)
	if *pAttribute != "" {
		for _, v := range strings.Split(*pAttribute, ",") {
			e := strings.Split(v, "=")
			attributes[e[0]] = e[1]
		}
	}

	MustNotEmpty(pProjectID, "'project' parameter is required.", fs.Usage)
	MustNotEmpty(pTopicName, "'topic' parameter is required.", fs.Usage)
	MustNotEmpty(pData, "'data' parameter is required.", fs.Usage)

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, *pProjectID)
	if err != nil {
		panic(err)
	}
	topic := client.Topic(*pTopicName)
	serverID, err :=
		topic.Publish(ctx,
			&pubsub.Message{
				ID:         time.Now().String(),
				Data:       []byte(*pData),
				Attributes: attributes,
			}).Get(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message(%s) published to %s.\n", serverID, *pTopicName)
}
