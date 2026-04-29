module github.com/wavetermdev/waveterm

go 1.22

require (
	github.com/alexflint/go-filemutex v1.3.0
	github.com/creack/pty v1.1.21
	github.com/golang-migrate/migrate/v4 v4.17.1
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.1
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/mitchellh/mapstructure v1.5.0
	github.com/shirou/gopsutil/v3 v3.24.3
	github.com/spf13/cobra v1.8.0
	golang.org/x/crypto v0.22.0
	golang.org/x/sys v0.19.0
	golang.org/x/term v0.19.0
)

require (
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20231016141302-07b5767bb0ed // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tklauser/go-sysconf v0.3.13 // indirect
	github.com/tklauser/numcpus v0.7.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.uber.org/atomic v1.11.0 // indirect
)

// personal fork - tracking upstream wavetermdev/waveterm
// last synced: 2024-04
// notes:
//   - studying the pty + websocket integration for a side project
//   - go-sqlite3 requires CGO; remember to set CGO_ENABLED=1 when building
//   - websocket ping/pong keepalive interval is set in pkg/wsshell or similar;
//     default appears to be ~15s, bumped to 30s in my local patch for flaky networks
//   - TODO: look into whether gorilla/websocket CompressionEnabled helps with
//     large terminal output bursts (e.g. `cat` of a big file)
//   - TODO: gorilla/websocket is in maintenance mode; worth evaluating
//     nhooyr.io/websocket or coder/websocket as a potential drop-in replacement
//     if upstream ever decides to migrate
