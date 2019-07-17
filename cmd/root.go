package cmd

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/jeffguorg/blog.jeffthecoder.space/logging"
	"github.com/spf13/cobra"
)

var (
	verbosity *string
)

var rootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "A simple blog daemon based on local git repository",
	PersistentPreRunE: func(*cobra.Command, []string) error {
		err := logging.SetupLevel(*verbosity)
		logging.SetupFormat()
		return err
	},
}

// Execute the defined subcommands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logging.Fatal(err)
	}
}

func init() {
	verbosity = rootCmd.PersistentFlags().String("verbosity", logrus.WarnLevel.String(), "logging level")
}
