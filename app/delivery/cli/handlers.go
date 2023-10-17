package cli

import (
	"github.com/spf13/cobra"
	"ip.com/app"
)

type cliHandlers struct {
	uc app.UseCase
	rc *cobra.Command // rootCmd
}

func NewCliHandlers(
	uc app.UseCase,
	rc *cobra.Command,
) *cliHandlers {
	return &cliHandlers{uc, rc}
}

func (p *cliHandlers) Handler(cmd *cobra.Command, args []string) {}
