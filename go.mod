module github.com/anthropics/cc-connect

go 1.21

require (
	github.com/gorilla/websocket v1.5.1
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.2
	go.uber.org/zap v1.26.0
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20231108232855-2478ac86f678 // indirect
	golang.org/x/net v0.23.0 // indirect; bumped from v0.20.0 to pick up CVE-2023-45288 fix
	golang.org/x/sys v0.18.0 // indirect; bumped alongside x/net
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Personal fork of chenhg5/cc-connect for local experimentation.
// Upstream: https://github.com/chenhg5/cc-connect
// Note: keeping go version pinned at 1.21 intentionally; upgrading to 1.22+
// changes loop variable semantics which may affect range-based goroutine usage
// in this codebase. Revisit before bumping.
//
// TODO: audit range-loop goroutine captures in pkg/hub before attempting 1.22
// upgrade — at least two spots in hub.go close over the loop variable 'client'.
