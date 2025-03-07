package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/odpf/salt/cmdx"

	"github.com/spf13/cobra"
)

// Execute runs the command line interface
func Execute() {
	rootCmd := &cobra.Command{
		Use:           "siren <command> <subcommand> [flags]",
		Short:         "siren",
		Long:          "Work seemlessly with your observability stack.",
		SilenceUsage:  true,
		SilenceErrors: true,
		Annotations: map[string]string{
			"group:core": "true",
			"help:learn": heredoc.Doc(`
				Use 'siren <command> <subcommand> --help' for more information about a command.
				Read the manual at https://odpf.gitbook.io/siren/
			`),
			"help:feedback": heredoc.Doc(`
				Open an issue here https://github.com/odpf/siren/issues
			`),
		},
	}

	cliConfig, err := readConfig()
	if err != nil {
		fmt.Println(err)
	}

	cmdx.SetHelp(rootCmd)

	rootCmd.AddCommand(configCmd())
	rootCmd.AddCommand(serveCmd())
	rootCmd.AddCommand(migrateCmd())
	rootCmd.AddCommand(providersCmd(cliConfig))
	rootCmd.AddCommand(namespacesCmd(cliConfig))
	rootCmd.AddCommand(receiversCmd(cliConfig))
	rootCmd.AddCommand(templatesCmd(cliConfig))
	rootCmd.AddCommand(rulesCmd(cliConfig))
	rootCmd.AddCommand(alertsCmd(cliConfig))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
