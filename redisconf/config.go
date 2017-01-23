package redisconf

import (
	"fmt"
	"strings"
)

//Config is a config value for a `redis.conf`
type Config struct {
	Name  string
	Value string
}

//NewConfig is a conventient way to initialise a Config
func NewConfig(name, value string) Config {
	return Config{Name: name, Value: value}
}

func (config Config) String() string {
	return fmt.Sprintf("%s %s", config.Name, config.Value)
}

//Validate returns an error if it considered invalid by Redis
func (config Config) Validate() error {
	if !config.hasKnownConfigName() {
		return fmt.Errorf("unknown config: `%s`", config.Name)
	}
	return nil
}

//DecodeConfig attempts to initialise a config from a raw `redis.conf` line
func DecodeConfig(rawDirective string) (Config, error) {
	config := fromRawDirective(rawDirective)

	if err := config.Validate(); err != nil {
		return Config{}, err
	}

	return config, nil
}

func fromRawDirective(rawDirective string) Config {
	words := strings.SplitN(strings.TrimSpace(rawDirective), " ", 2)

	switch len(words) {
	case 0:
		return Config{"", ""}
	case 1:
		return Config{words[0], ""}
	default:
		return Config{words[0], words[1]}
	}
}

func (config Config) hasKnownConfigName() bool {
	for _, validConfig := range validConfigs {
		if config.Name == validConfig {
			return true
		}
	}
	return false
}

var validConfigs = []string{
	"dbfilename",
	"requirepass",
	"masterauth",
	"unixsocket",
	"logfile",
	"pidfile",
	"slave-announce-ip",
	"maxmemory",
	"maxmemory-samples",
	"timeout",
	"auto-aof-rewrite-percentage",
	"auto-aof-rewrite-min-size",
	"hash-max-ziplist-entries",
	"hash-max-ziplist-value",
	"list-max-ziplist-size",
	"list-compress-depth",
	"set-max-intset-entries",
	"zset-max-ziplist-entries",
	"zset-max-ziplist-value",
	"hll-sparse-max-bytes",
	"lua-time-limit",
	"slowlog-log-slower-than",
	"latency-monitor-threshold",
	"slowlog-max-len",
	"port",
	"tcp-backlog",
	"databases",
	"repl-ping-slave-period",
	"repl-timeout",
	"repl-backlog-size",
	"repl-backlog-ttl",
	"maxclients",
	"watchdog-period",
	"slave-priority",
	"slave-announce-port",
	"min-slaves-to-write",
	"min-slaves-max-lag",
	"hz",
	"cluster-node-timeout",
	"cluster-migration-barrier",
	"cluster-slave-validity-factor",
	"repl-diskless-sync-delay",
	"tcp-keepalive",
	"cluster-require-full-coverage",
	"no-appendfsync-on-rewrite",
	"slave-serve-stale-data",
	"slave-read-only",
	"stop-writes-on-bgsave-error",
	"daemonize",
	"rdbcompression",
	"rdbchecksum",
	"activerehashing",
	"protected-mode",
	"repl-disable-tcp-nodelay",
	"repl-diskless-sync",
	"aof-rewrite-incremental-fsync",
	"aof-load-truncated",
	"maxmemory-policy",
	"loglevel",
	"supervised",
	"appendfsync",
	"syslog-facility",
	"appendonly",
	"dir",
	"save",
	"client-output-buffer-limit",
	"unixsocketperm",
	"slaveof",
	"notify-keyspace-events",
	"bind",
}
