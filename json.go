package tzot

type ZoneJSON struct {
	ID          string           `json:"id"`
	Transitions []TransitionJSON `json:"transitions"`
	Rules       []RuleJSON       `json:"rules"`
}

type TransitionJSON struct {
	TransitionTimestamp string `json:"transition_timestamp"`
	OffsetSecondsBefore int    `json:"offset_seconds_before"`
	OffsetSecondsAfter  int    `json:"offset_seconds_after"`
}

type RuleType string

const (
	RuleTypeWeekDayPositive  RuleType = "net.jumpaku.tzot.Rule.WeekDayPositive"
	RuleTypeWeekDayNegative  RuleType = "net.jumpaku.tzot.Rule.WeekDayNegative"
	RuleTypeMonthDayPositive RuleType = "net.jumpaku.tzot.Rule.MonthDayPositive"
	RuleTypeMonthDayNegative RuleType = "net.jumpaku.tzot.Rule.MonthDayNegative"
)

type RuleJSON struct {
	Type                RuleType `json:"type"`
	OffsetSecondsBefore int      `json:"offset_seconds_before"`
	OffsetSecondsAfter  int      `json:"offset_seconds_after"`
	Month               int      `json:"month"`
	OffsetTime          string   `json:"offset_time"`
	DayOfWeek           int      `json:"day_of_week"`    // DayOfWeek has a valid value if and only if Type match RuleTypeWeekDayPositive or RuleTypeWeekDayNegative.
	Days                int      `json:"days"`           // Days has a valid value if and only if Type match RuleTypeWeekDayPositive or RuleTypeMonthDayPositive.
	DaysFromLast        int      `json:"days_from_last"` // DaysFromLast has a valid value if and only if Type match RuleTypeWeekDayPositive or RuleTypeMonthDayPositive.
}
