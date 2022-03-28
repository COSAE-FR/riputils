package svc

import (
	"fmt"
	"strings"
)

func OSDefaultConfigurationFile(name string) string {
	return fmt.Sprintf("/usr/local/etc/%s.yml", strings.ToLower(name))
}
