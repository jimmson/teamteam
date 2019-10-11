package teamteam

type TeamTeamEventType int

const (
	EventTypeUnknown   TeamTeamEventType = 0
	EventTypeCollected TeamTeamEventType = 1
	EventTypeSubmitted TeamTeamEventType = 2
	eventTypeSentinel  TeamTeamEventType = 3
)

func (t TeamTeamEventType) Valid() bool {
	return t > EventTypeUnknown && t < eventTypeSentinel
}

func (t TeamTeamEventType) ReflexType() int {
	return int(t)
}

func String(t TeamTeamEventType) string {
	switch t {
	case EventTypeUnknown:
		return "Unknown"
	case EventTypeCollected:
		return "Collected"
	case EventTypeSubmitted:
		return "Submitted"
	default:
		return "Invalid event type"
	}
}
