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
	var bindAddr string

	// muxCmd represents the mux command
	var muxCmd = &cobra.Command{
		Use:   "mux",
		Short: "Multiplex incoming connections to another address",
		Run: func(cmd *cobra.Command, args []string) {
			servers.MuxToPipe(listenAddr, writeAddr, bindAddr)
		},
	}
	muxCmd.Flags().StringVarP(&listenAddr, "listen", "l", ":8889", "address to listen on")
	muxCmd.Flags().StringVarP(&writeAddr, "write", "w", "", "address to mux to")
	muxCmd.MarkFlagRequired("write")
	muxCmd.Flags().StringVarP(&bindAddr, "bind", "b", ":9999", "address to send messages from")

	rootCmd.AddCommand(muxCmd)
}
