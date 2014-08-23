package model

type Event struct {
	Finish          int64
	Kind            string
	Label           string
	Message         string
	Prefix          string
	Identifier      string
	Unit            string
	OptionalStrings []string
}

func CreateEventFromRequestParams(
	finish int64,
	message,
	label,
	prefix,
	identifier,
	unit,
	kind string,
	optionalStrings ...string,
) (event Event) {
	event.Finish = finish
	event.Message = message
	event.Kind = kind
	event.Label = label
	event.Prefix = prefix
	event.Identifier = identifier
	event.Unit = unit
	for _, s := range optionalStrings {
		event.OptionalStrings = append(event.OptionalStrings, s)
	}
	return event
}

type Events []Event

func (events Events) Len() int {
	return len(events)
}
func (events Events) Less(i, j int) bool {
	return events[i].Finish < events[j].Finish
}
func (events Events) Swap(i, j int) {
	events[i], events[j] = events[j], events[i]
}
