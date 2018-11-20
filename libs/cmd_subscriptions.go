package libs

import (
	"fmt"
	"strings"
)

type SubscriptionsCommand struct {
	SubCommand map[string]Command
}

func (c *SubscriptionsCommand) GetSubCommands() map[string]Command {
	return c.SubCommand
}
func (c *SubscriptionsCommand) Run(cmdTree, args []string) {
	fmt.Println(strings.Join(cmdTree, " ") + ` <subcommand>

<subcommand>
  list
  consume`)

}
