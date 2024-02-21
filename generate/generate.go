package generate

import (
	_ "embed"
	"fmt"
	"github.com/Jumpaku/go-tzot"
	"github.com/samber/lo"
	"io"
	"text/template"
)

type data struct {
	Version     string
	TZVersion   string
	PackageName string
	Zones       []zoneData
}
type zoneData struct {
	IDLiteral   string
	Transitions []transitionData
}
type transitionData struct {
	WhenUnix         int64
	OffsetBeforeNano int64
	OffsetAfterNano  int64
}

//go:embed tzot.gen.go.tpl
var tzotGenGoTemplate string
var executor = template.Must(template.New("tzot.gen.go.tpl").Parse(tzotGenGoTemplate))

func Generate(packageName string, zones []tzot.Zone, writer io.Writer) error {
	err := executor.Execute(writer, data{
		Version:     tzot.ModuleVersion(),
		TZVersion:   tzot.GetTZVersion(),
		PackageName: packageName,
		Zones: lo.Map(zones, func(zone tzot.Zone, index int) zoneData {
			return zoneData{
				IDLiteral: fmt.Sprintf(`%q`, zone.ID),
				Transitions: lo.Map(zone.Transitions, func(transition tzot.Transition, index int) transitionData {
					return transitionData{
						WhenUnix:         transition.When.Unix(),
						OffsetBeforeNano: transition.OffsetBefore.Nanoseconds(),
						OffsetAfterNano:  transition.OffsetAfter.Nanoseconds(),
					}
				}),
			}
		}),
	})
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}
	return nil
}
