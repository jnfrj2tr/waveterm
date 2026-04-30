// Copyright 2024, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

// Package wavebase provides core utilities and constants for the WaveTerm application.
package wavebase

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	AppName        = "waveterm"
	AppVersion     = "0.1.0"
	WaveHomeDirEnv = "WAVETERM_HOME"
	DefaultHomeDir = ".waveterm"
	DefaultLogDir  = "log"
	DefaultDataDir = "data"
	DefaultTempDir = "tmp"
	DefaultConfigDir = "config"
)

var (
	waveHomeDir     string
	waveHomeDirOnce sync.Once
)

// GetWaveHomeDir returns the root directory for WaveTerm data.
// It respects the WAVETERM_HOME environment variable, falling back
// to ~/.waveterm on Unix-like systems.
func GetWaveHomeDir() string {
	waveHomeDirOnce.Do(func() {
		if envHome := os.Getenv(WaveHomeDirEnv); envHome != "" {
			waveHomeDir = envHome
			return
		}
		userHome, err := os.UserHomeDir()
		if err != nil {
			// Fallback to current directory if home cannot be determined
			userHome = "."
		}
		waveHomeDir = filepath.Join(userHome, DefaultHomeDir)
	})
	return waveHomeDir
}

// EnsureWaveHomeDir creates the WaveTerm home directory and required
// subdirectories if they do not already exist.
func EnsureWaveHomeDir() error {
	homeDir := GetWaveHomeDir()
	subDirs := []string{
		homeDir,
		filepath.Join(homeDir, DefaultLogDir),
		filepath.Join(homeDir, DefaultDataDir),
		filepath.Join(homeDir, DefaultTempDir),
		filepath.Join(homeDir, DefaultConfigDir),
	}
	for _, dir := range subDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("wavebase: failed to create directory %q: %w", dir, err)
		}
	}
	return nil
}

// GetLogDir returns the path to the log directory.
func GetLogDir() string {
	return filepath.Join(GetWaveHomeDir(), DefaultLogDir)
}

// GetDataDir returns the path to the data directory.
func GetDataDir() string {
	return filepath.Join(GetWaveHomeDir(), DefaultDataDir)
}

// GetTempDir returns the path to the temp directory.
func GetTempDir() string {
	return filepath.Join(GetWaveHomeDir(), DefaultTempDir)
}

// GetConfigDir returns the path to the config directory.
func GetConfigDir() string {
	return filepath.Join(GetWaveHomeDir(), DefaultConfigDir)
}

// GetOS returns a normalized OS identifier string.
func GetOS() string {
	return runtime.GOOS
}

// GetArch returns the current CPU architecture.
func GetArch() string {
	return runtime.GOARCH
}

// VersionString returns a human-readable version string including
// the application name, version, OS, and architecture.
func VersionString() string {
	return fmt.Sprintf("%s v%s (%s/%s)", AppName, AppVersion, GetOS(), GetArch())
}
