package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"log"
	"strings"
)

func ActionRun(ctx *cli.Context) (err error) {
	if ctx.Args().Len() < 1 {
		return fmt.Errorf("image name missing")
	}

	cfg := &DockerComposeConfig{Services: make(map[string]*Service)}
	srv := &Service{}

	// load image
	srv.Image = ctx.Args().Get(0)
	// load entrypoint
	if ctx.Args().Len() > 1 {
		srv.Entrypoint = strings.Join(ctx.Args().Slice()[1:], " ")
	}

	// insert values
	// name
	name := ctx.String("name")
	if name == "" {
		name = "unnamed"
	}

	// env
	for _, e := range ctx.StringSlice("env") {
		if !strings.Contains(e, "=") {
			return fmt.Errorf("invalid environment: %s (missing '='-separator)", e)
		}
		if srv.Environment == nil {
			srv.Environment = make(map[string]string)
		}

		index := strings.Index(e, "=")
		key := e[:index]
		value := e[index+1:]

		srv.Environment[key] = value
	}

	// env-file
	for _, e := range ctx.StringSlice("env-file") {
		srv.EnvFile = append(srv.EnvFile, e)
	}

	// expose
	for _, e := range ctx.StringSlice("expose") {
		srv.Expose = append(srv.Expose, e)
	}

	// publish
	for _, e := range ctx.StringSlice("publish") {
		srv.Publish = append(srv.Publish, e)
	}

	// interactive
	if ctx.Bool("interactive") {
		srv.StdinOpen = true
	}

	// tty
	if ctx.Bool("tty") {
		srv.Tty = true
	}

	// pid
	if pid := ctx.String("pid"); pid != "" {
		srv.Pid = pid
	}

	// labels
	for _, l := range ctx.StringSlice("label") {
		if !strings.Contains(l, "=") {
			return fmt.Errorf("invalid label: %s (missing '='-separator)", l)
		}
		if srv.Labels == nil {
			srv.Labels = make(map[string]string)
		}

		index := strings.Index(l, "=")
		key := l[:index]
		value := l[index+1:]

		srv.Labels[key] = value
	}

	// mount
	for _, m := range ctx.StringSlice("mount") {
		srv.Tmpfs = append(srv.Tmpfs, m)
	}

	// network
	for _, n := range ctx.StringSlice("network") {
		srv.Networks = append(srv.Networks, n)
	}

	// restart
	if restart := ctx.String("restart"); restart != "" {
		srv.Restart = restart
	}

	// volumes
	for _, v := range ctx.StringSlice("volume") {
		srv.Volumes = append(srv.Volumes, v)
	}

	// entrypoint
	if entrypoint := ctx.String("entrypoint"); entrypoint != "" {
		srv.Entrypoint = entrypoint
	}

	// serialize
	cfg.Services[name] = srv
	var out []byte
	out, err = yaml.Marshal(cfg)
	if err != nil {
		return
	}

	log.Println("YAML:")
	fmt.Println(string(out))

	return nil
}

func CommandRun() *cli.Command {
	return &cli.Command{
		Name:   "run",
		Action: ActionRun,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "env",
				Usage: "Set environment variables",
				Aliases: []string{
					"e",
				},
			},
			&cli.StringSliceFlag{
				Name:  "env-file",
				Usage: "Read in a file of environment variables",
			},
			&cli.StringSliceFlag{
				Name:  "expose",
				Usage: "Expose a port or a range of ports",
			},
			&cli.StringSliceFlag{
				Name:  "publish",
				Usage: "Expose ports",
				Aliases: []string{
					"p",
				},
			},
			// $ docker run -it
			&cli.BoolFlag{
				Name:  "interactive",
				Usage: "Keep STDIN open even if not attached",
				Aliases: []string{
					"i",
				},
			},
			&cli.StringSliceFlag{
				Name:  "label",
				Usage: "Set meta data on a container",
				Aliases: []string{
					"l",
				},
			},
			&cli.StringSliceFlag{
				Name:  "mount",
				Usage: "Attach a filesystem mount to the container",
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "Assign a name to the container",
			},
			&cli.StringSliceFlag{
				Name:  "network",
				Usage: "Connect a container to a network",
				Aliases: []string{
					"net",
				},
			},

			&cli.StringFlag{
				Name:  "restart",
				Value: "no",
				Usage: "Restart policy to apply when a container exits",
			},
			// $ docker run -it
			&cli.BoolFlag{
				Name:  "tty",
				Usage: "Allocate a pseudo-TTY",
				Aliases: []string{
					"t",
				},
			},
			&cli.StringSliceFlag{
				Name:  "volume",
				Usage: "Bind mount a volume",
				Aliases: []string{
					"v",
				},
			},
			&cli.StringFlag{
				Name:  "entrypoint",
				Usage: "Overwrite the default ENTRYPOINT of the image",
			},

			&cli.StringFlag{
				Name:  "pid",
				Usage: "PID namespace to use",
			},

			//// Not available in docker-compose (on first sight)
			// $ echo "test" | docker run -i -a stdin ubuntu cat -
			&cli.StringFlag{
				Name:  "attach",
				Usage: "Attach to STDIN, STDOUT or STDERR",
				Aliases: []string{
					"a",
				},
			},
			&cli.BoolFlag{
				Name:  "detach",
				Usage: "Run container in background and print container ID",
				Aliases: []string{
					"d",
				},
			},
			&cli.StringFlag{
				Name:  "ip",
				Usage: "IPv4 address (e.g., 172.30.100.104)",
			},
			&cli.StringFlag{
				Name:  "ip6",
				Usage: "IPv6 address (e.g., 2001:db8::33)",
			},
			&cli.PathFlag{
				Name:  "label-file",
				Usage: "Read in a line delimited file of labels",
			},
			&cli.StringFlag{
				Name:  "memory",
				Usage: "Memory limit",
				Aliases: []string{
					"m",
				},
			},
			&cli.BoolFlag{
				Name:  "publish-all",
				Usage: "Publish all exposed ports to random ports",
				Aliases: []string{
					"P",
				},
			},
			&cli.StringFlag{
				Name:  "pull",
				Usage: "Pull image before running (\"always\"|\"missing\"|\"never\")",
			},
			&cli.BoolFlag{
				Name:  "rm",
				Usage: "Automatically remove the container when it exits",
			},
			&cli.StringFlag{
				Name:  "workdir",
				Usage: "Working directory inside the container",
				Aliases: []string{
					"w",
				},
			},

			//// Implement later
			&cli.StringFlag{
				Name:  "health-timeout",
				Usage: "Maximum time to allow one check to run (ms|s|m|h) (default 0s)",
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
				Name:  "hostname",
				Usage: "Container host name",
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
				Name:  "privileged",
				Usage: "Give extended privileges to this container",
			},
			&cli.StringFlag{
				Name:  "link",
				Usage: "Add link to another container",
			},
			&cli.UintFlag{
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
				Usage: "API 1.41+Cgroup namespace to use (host|private)\n'host':    Run the container in the Docker host's cgroup namespace\n'private': Run the container in its own private cgroup namespace\n'':        Use the cgroup namespace as configured by the\n           default-cgroupns-mode option on the daemon (default)",
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
				Name:  "gpus",
				Usage: "API 1.40+GPU devices to add to the container ('all' to pass all GPUs)",
			},
			&cli.StringFlag{
				Name:  "group-add",
				Usage: "Add additional groups to join",
			},
			&cli.StringFlag{
				Name:  "init",
				Usage: "API 1.25+Run an init inside the container that forwards signals and reaps processes",
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
				Name:  "net-alias",
				Usage: "Add network-scoped alias for the container",
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
				Name:  "pids-limit",
				Usage: "Tune container pids limit (set -1 for unlimited)",
			},
			&cli.StringFlag{
				Name:  "platform",
				Usage: "API 1.32+Set platform if server is multi-platform capable",
			},
			&cli.StringFlag{
				Name:  "read-only",
				Usage: "Mount the container's root filesystem as read only",
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
				Name:  "volume-driver",
				Usage: "Optional volume driver for the container",
			},
			&cli.StringFlag{
				Name:  "volumes-from",
				Usage: "Mount volumes from the specified container(s)",
			},
		},
	}
}
