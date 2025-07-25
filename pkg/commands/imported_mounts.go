// Code generated by go run tools/import_commands.go --talos-version v1.10.3 mounts
// DO NOT EDIT.

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"

	"github.com/siderolabs/talos/pkg/cli"
	"github.com/siderolabs/talos/pkg/machinery/client"
	"github.com/siderolabs/talos/pkg/machinery/formatters"
)

// mountsCmd represents the mounts command.
var mountsCmd = &cobra.Command{
	Use:     "mounts",
	Aliases: []string{"mount"},
	Short:   "List mounts",
	Long:    ``,
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return WithClient(func(ctx context.Context, c *client.Client) error {
			var remotePeer peer.Peer

			resp, err := c.Mounts(ctx, grpc.Peer(&remotePeer))
			if err != nil {
				if resp == nil {
					return fmt.Errorf("error getting mount information: %s", err)
				}

				cli.Warning("%s", err)
			}

			return formatters.RenderMounts(resp, os.Stdout, &remotePeer)
		})
	},
}

func init() {
	mountsCmd.Flags().StringSliceVarP(&mountsCmdFlags.configFiles,
		"file", "f", nil, "specify config files or patches in a YAML file (can specify multiple)",
	)
	mountsCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		nodesFromArgs :=
			len(GlobalArgs.Nodes) > 0
		endpointsFromArgs := len(
			GlobalArgs.Endpoints) > 0
		for _, configFile := range mountsCmdFlags.configFiles {
			if err := processModelineAndUpdateGlobals(configFile, nodesFromArgs, endpointsFromArgs, false); err != nil {
				return err
			}
		}
		return nil
	}

	addCommand(mountsCmd)
}

var mountsCmdFlags struct {
	configFiles []string
}
