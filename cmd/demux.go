/*
Copyright © 2023 Asad-ullah Khan <asadullah@kh4n.io>
*/
package cmd

import (
	"github.com/kh4n/demuxpipe/servers"
	"github.com/spf13/cobra"
)

func init() {
	var listenAddr string
	var writeAddr string
	// demuxCmd represents the demux command
	var demuxCmd = &cobra.Command{
		Use:   "demux",
		Short: "Demuxes incoming messages to another address",
		Run: func(cmd *cobra.Command, args []string) {
			servers.PipeToDemux(listenAddr, writeAddr)
		},
	}
	demuxCmd.Flags().StringVarP(&listenAddr, "listen", "l", ":8889", "address to receive mux")
	demuxCmd.Flags().StringVarP(&writeAddr, "write", "w", "", "address to demux to (eg. localhost proxy)")
	demuxCmd.MarkFlagRequired("write")

	rootCmd.AddCommand(demuxCmd)
}
