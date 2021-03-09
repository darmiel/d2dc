package cfg

import "gopkg.in/yaml.v2"

type DockerComposeConfig struct {
	Services map[string]*Service // --name
}

type Build struct {
	Context    string `yaml:"context,omitempty"`
	Dockerfile string `yaml:"dockerfile,omitempty"`
}

type Service struct {
	Build       *Build            `yaml:"build,omitempty"`
	Image       string            `yaml:"image,omitempty"`
	StdinOpen   bool              `yaml:"stdin_open,omitempty"`
	Tty         bool              `yaml:"tty,omitempty"`
	Entrypoint  string            `yaml:"entrypoint,omitempty"`
	Command     string            `yaml:"command,omitempty"`
	EnvFile     []string          `yaml:"env_file,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	Volumes     []string          `yaml:"volumes,omitempty"`
	Labels      map[string]string `yaml:"labels,omitempty"`
	Expose      []string          `yaml:"expose,omitempty"`
	Publish     []string          `yaml:"ports,omitempty"`
	Networks    []string          `yaml:"networks,omitempty"`
	Restart     string            `yaml:"restart,omitempty"`
	Tmpfs       []string          `yaml:"tmpfs,omitempty"`
	PID         string            `yaml:"pid,omitempty"`
}

func (s *Service) Marshal(name string, configFormat ...bool) (out []byte, err error) {
	useConfigFormat := true
	if len(configFormat) > 0 {
		useConfigFormat = configFormat[0]
	}

	var data interface{}
	if useConfigFormat {
		cfg := &DockerComposeConfig{}
		cfg.Services = make(map[string]*Service)
		cfg.Services[name] = s
		data = cfg
	} else {
		data = s
	}

	out, err = yaml.Marshal(data)
	return
}

var SupportedFlags = []string{
	"name",        // ok
	"env",         // ok
	"env-file",    // ok
	"expose",      // ok
	"publish",     // ok
	"interactive", // ok
	"tty",         // ok
	"pid",         // ok
	"label",       // ok
	"mount",       // ok
	"network",     // ok
	"restart",     // ok
	"volume",      // ok
	"entrypoint",  // ok
	//

}
