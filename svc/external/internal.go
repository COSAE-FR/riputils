package external

import (
	"github.com/natefinch/pie"
	log "github.com/sirupsen/logrus"
	"gopkg.in/hlandau/svcutils.v1/exepath"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strings"
)

type Client struct {
	CLient *rpc.Client
	Log *log.Entry
	ServerName string
}

func NewClient(logger *log.Entry, name string, component string) (*Client, error) {
	client := &Client{Log: logger.WithField("component", "internal"), ServerName: GetServerName(name, component)}
	args := []string{GetArgumentMark(name, component)}
	// Append daemon invocation arguments
	if len(os.Args) > 1 {
		for _, a := range os.Args {
			if !strings.HasPrefix(a, ArgumentMarkPrefix) {
				args = append(args, a)
			}
		}
	}
	clt, err := pie.StartProviderCodec(jsonrpc.NewClientCodec, logger.Logger.Out, exepath.Abs, args...)
	if err != nil {
		return nil, err
	}
	client.CLient = clt
	return client, nil
}
