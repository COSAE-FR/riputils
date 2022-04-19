package svc

import (
	"github.com/COSAE-FR/riputils/common/logging"
	"github.com/COSAE-FR/riputils/svc/external"
	"github.com/COSAE-FR/riputils/svc/helpers"
	"github.com/COSAE-FR/riputils/svc/shared"
	"github.com/sirupsen/logrus"
	"gopkg.in/hlandau/easyconfig.v1"
	"gopkg.in/hlandau/service.v2"
	"path/filepath"
)

type serviceConfiguration struct {
	Name                     string
	Version                  string
	ForbidRoot               bool
	DefaultConfigurationPath string
}

type Daemonizer interface {
	Start() error
	Stop() error
}

type Configurable interface {
	Daemonizer
	Configure() error
}

type Loggable interface {
	Daemonizer
	SetupLog(name, version string) *logrus.Entry
}

type SubProcessable interface {
	Daemonizer
	GetSubProcess(name, component string) error
}

type Option func(*serviceConfiguration)

type Configurer func(log *logrus.Entry, cfg shared.Config) (Daemonizer, error)

func WithForbidRoot() Option {
	return func(configuration *serviceConfiguration) {
		configuration.ForbidRoot = true
	}
}

func WithVersion(version string) Option {
	return func(configuration *serviceConfiguration) {
		configuration.Version = version
	}
}

func WithDefaultConfigurationPath(path string) Option {
	return func(configuration *serviceConfiguration) {
		configuration.DefaultConfigurationPath = path
	}
}

func WillBePrivileged() bool {
	return helpers.RunWithPrivileges()
}

func StartService(name string, configParser Configurer, opts ...Option) {
	config := &serviceConfiguration{Name: name}
	for _, opt := range opts {
		if opt != nil {
			opt(config)
		}
	}
	if len(config.DefaultConfigurationPath) == 0 {
		config.DefaultConfigurationPath = OSDefaultConfigurationFile(name)
	}
	logger := logging.SetupLog(logging.Config{
		Level:     "error",
		App:       config.Name,
		Version:   config.Version,
		Component: "main",
	})

	cfg := shared.Config{}

	configurator := &easyconfig.Configurator{
		ProgramName: config.Name,
	}

	err := easyconfig.Parse(configurator, &cfg)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	if len(cfg.Conf) == 0 {
		cfg.Conf = config.DefaultConfigurationPath
	}
	confPath, err := filepath.Abs(cfg.Conf)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	cfg.Conf = confPath
	logger.Tracef("Creating daemon %s", config.Name)
	dmn, err := configParser(logger, cfg)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	ext := external.GetRegisteredProcess()
	if ext != nil {
		if err := ext.Configure(); err != nil {
			logger.Fatalf("Cannot configure external process: %s", err)
		}
		if err := ext.Start(); err != nil {
			logger.Fatalf("Cannot start external process: %s", err)
		}
		return
	}
	logger.Debugf("Starting %s daemon", config.Name)
	service.Main(&service.Info{
		Name:      config.Name,
		AllowRoot: !config.ForbidRoot,
		NewFunc: func() (service.Runnable, error) {
			cfg, ok := dmn.(Configurable)
			if ok {
				return dmn, cfg.Configure()
			}
			return dmn, err
		},
	})
}
