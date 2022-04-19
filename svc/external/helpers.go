package external

import (
	"fmt"
	"regexp"
	"strings"
)

const ArgumentMarkPrefix = "__SUBPROCESS__"
const ArgumentMark = "%s%s__EXTERNAL__%s__"
var ArgumentMarkRex *regexp.Regexp

func init()  {
	ArgumentMarkRex = regexp.MustCompile("^__SUBPROCESS__([a-zA-Z\\-_0-9]+)__EXTERNAL__([a-zA-Z\\-_0-9]+)__$")
}

func GetArgumentMark(name string, component string) string {
	return fmt.Sprintf(ArgumentMark, ArgumentMarkPrefix, name, component)
}

func GetServerName(name string, component string) string {
	Name :=  strings.Title(name) + strings.Title(component)
	return Name
}

func ParseArgumentMark(mark string) (string, string, error) {
	results := ArgumentMarkRex.FindStringSubmatch(mark)
	if len(results) == 3 {
		return results[1], results[2], nil
	}
	return "", "", fmt.Errorf("invalid argument mark %s", mark)
}
