package tzot

import (
	_ "embed"
	"strings"
)

//go:embed version.txt
var modVersion string

func ModuleVersion() string {
	return strings.TrimSpace(modVersion)
}
