var rootCmd = &cobra.Command{
	Use: 'hugo',
	Short: "Hugo is a very fast static site generator",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}