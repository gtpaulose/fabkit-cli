/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package network

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
func newNetworkStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start a Hyperledger Fabric Network",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("start called")
		},
	}
}
