// Copyright 2024, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

// Package main is the entry point for the WaveTerm application.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/wavetermdev/waveterm/pkg/wavebase"
	"github.com/wavetermdev/waveterm/pkg/waveobj"
)

const WaveTermVersion = "0.1.0"

func main() {
	// Handle OS signals for graceful shutdown
	// Also handle SIGHUP so terminal closes don't leave orphaned processes
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Initialize the base application environment
	if err := wavebase.InitWaveBase(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to initialize waveterm base: %v\n", err)
		os.Exit(1)
	}

	// Initialize the object store
	if err := waveobj.InitObjectStore(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: failed to initialize object store: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("WaveTerm v%s starting...\n", WaveTermVersion)

	// Wait for shutdown signal
	go func() {
		sig := <-sigCh
		fmt.Printf("\nReceived signal %v, shutting down...\n", sig)
		shutdown()
		os.Exit(0)
	}()

	// Block main goroutine
	select {}
}

// shutdown performs a graceful shutdown of all application components.
func shutdown() {
	waveobj.CloseObjectStore()
	fmt.Println("WaveTerm shutdown complete.")
}
