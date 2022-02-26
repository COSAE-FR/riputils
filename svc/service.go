package svc

import (
	"github.com/COSAE-FR/riputils/common/logging"
	"github.com/sirupsen/logrus"
	"gopkg.in/hlandau/easyconfig.v1"
	"gopkg.in/hlandau/service.v2"
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

type Config struct {
	Conf string `usage:"configuration file"`
}

type Option func(*serviceConfiguration)

type Configurer func(log *logrus.Entry, cfg Config) (Daemonizer, error)

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

func StartService(name string, configParser Configurer, opts ...Option) {
	config := &serviceConfiguration{Name: name}
	for _, opt := range opts {
		opt(config)
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

	cfg := Config{}

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
	logger.Debugf("Starting %s daemon", config.Name)
	service.Main(&service.Info{
		Name:      config.Name,
		AllowRoot: true,
		NewFunc: func() (service.Runnable, error) {
			dmn, err := configParser(logger, cfg)
			if err != nil {
				return dmn, err
			}
			cfg, ok := dmn.(Configurable)
			if ok {
				return dmn, cfg.Configure()
			}
			return dmn, err
		},
	})
}
