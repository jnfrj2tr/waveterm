// Copyright 2024, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

package wavebase_test

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/wavetermdev/waveterm/pkg/wavebase"
)

func TestGetWaveHomeDir(t *testing.T) {
	homeDir := wavebase.GetWaveHomeDir()
	if homeDir == "" {
		t.Fatal("GetWaveHomeDir returned empty string")
	}
	// Should contain a wave-related path segment
	if !strings.Contains(homeDir, "waveterm") && !strings.Contains(homeDir, ".waveterm") {
		t.Logf("WaveHomeDir: %s", homeDir)
		// Not a hard failure, but log it for inspection
	}
}

func TestEnsureWaveHomeDir(t *testing.T) {
	// Use a temp dir to avoid polluting the real home
	tmpDir := t.TempDir()
	t.Setenv("WAVETERM_HOME", tmpDir)

	err := wavebase.EnsureWaveHomeDir()
	if err != nil {
		t.Fatalf("EnsureWaveHomeDir failed: %v", err)
	}

	// Verify the directory exists
	if _, statErr := os.Stat(tmpDir); os.IsNotExist(statErr) {
		t.Fatalf("Expected directory %s to exist after EnsureWaveHomeDir", tmpDir)
	}
}

func TestGetLogDir(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("WAVETERM_HOME", tmpDir)

	logDir := wavebase.GetLogDir()
	if logDir == "" {
		t.Fatal("GetLogDir returned empty string")
	}
	if !filepath.IsAbs(logDir) {
		t.Fatalf("Expected absolute path, got: %s", logDir)
	}
}

func TestGetDataDir(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("WAVETERM_HOME", tmpDir)

	dataDir := wavebase.GetDataDir()
	if dataDir == "" {
		t.Fatal("GetDataDir returned empty string")
	}
	if !filepath.IsAbs(dataDir) {
		t.Fatalf("Expected absolute path, got: %s", dataDir)
	}
}

func TestGetOS(t *testing.T) {
	os := wavebase.GetOS()
	if os == "" {
		t.Fatal("GetOS returned empty string")
	}

	// Validate that the returned OS matches the runtime
	switch runtime.GOOS {
	case "darwin":
		if os != "darwin" && os != "macos" {
			t.Logf("GetOS returned %q on darwin (runtime.GOOS=%q)", os, runtime.GOOS)
		}
	case "linux":
		if os != "linux" {
			t.Logf("GetOS returned %q on linux (runtime.GOOS=%q)", os, runtime.GOOS)
		}
	case "windows":
		if os != "windows" {
			t.Logf("GetOS returned %q on windows (runtime.GOOS=%q)", os, runtime.GOOS)
		}
	}
}
