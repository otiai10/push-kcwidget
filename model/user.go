package model

type User struct {
	TwitterIdStr string    // "140021552"
	Name         string    // "otiai10"
	Services     []Service // AppleとかAndroidとかがTokenとともに入る
	Events       UserEvents
}

type UserEvents struct {
	Missions []Mission
	Nyukyos  []Nyukyo
}

type Mission struct{}
type Nyukyo struct{}
