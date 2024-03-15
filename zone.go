package tzot

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/Jumpaku/go-assert"
	"github.com/samber/lo"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Zone struct {
	ID          string
	Transitions []Transition
	Rules       []Rule
}

type Transition struct {
	When         time.Time
	OffsetBefore time.Duration
	OffsetAfter  time.Duration
}

type Rule struct {
	RuleType          RuleType
	OffsetBefore      time.Duration
	OffsetAfter       time.Duration
	Month             time.Month
	TimeOffset        time.Duration
	TimeHour          int
	TimeMinute        int
	TimeSecond        int
	DayOfWeek         time.Weekday
	MonthDays         int
	MonthDaysFromLast int
}

func (r Rule) Create(year int) Transition {
	var when time.Time
	switch r.RuleType {
	default:
		panic(fmt.Sprintf("invalid RuleType %q", r.RuleType))
	case RuleTypeWeekDayPositive:
		loc := time.FixedZone("", int(r.TimeOffset/time.Second))
		t := time.Date(year, r.Month, r.MonthDays, r.TimeHour, r.TimeMinute, r.TimeSecond, 0, loc)
		if r.TimeHour == 24 {
			t = time.Date(year, r.Month, r.MonthDays, 0, r.TimeMinute, r.TimeSecond, 0, loc)
		}
		_, dow := t.ISOWeek()
		addDays := int(r.DayOfWeek) - dow
		if addDays < 0 {
			addDays += 7
		}
		if r.TimeHour == 24 {
			addDays += 1
		}
		when = t.AddDate(0, 0, addDays)
	case RuleTypeWeekDayNegative:
		loc := time.FixedZone("", int(r.TimeOffset/time.Second))
		t := time.Date(year, r.Month+1, 1-r.MonthDaysFromLast, r.TimeHour, r.TimeMinute, r.TimeSecond, 0, loc)
		if r.TimeHour == 24 {
			t = time.Date(year, r.Month+1, 1-r.MonthDaysFromLast, 0, r.TimeMinute, r.TimeSecond, 0, loc)
		}
		_, dow := t.ISOWeek()
		addDays := int(r.DayOfWeek) - dow
		if addDays > 0 {
			addDays -= 7
		}
		if r.TimeHour == 24 {
			addDays += 1
		}
		when = t.AddDate(0, 0, addDays)
	case RuleTypeMonthDayPositive:
		loc := time.FixedZone("", int(r.TimeOffset/time.Second))
		when = time.Date(year, r.Month, r.MonthDays, r.TimeHour, r.TimeMinute, r.TimeSecond, 0, loc)
	case RuleTypeMonthDayNegative:
		loc := time.FixedZone("", int(r.TimeOffset/time.Second))
		t := time.Date(year, r.Month+1, 1-r.MonthDaysFromLast, r.TimeHour, r.TimeMinute, r.TimeSecond, 0, loc)
		if r.TimeHour == 24 {
			t = time.Date(year, r.Month+1, 1-r.MonthDaysFromLast, 0, r.TimeMinute, r.TimeSecond, 0, loc)
		}
		addDays := 0
		if r.TimeHour == 24 {
			addDays += 1
		}
		when = t.AddDate(0, 0, addDays)
	}
	return Transition{
		When:         when,
		OffsetBefore: r.OffsetBefore,
		OffsetAfter:  r.OffsetAfter,
	}
}

func AvailableZoneIDs() []string {
	return append([]string{}, zoneIDs...)
}

func GetZone(zoneID string) (zone Zone, found bool) {
	zone, found = zones[zoneID]
	return zone, found
}

func GetTZVersion() string {
	return strings.TrimSpace(tzVersion)
}

//go:embed tz-offset-transitions/version
var tzVersion string

//go:embed tz-offset-transitions/tzot.json
var tzotJson []byte

var zones = (func() map[string]Zone {
	var zones []ZoneJSON
	err := json.Unmarshal(tzotJson, &zones)
	assert.State(err == nil, "expect that tzotJson can be unmarshalled: %w", err)

	return lo.SliceToMap(zones, func(zone ZoneJSON) (string, Zone) {
		transitions := lo.Map(zone.Transitions, func(transition TransitionJSON, index int) Transition {
			when, err := time.Parse(time.RFC3339, transition.TransitionTimestamp)
			assert.State(err == nil, "transition.TransitionTimestamp can be parsed as a datetime: %w", err)

			return Transition{
				When:         when,
				OffsetBefore: time.Duration(transition.OffsetSecondsBefore) * time.Second,
				OffsetAfter:  time.Duration(transition.OffsetSecondsAfter) * time.Second,
			}
		})
		slices.SortFunc(transitions, func(a, b Transition) int { return a.When.Compare(b.When) })

		rules := lo.Map(zone.Rules, func(rule RuleJSON, index int) Rule {
			assert.State(len(rule.OffsetTime) >= 9, fmt.Sprintf("rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ': %q", rule.OffsetTime))

			timeStr := rule.OffsetTime[:8]

			hour, err := strconv.Atoi(timeStr[:2])
			assert.State(err == nil, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)
			assert.State(0 <= hour && hour <= 24, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)

			minute, err := strconv.Atoi(timeStr[3:5])
			assert.State(err == nil, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)
			assert.State(0 <= minute && minute < 60, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)

			second, err := strconv.Atoi(timeStr[6:8])
			assert.State(err == nil, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)
			assert.State(0 <= second && second < 60, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)

			if hour == 24 {
				assert.State(minute == 0 && second == 0, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)
			}

			offset := time.Duration(0)
			if offsetStr := rule.OffsetTime[8:]; offsetStr != "Z" {
				assert.State(len(offsetStr) == 6, fmt.Sprintf("rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ': %q", rule.OffsetTime))

				sign := time.Duration(1)
				if offsetStr[0] == '-' {
					sign = -1
				}

				hour, err := strconv.Atoi(offsetStr[1:3])
				assert.State(err == nil, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)
				assert.State(0 <= hour && hour <= 14, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)

				minute, err := strconv.Atoi(offsetStr[4:6])
				assert.State(err == nil, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)
				assert.State(0 <= minute && minute < 60, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)

				if hour == 14 {
					assert.State(minute == 0, "rule.OffsetTime must match 'hh:mm:ss±hh:mm' or 'hh:mm:ssZ'", err)
				}

				offset = sign * (time.Duration(hour)*time.Hour + time.Duration(minute)*time.Minute)
			}

			return Rule{
				RuleType:          rule.Type,
				OffsetBefore:      time.Duration(rule.OffsetSecondsBefore) * time.Second,
				OffsetAfter:       time.Duration(rule.OffsetSecondsAfter) * time.Second,
				Month:             time.Month(rule.Month),
				TimeHour:          hour,
				TimeMinute:        minute,
				TimeSecond:        second,
				TimeOffset:        offset,
				DayOfWeek:         time.Weekday(rule.DayOfWeek),
				MonthDays:         rule.Days,
				MonthDaysFromLast: rule.DaysFromLast,
			}
		})

		return zone.ID, Zone{
			ID:          zone.ID,
			Transitions: transitions,
			Rules:       rules,
		}
	})
})()

var zoneIDs = func() []string {
	zoneIDs := lo.MapToSlice(zones, func(zoneID string, zone Zone) string { return zoneID })
	slices.Sort(zoneIDs)
	return zoneIDs
}()
