package libs

import "fmt"

type TopicsPublishCommand struct {
	SubCommand map[string]Command
}

func (c *TopicsPublishCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}

func (c *TopicsPublishCommand) Run(cmdTree, args []string) {
	fmt.Println("Run TopicsPublishCommand", args)
	// Show usage
	panic("Not yet implemented")
}
