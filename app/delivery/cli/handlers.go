package cli

import (
	"ip.com/app"
)

type cliHandlers struct {
	uc app.UseCase
}

func NewCliHandlers(
	uc app.UseCase,
) *cliHandlers {
	return &cliHandlers{uc}
}
