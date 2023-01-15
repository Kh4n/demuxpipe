/*
Copyright © 2023 Asad-ullah Khan <asadullah@kh4n.io>
*/
package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "demuxpipe",
	Short: "Multiplexes or Demultiplexes incoming data on to one stream",
}

func initCobra() {
	if !verbose {
		log.SetFlags(0)
		log.SetOutput(io.Discard)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable all debug output")

	cobra.OnInitialize(initCobra)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
