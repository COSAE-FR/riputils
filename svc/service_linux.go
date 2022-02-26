package svc

import (
	"fmt"
	"strings"
)

func OSDefaultConfigurationFile(name string) string {
	return fmt.Sprintf("/etc/%s.yml", strings.ToLower(name))
}
