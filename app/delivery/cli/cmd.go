package cli

func SetupHandlers(h *cliHandlers) {
	h.rc.Run = h.Handler
	//должно быть h.rc.AddCommand(cobra.Command{..., Run: h.Handler}) если не rootCmd
}
