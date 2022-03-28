package tls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/COSAE-FR/riputils/common"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/acme/autocert"
	"os"
	"strings"
)

type Configuration struct {
	Hostname          string `yaml:"hostname" validate:"fqdn,required"`
	ServerCertificate string `yaml:"cert" validate:"required"`                                       // PEM, path to PEM or acme
	ServerKey         string `yaml:"key" validate:"required_unless=ServerCertificate acme"`          // PEM, path to PEM or empty if ACME
	ServerCa          string `yaml:"ca"`                                                             // PEM, path to PEM or empty if ACME
	Cache             string `yaml:"cache"`                                                          // Recommended if using ACME
	ClientCA          string `yaml:"client_ca"`                                                      // PEM, path to PEM or empty if no client certificate required
	ACMEEmail         string `yaml:"acme_email" validate:"email,required_if=ServerCertificate acme"` // ACME account Email
}

func (c *Configuration) Check() error {
	if err := defaults.Set(c); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil
}

func (c *Configuration) Generate() (*tls.Config, error) {
	return generateTLSConfig(c)
}

func generateTLSConfig(config *Configuration) (*tls.Config, error) {
	if config == nil {
		return nil, nil
	}
	var tlsConfig *tls.Config
	if strings.ToLower(config.ServerCertificate) == "acme" {
		mgr := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(config.Hostname),
			Email:      config.ACMEEmail,
		}
		if config.Cache != "" {
			if !common.IsDirectory(config.Cache) {
				if err := os.MkdirAll(config.Cache, 0750); err != nil {
					return nil, err
				}
			}
			mgr.Cache = autocert.DirCache(config.Cache)
		}
		tlsConfig = mgr.TLSConfig()
	} else {
		config.ServerCertificate = common.FilePathToContent(config.ServerCertificate)
		config.ServerKey = common.FilePathToContent(config.ServerKey)
		config.ServerCa = common.FilePathToContent(config.ServerCa)
		var certPool *x509.CertPool
		if config.ServerCa != "" {
			certPool = x509.NewCertPool()
			if !certPool.AppendCertsFromPEM([]byte(config.ServerCa)) {
				return nil, fmt.Errorf("cannot read server Certificate Authorities")
			}
		}
		cert, err := tls.X509KeyPair([]byte(config.ServerCertificate), []byte(config.ServerKey))
		if err != nil {
			return nil, err
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      certPool,
			ServerName:   config.Hostname,
		}
	}
	if config.ClientCA != "" {
		config.ClientCA = common.FilePathToContent(config.ClientCA)
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM([]byte(config.ClientCA)) {
			return nil, fmt.Errorf("cannot read client Certificate Authorities")
		}
		tlsConfig.ClientCAs = certPool
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}
	return tlsConfig, nil
}
