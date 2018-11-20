package libs

import (
	"fmt"
	"strings"
)

type RootCommand struct {
	SubCommand map[string]Command
}

func (c *RootCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}
func (c *RootCommand) Run(cmdTree, args []string) {
	fmt.Println(strings.Join(cmdTree, " ") + ` <subcommand>

<subcommand>
  topics
  subscriptions`)
}
