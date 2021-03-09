package main

import (
	"github.com/urfave/cli/v2"
)

func CommandRun() *cli.Command {
	return &cli.Command{
		Name: "run",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "add-host",
				Usage: "Add a custom host-to-IP mapping (host:ip)",
			},
			&cli.StringFlag{
				Name:  "attach",
				Usage: "Attach to STDIN, STDOUT or STDERR",
				Aliases: []string{
					"a",
				},
			},
			&cli.StringFlag{
				Name:  "blkio-weight",
				Usage: "Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)",
			},
			&cli.StringFlag{
				Name:  "blkio-weight-device",
				Usage: "Block IO weight (relative device weight)",
			},
			&cli.StringFlag{
				Name:  "cap-add",
				Usage: "Add Linux capabilities",
			},
			&cli.StringFlag{
				Name:  "cap-drop",
				Usage: "Drop Linux capabilities",
			},
			&cli.StringFlag{
				Name:  "cgroup-parent",
				Usage: "Optional parent cgroup for the container",
			},
			&cli.StringFlag{
				Name:  "cgroupns",
				Usage: "API 1.41+Cgroup namespace to use (host|private)\\n'host':    Run the container in the Docker host's cgroup namespace\\n'private': Run the container in its own private cgroup namespace\\n'':        Use the cgroup namespace as configured by the\\n           default-cgroupns-mode option on the daemon (default)",
			},
			&cli.StringFlag{
				Name:  "cidfile",
				Usage: "Write the container ID to the file",
			},
			&cli.StringFlag{
				Name:  "cpu-count",
				Usage: "CPU count (Windows only)",
			},
			&cli.StringFlag{
				Name:  "cpu-percent",
				Usage: "CPU percent (Windows only)",
			},
			&cli.StringFlag{
				Name:  "cpu-period",
				Usage: "Limit CPU CFS (Completely Fair Scheduler) period",
			},
			&cli.StringFlag{
				Name:  "cpu-quota",
				Usage: "Limit CPU CFS (Completely Fair Scheduler) quota",
			},
			&cli.StringFlag{
				Name:  "cpu-rt-period",
				Usage: "API 1.25+Limit CPU real-time period in microseconds",
			},
			&cli.StringFlag{
				Name:  "cpu-rt-runtime",
				Usage: "API 1.25+Limit CPU real-time runtime in microseconds",
			},
			&cli.StringFlag{
				Name:  "cpu-shares",
				Usage: "CPU shares (relative weight)",
				Aliases: []string{
					"c",
				},
			},
			&cli.StringFlag{
				Name:  "cpus",
				Usage: "API 1.25+Number of CPUs",
			},
			&cli.StringFlag{
				Name:  "cpuset-cpus",
				Usage: "CPUs in which to allow execution (0-3, 0,1)",
			},
			&cli.StringFlag{
				Name:  "cpuset-mems",
				Usage: "MEMs in which to allow execution (0-3, 0,1)",
			},
			&cli.StringFlag{
				Name:  "detach",
				Usage: "Run container in background and print container ID",
				Aliases: []string{
					"d",
				},
			},
			&cli.StringFlag{
				Name:  "detach-keys",
				Usage: "Override the key sequence for detaching a container",
			},
			&cli.StringFlag{
				Name:  "device",
				Usage: "Add a host device to the container",
			},
			&cli.StringFlag{
				Name:  "device-cgroup-rule",
				Usage: "Add a rule to the cgroup allowed devices list",
			},
			&cli.StringFlag{
				Name:  "device-read-bps",
				Usage: "Limit read rate (bytes per second) from a device",
			},
			&cli.StringFlag{
				Name:  "device-read-iops",
				Usage: "Limit read rate (IO per second) from a device",
			},
			&cli.StringFlag{
				Name:  "device-write-bps",
				Usage: "Limit write rate (bytes per second) to a device",
			},
			&cli.StringFlag{
				Name:  "device-write-iops",
				Usage: "Limit write rate (IO per second) to a device",
			},
			&cli.StringFlag{
				Name:  "disable-content-trust",
				Value: "true",
				Usage: "Skip image verification",
			},
			&cli.StringFlag{
				Name:  "dns",
				Usage: "Set custom DNS servers",
			},
			&cli.StringFlag{
				Name:  "dns-opt",
				Usage: "Set DNS options",
			},
			&cli.StringFlag{
				Name:  "dns-option",
				Usage: "Set DNS options",
			},
			&cli.StringFlag{
				Name:  "dns-search",
				Usage: "Set custom DNS search domains",
			},
			&cli.StringFlag{
				Name:  "domainname",
				Usage: "Container NIS domain name",
			},
			&cli.StringFlag{
				Name:  "entrypoint",
				Usage: "Overwrite the default ENTRYPOINT of the image",
			},
			&cli.StringFlag{
				Name:  "env",
				Usage: "Set environment variables",
				Aliases: []string{
					"e",
				},
			},
			&cli.StringFlag{
				Name:  "env-file",
				Usage: "Read in a file of environment variables",
			},
			&cli.StringFlag{
				Name:  "expose",
				Usage: "Expose a port or a range of ports",
			},
			&cli.StringFlag{
				Name:  "gpus",
				Usage: "API 1.40+GPU devices to add to the container ('all' to pass all GPUs)",
			},
			&cli.StringFlag{
				Name:  "group-add",
				Usage: "Add additional groups to join",
			},
			&cli.StringFlag{
				Name:  "health-cmd",
				Usage: "Command to run to check health",
			},
			&cli.StringFlag{
				Name:  "health-interval",
				Usage: "Time between running the check (ms|s|m|h) (default 0s)",
			},
			&cli.StringFlag{
				Name:  "health-retries",
				Usage: "Consecutive failures needed to report unhealthy",
			},
			&cli.StringFlag{
				Name:  "health-start-period",
				Usage: "API 1.29+Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)",
			},
			&cli.StringFlag{
				Name:  "health-timeout",
				Usage: "Maximum time to allow one check to run (ms|s|m|h) (default 0s)",
			},
			/*
				&cli.StringFlag{
					Name: "hostname",
					Usage: "Container host name",
					Aliases: []string{
						"h",
					},
				},
			*/
			&cli.StringFlag{
				Name:  "init",
				Usage: "API 1.25+Run an init inside the container that forwards signals and reaps processes",
			},
			&cli.StringFlag{
				Name:  "interactive",
				Usage: "Keep STDIN open even if not attached",
				Aliases: []string{
					"i",
				},
			},
			&cli.StringFlag{
				Name:  "io-maxbandwidth",
				Usage: "Maximum IO bandwidth limit for the system drive (Windows only)",
			},
			&cli.StringFlag{
				Name:  "io-maxiops",
				Usage: "Maximum IOps limit for the system drive (Windows only)",
			},
			&cli.StringFlag{
				Name:  "ip",
				Usage: "IPv4 address (e.g., 172.30.100.104)",
			},
			&cli.StringFlag{
				Name:  "ip6",
				Usage: "IPv6 address (e.g., 2001:db8::33)",
			},
			&cli.StringFlag{
				Name:  "ipc",
				Usage: "IPC mode to use",
			},
			&cli.StringFlag{
				Name:  "isolation",
				Usage: "Container isolation technology",
			},
			&cli.StringFlag{
				Name:  "kernel-memory",
				Usage: "Kernel memory limit",
			},
			&cli.StringFlag{
				Name:  "label",
				Usage: "Set meta data on a container",
				Aliases: []string{
					"l",
				},
			},
			&cli.StringFlag{
				Name:  "label-file",
				Usage: "Read in a line delimited file of labels",
			},
			&cli.StringFlag{
				Name:  "link",
				Usage: "Add link to another container",
			},
			&cli.StringFlag{
				Name:  "link-local-ip",
				Usage: "Container IPv4/IPv6 link-local addresses",
			},
			&cli.StringFlag{
				Name:  "log-driver",
				Usage: "Logging driver for the container",
			},
			&cli.StringFlag{
				Name:  "log-opt",
				Usage: "Log driver options",
			},
			&cli.StringFlag{
				Name:  "mac-address",
				Usage: "Container MAC address (e.g., 92:d0:c6:0a:29:33)",
			},
			&cli.StringFlag{
				Name:  "memory",
				Usage: "Memory limit",
				Aliases: []string{
					"m",
				},
			},
			&cli.StringFlag{
				Name:  "memory-reservation",
				Usage: "Memory soft limit",
			},
			&cli.StringFlag{
				Name:  "memory-swap",
				Usage: "Swap limit equal to memory plus swap: '-1' to enable unlimited swap",
			},
			&cli.StringFlag{
				Name:  "memory-swappiness",
				Value: "-1",
				Usage: "Tune container memory swappiness (0 to 100)",
			},
			&cli.StringFlag{
				Name:  "mount",
				Usage: "Attach a filesystem mount to the container",
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "Assign a name to the container",
			},
			&cli.StringFlag{
				Name:  "net",
				Usage: "Connect a container to a network",
			},
			&cli.StringFlag{
				Name:  "net-alias",
				Usage: "Add network-scoped alias for the container",
			},
			&cli.StringFlag{
				Name:  "network",
				Usage: "Connect a container to a network",
			},
			&cli.StringFlag{
				Name:  "network-alias",
				Usage: "Add network-scoped alias for the container",
			},
			&cli.StringFlag{
				Name:  "no-healthcheck",
				Usage: "Disable any container-specified HEALTHCHECK",
			},
			&cli.StringFlag{
				Name:  "oom-kill-disable",
				Usage: "Disable OOM Killer",
			},
			&cli.StringFlag{
				Name:  "oom-score-adj",
				Usage: "Tune host's OOM preferences (-1000 to 1000)",
			},
			&cli.StringFlag{
				Name:  "pid",
				Usage: "PID namespace to use",
			},
			&cli.StringFlag{
				Name:  "pids-limit",
				Usage: "Tune container pids limit (set -1 for unlimited)",
			},
			&cli.StringFlag{
				Name:  "platform",
				Usage: "API 1.32+Set platform if server is multi-platform capable",
			},
			&cli.StringFlag{
				Name:  "privileged",
				Usage: "Give extended privileges to this container",
			},
			&cli.StringFlag{
				Name:  "publish",
				Usage: "Publish a container's port(s) to the host",
				Aliases: []string{
					"p",
				},
			},
			&cli.StringFlag{
				Name:  "publish-all",
				Usage: "Publish all exposed ports to random ports",
				Aliases: []string{
					"P",
				},
			},
			&cli.StringFlag{
				Name:  "pull",
				Value: "missing",
				Usage: "Pull image before running (\"always\"|\"missing\"|\"never\")",
			},
			&cli.StringFlag{
				Name:  "read-only",
				Usage: "Mount the container's root filesystem as read only",
			},
			&cli.StringFlag{
				Name:  "restart",
				Value: "no",
				Usage: "Restart policy to apply when a container exits",
			},
			&cli.StringFlag{
				Name:  "rm",
				Usage: "Automatically remove the container when it exits",
			},
			&cli.StringFlag{
				Name:  "runtime",
				Usage: "Runtime to use for this container",
			},
			&cli.StringFlag{
				Name:  "security-opt",
				Usage: "Security Options",
			},
			&cli.StringFlag{
				Name:  "shm-size",
				Usage: "Size of /dev/shm",
			},
			&cli.StringFlag{
				Name:  "sig-proxy",
				Value: "true",
				Usage: "Proxy received signals to the process",
			},
			&cli.StringFlag{
				Name:  "stop-signal",
				Value: "SIGTERM",
				Usage: "Signal to stop a container",
			},
			&cli.StringFlag{
				Name:  "stop-timeout",
				Usage: "API 1.25+Timeout (in seconds) to stop a container",
			},
			&cli.StringFlag{
				Name:  "storage-opt",
				Usage: "Storage driver options for the container",
			},
			&cli.StringFlag{
				Name:  "sysctl",
				Usage: "Sysctl options",
			},
			&cli.StringFlag{
				Name:  "tmpfs",
				Usage: "Mount a tmpfs directory",
			},
			&cli.StringFlag{
				Name:  "tty",
				Usage: "Allocate a pseudo-TTY",
				Aliases: []string{
					"t",
				},
			},
			&cli.StringFlag{
				Name:  "ulimit",
				Usage: "Ulimit options",
			},
			&cli.StringFlag{
				Name:  "user",
				Usage: "Username or UID (format: <name|uid>[:<group|gid>])",
				Aliases: []string{
					"u",
				},
			},
			&cli.StringFlag{
				Name:  "userns",
				Usage: "User namespace to use",
			},
			&cli.StringFlag{
				Name:  "uts",
				Usage: "UTS namespace to use",
			},
			&cli.StringFlag{
				Name:  "volume",
				Usage: "Bind mount a volume",
				Aliases: []string{
					"v",
				},
			},
			&cli.StringFlag{
				Name:  "volume-driver",
				Usage: "Optional volume driver for the container",
			},
			&cli.StringFlag{
				Name:  "volumes-from",
				Usage: "Mount volumes from the specified container(s)",
			},
			&cli.StringFlag{
				Name:  "workdir",
				Usage: "Working directory inside the container",
				Aliases: []string{
					"w",
				},
			},
		},
	}
}
