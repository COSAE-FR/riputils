package external

import (
	"github.com/COSAE-FR/riputils/common/logging"
	"github.com/COSAE-FR/riputils/svc/helpers"
	"github.com/COSAE-FR/riputils/svc/shared"
	"github.com/natefinch/pie"
	log "github.com/sirupsen/logrus"
	"gopkg.in/hlandau/easyconfig.v1"
	"net/rpc/jsonrpc"
	"os"
)

type Process struct {
	Name       string
	Version    string
	Component  string
	Privileges bool
	Provider   pie.Server
	Log        *log.Entry
}

type Option func(process *Process)

func WithVersion(version string) Option {
	return func(configuration *Process) {
		configuration.Version = version
	}
}

func WithComponent(name string) Option {
	return func(configuration *Process) {
		configuration.Component = name
	}
}

func WithPrivileges() Option {
	return func(configuration *Process) {
		configuration.Privileges = true
	}
}

type Configurer func(logger *log.Entry, path string) (interface{}, error)

func NewProcess(name string, newServer Configurer, opts ...Option) (*Process, error) {
	log.SetOutput(os.Stderr)
	config := &Process{Name: name}
	for _, opt := range opts {
		if opt != nil {
			opt(config)
		}
	}
	config.Log = logging.SetupLog(logging.Config{
		Level:     "error",
		App:       config.Name,
		Version:   config.Version,
		Component: "external",
	})

	cfg := shared.Config{}

	configurator := &easyconfig.Configurator{
		ProgramName: config.Name,
	}
	err := easyconfig.Parse(configurator, &cfg)
	if err != nil {
		return config, err
	}
	server, err := newServer(config.Log, cfg.Conf)
	if err != nil {
		return nil, err
	}
	if !config.Privileges {
		err := helpers.DropPrivileges()
		if err != nil {
			return nil, err
		}
	}
	config.Provider = pie.NewProvider()
	if err := config.Provider.RegisterName(GetServerName(config.Name, config.Component), server); err != nil {
		return nil, err
	}
	return config, nil
}

func (p *Process) Start() error {
	p.Provider.ServeCodec(jsonrpc.NewServerCodec)
	return nil
}
