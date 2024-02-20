package tzot

type ZoneJSON struct {
	Zone        string           `json:"zone"`
	Transitions []TransitionJSON `json:"transitions"`
}

type TransitionJSON struct {
	TransitionTimestamp string `json:"transition_timestamp"`
	OffsetSecondsBefore int    `json:"offset_seconds_before"`
	OffsetSecondsAfter  int    `json:"offset_seconds_after"`
}
