package cmd

import (
	"github.com/jeffguorg/blog.jeffthecoder.space/web"
	"github.com/spf13/cobra"
)

var (
	addr    *string
	root    *string
	metrics *bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A simple blog daemon based on local git repository",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return web.Serve(*addr, *root, *metrics)
	},
}

// Execute the defined subcommands
func init() {
	rootCmd.AddCommand(serveCmd)

	addr = serveCmd.Flags().String("addr", ":80", "the address which it will bind to")
	root = serveCmd.Flags().String("root", "/blog", "web root location.")
	metrics = serveCmd.Flags().Bool("metrics", false, "mount prometheus metric route")
}
