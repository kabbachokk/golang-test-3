package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	// здесь же можно добавлять команды приложения, типа server, ...
	return &cobra.Command{
		Use:   "",
		Short: "",
		Long:  ``,
		Run:   nil,
	}
}
