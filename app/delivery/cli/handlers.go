package cli

import (
	"log"

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

func (p *cliHandlers) RootHandler(cmd *cobra.Command, args []string) {
	log.Print(args)
}
