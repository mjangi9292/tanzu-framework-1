// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestGenerateDescriptor(t *testing.T) {
	m := MainUsage{}

	f := m.Func()

	c := &cobra.Command{
		Use:   "tanzu",
		Short: aurora.Bold(`Tanzu CLI`).String(),
	}

	err := f(c)
	require.NoError(t, err)
}
