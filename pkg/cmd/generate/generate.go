/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package generate

import (
	"context"
	"fmt"
	"log"

	"github.com/czar0/fabkit-cli/internal/docker"
	"github.com/czar0/fabkit-cli/internal/spinner"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// NewGenerateCmd represents the generate command
func NewGenerateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate cryptos and aritfacts required to start the network",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				log.Fatalln(err)
			}

			if err := docker.CheckServerRunning(); err != nil {
				log.Fatalln(err)
			}

			spinner.Spin.Start()
			spinner.Spin.Suffix = fmt.Sprintf(" Generating cryptos")

			resp, err := cli.ContainerCreate(ctx, &container.Config{
				Image: "hyperledger/fabric-tools:2.3.3",
				Cmd:   []string{"cryptogen", "generate", "--config=/crypto-config.yaml", "--output=/crypto-config"},
				Tty:   false,
			}, &container.HostConfig{
				Mounts: []mount.Mount{
					{
						Type:   mount.TypeBind,
						Source: "/Users/georgep/Projects/fabkit-cli/network/config/crypto-config.yaml",
						Target: "/crypto-config.yaml",
					},
					{
						Type:   mount.TypeBind,
						Source: "/Users/georgep/Projects/fabkit-cli/network/crypto-config",
						Target: "/crypto-config",
					},
				},
			}, nil, nil, "")
			if err != nil {
				log.Fatalln(err)
			}

			if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
				log.Fatalln(err)
			}

			statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
			select {
			case err := <-errCh:
				if err != nil {
					log.Fatalln(err)
				}
			case <-statusCh:
			}

			spinner.Spin.Stop()
		},
	}
}
