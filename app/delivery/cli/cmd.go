package cli

func SetupHandlers(h *cliHandlers) {
	h.rc.Run = h.RootHandler
	//должно быть h.rc.AddCommand(&cobra.Command{..., Run: h.SearchHandler}), например для команды поиска
}
