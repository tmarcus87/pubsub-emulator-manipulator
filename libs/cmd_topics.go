package libs

import (
	"fmt"
	"strings"
)

type TopicsCommand struct {
	SubCommand map[string]Command
}

func (c *TopicsCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}
func (c *TopicsCommand) Run(cmdTree, args []string) {
	fmt.Println(strings.Join(cmdTree, " ") + ` <subcommand>

<subcommand>
  list
  create
  subscription
  publish`)

}
