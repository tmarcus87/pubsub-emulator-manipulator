package libs

import "fmt"

type TopicsSubscriptionCommand struct {
	SubCommand map[string]Command
}

func (c *TopicsSubscriptionCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *TopicsSubscriptionCommand) Run(cmdTree, args []string) {
	fmt.Println("Run TopicsSubscriptionCommand", args)
	// Show usage
	panic("Not yet implemented")
}
