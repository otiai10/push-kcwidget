package model

type User struct {
	TwitterIdStr string    // "140021552"
	Name         string    // "otiai10"
	Services     []Service // AppleとかAndroidとかがTokenとともに入る
	Events       UserEvents
}

type UserEvents struct {
	Missions []Event
	Nyukyos  []Event
}

type Mission struct{} // implements Event
type Nyukyo struct{}  // implements Event
