package libs

import "fmt"

type SubscriptionsListCommand struct {
	SubCommand map[string]Command
}

func (c *SubscriptionsListCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *SubscriptionsListCommand) Run(cmdTree, args []string) {
	fmt.Println("Run SubscriptionsListCommand", args)
	// Show usage
	panic("Not yet implemented")
}
