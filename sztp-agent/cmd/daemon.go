/*
SPDX-License-Identifier: Apache-2.0
Copyright (C) 2022 Intel Corporation
Copyright (c) 2022 Dell Inc, or its subsidiaries.
Copyright (C) 2022 Red Hat.
*/

package cmd

import (
	"github.com/opiproject/sztp/sztp-agent/pkg/secureAgent"
	"github.com/spf13/cobra"
)

func NewDaemonCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "daemon",
		Short: "Run the daemon command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return secureAgent.RunCommandDaemon()
		},
	}
	return cmd
}