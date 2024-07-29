// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !enterprise && !minimal

package command

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_extendAddonHandlers tests extendAddonHandlers() extends the minimal Vault handlers with handlers
// generated by newFullAddonHandlers()
func Test_extendAddonHandlers(t *testing.T) {
	handlers := newMinimalVaultHandlers()
	expMinPhysicalBackends := maps.Clone(handlers.physicalBackends)
	expMinLoginHandlers := maps.Clone(handlers.loginHandlers)

	expAddonPhysicalBackends, expAddonLoginHandlers := newFullAddonHandlers()

	extendAddonHandlers(handlers)

	require.Equal(t, len(expMinPhysicalBackends)+len(expAddonPhysicalBackends), len(handlers.physicalBackends),
		"extended total physical backends mismatch total of minimal and full addon physical backends")
	require.Equal(t, len(expMinLoginHandlers)+len(expAddonLoginHandlers), len(handlers.loginHandlers),
		"extended total login handlers mismatch total of minimal and full addon login handlers")

	for k := range expMinPhysicalBackends {
		require.Contains(t, handlers.physicalBackends, k, "expected to contain minimal physical backend")
	}

	for k := range expAddonPhysicalBackends {
		require.Contains(t, handlers.physicalBackends, k, "expected to contain full addon physical backend")
	}

	for k := range expMinLoginHandlers {
		require.Contains(t, handlers.loginHandlers, k, "expected to contain minimal login handler")
	}

	for k := range expAddonLoginHandlers {
		require.Contains(t, handlers.loginHandlers, k, "expected to contain full addon login handler")
	}
}
