package libs

import (
	"os"
)

type Command interface {
	GetSubCommands() map[string]Command
	Run(cmdTree, args []string)
}

func NewCommand() Command {
	return &RootCommand{
		SubCommand: map[string]Command{
			"topics": &TopicsCommand{
				SubCommand: map[string]Command{
					"list":         &TopicsListCommand{},
					"create":       &TopicsCreateCommand{},
					"subscription": &TopicsSubscriptionCommand{},
					"publish":      &TopicsPublishCommand{},
				},
			},
			"subscriptions": &SubscriptionsCommand{
				SubCommand: map[string]Command{
					"list":    &SubscriptionsListCommand{},
					"consume": &SubscriptionsConsumeCommand{},
				},
			},
		},
	}
}

func Execute(cmd Command) {
	var (
		targetCmd = NewCommand()
		cmdTree   = make([]string, 0)
		subArgs   = make([]string, 0)
	)

	cmdTree = append(cmdTree, os.Args[0])

	args := os.Args[1:]

COMMAND_PARSER:
	for i := 0; i < len(args); i++ {
		cmdTree = append(cmdTree, args[i])

		subCmdMap := targetCmd.GetSubCommands()

		if newCmd, ok := subCmdMap[args[i]]; ok {
			targetCmd = newCmd
		} else {
			for j := 0; j < i; j++ {
				cmdTree = append(cmdTree, args[i])
			}
			for j := i; j < len(args); j++ {
				subArgs = append(subArgs, args[j])
			}
			break COMMAND_PARSER
		}
	}

	targetCmd.Run(cmdTree, subArgs)

}
