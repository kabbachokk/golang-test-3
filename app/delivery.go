package app

import "github.com/spf13/cobra"

// ConsoleDelivery
type ConsoleDelivery interface {
	RootHandler() func(cmd *cobra.Command, args []string)
}
