package libs

import "fmt"

type SubscriptionsConsumeCommand struct {
	SubCommand map[string]Command
}

func (c *SubscriptionsConsumeCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *SubscriptionsConsumeCommand) Run(cmdTree, args []string) {
	fmt.Println("Run SubscriptionsConsumeCommand", args)
	// Show usage
	panic("Not yet implemented")
}
