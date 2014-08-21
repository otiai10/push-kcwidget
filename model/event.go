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
