package app

import "github.com/spf13/cobra"

// ConsoleDelivery
type ConsoleDelivery interface {
	Cmd() func(cmd *cobra.Command, args []string)
}
