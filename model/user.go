package model

type User struct {
	TwitterIdStr string    // "140021552"
	Name         string    // "otiai10"
	Services     []Service // AppleとかAndroidとかがTokenとともに入る
}
