package external

import (
	"fmt"
	"github.com/COSAE-FR/riputils/common/logging"
	"github.com/COSAE-FR/riputils/svc/helpers"
	"github.com/COSAE-FR/riputils/svc/shared"
	"github.com/natefinch/pie"
	log "github.com/sirupsen/logrus"
	"gopkg.in/hlandau/easyconfig.v1"
	"net/rpc/jsonrpc"
	"os"
)

var subProcessRegistry map[string]*SubProcess

var (
	AlreadyRegisteredError = fmt.Errorf("already registered process")
)

func init() {
	subProcessRegistry = make(map[string]*SubProcess)
}

type SubProcess struct {
	Name        string
	Version     string
	Component   string
	Privileges  bool
	Mark        string
	ConfigureFn Configurer
	log         *log.Entry
	provider    *pie.Server
}

func (p *SubProcess) ServerName() string {
	return GetServerName(p.Name, p.Component)
}

func (p *SubProcess) Configure() error {
	log.SetOutput(os.Stderr)
	p.log = logging.SetupLog(logging.Config{
		Level:     "error",
		App:       p.Name,
		Version:   p.Version,
		Component: "external",
	})

	cfg := shared.Config{}

	configurator := &easyconfig.Configurator{
		ProgramName: p.Name,
	}
	err := easyconfig.Parse(configurator, &cfg)
	if err != nil {
		return err
	}
	server, err := p.ConfigureFn(p.log, cfg.Conf)
	if err != nil {
		return err
	}
	if !p.Privileges {
		err := helpers.DropPrivileges()
		if err != nil {
			return err
		}
	}
	provider := pie.NewProvider()
	p.provider = &provider
	if err := p.provider.RegisterName(p.ServerName(), server); err != nil {
		return err
	}
	return nil
}

func (p *SubProcess) Start() error {
	p.provider.ServeCodec(jsonrpc.NewServerCodec)
	return nil
}

func (p *SubProcess) Stop() error {
	return nil
}

func RegisterSubProcess(mark string, process *SubProcess) error {
	_, found := subProcessRegistry[mark]
	if found {
		return AlreadyRegisteredError
	}
	subProcessRegistry[mark] = process
	return nil
}

func GetRegisteredProcess() *SubProcess {
	if len(os.Args) > 1 {
		p, found := subProcessRegistry[os.Args[1]]
		if found {
			return p
		}
	}
	return nil
}
