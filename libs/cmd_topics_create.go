package libs

import "fmt"

type TopicsCreateCommand struct {
	SubCommand map[string]Command
}

func (c *TopicsCreateCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *TopicsCreateCommand) Run(cmdTree, args []string) {
	fmt.Println("Run TopicsCreateCommand", args)
	// Show usage
	panic("Not yet implemented")
}
