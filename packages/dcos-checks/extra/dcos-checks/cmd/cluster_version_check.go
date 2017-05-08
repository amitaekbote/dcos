// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

// clusterVersionCmd represents the check_cluster_version command
var clusterVersionCmd = &cobra.Command{
	Use:   "cluster_version",
	Short: "Check all nodes on the cluster are no more than 2 versions",
	Long: `A cluster could be running 2 versions during an upgrade. For every other
situation we expect that all nodes on the cluster be one version`,

	Run: func(cmd *cobra.Command, args []string) {
		RunCheck(NewClusterVersionCheck("DC/OS cluster version check"))
	},
}

func init() {
	RootCmd.AddCommand(clusterVersionCmd)
}

// ClusterVersionCheck validates the cluster has no more than 2 versions
type ClusterVersionCheck struct {
	Name string
}

// Run invokes a cluster version check and return error output, exit code and error.
func (c *ClusterVersionCheck) Run(ctx context.Context, cfg *CLIConfigFlags) (string, int, error) {
	// Get a list of all masters
	// Get a list of all agents
	// Get versions for each and throw in array?
	// Error if more than 2
	return "", 0, nil
}

// ID returns a unique check identifier.
func (c *ClusterVersionCheck) ID() string {
	return c.Name
}
