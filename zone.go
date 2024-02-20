package tzot

import (
	_ "embed"
	"encoding/json"
	"github.com/Jumpaku/go-assert"
	"github.com/samber/lo"
	"slices"
	"time"
)

type Zone struct {
	ID          string
	Transitions []Transition
}
type Transition struct {
	When         time.Time
	OffsetBefore time.Duration
	OffsetAfter  time.Duration
}

//go:embed tz-offset-transitions/tzot.json
var tzotJson []byte

var zones = (func() map[string]Zone {
	var zones []ZoneJSON
	err := json.Unmarshal(tzotJson, &zones)
	assert.State(err == nil, "expect that tzotJson can be unmarshalled: %w", err)

	return lo.SliceToMap(zones, func(zone ZoneJSON) (string, Zone) {
		transitions := lo.Map(zone.Transitions, func(transition TransitionJSON, index int) Transition {
			when, err := time.Parse(time.RFC3339, transition.TransitionTimestamp)
			assert.State(err == nil, "expect that transition.TransitionTimestamp can be parsed: %w", err)

			return Transition{
				When:         when,
				OffsetBefore: time.Duration(transition.OffsetSecondsBefore) * time.Second,
				OffsetAfter:  time.Duration(transition.OffsetSecondsAfter) * time.Second,
			}
		})

		slices.SortFunc(transitions, func(a, b Transition) int { return a.When.Compare(b.When) })

		return zone.Zone, Zone{
			ID:          zone.Zone,
			Transitions: transitions,
		}
	})
})()

var zoneIDs = func() []string {
	zoneIDs := lo.MapToSlice(zones, func(zoneID string, zone Zone) string { return zoneID })
	slices.Sort(zoneIDs)
	return zoneIDs
}()

func AvailableZoneIDs() []string {
	return append([]string{}, zoneIDs...)
}

func GetZone(zoneID string) (zone Zone, found bool) {
	zone, found = zones[zoneID]
	return zone, found
}
