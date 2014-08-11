package model

type User struct {
	TwitterIdStr string   // "140021552"
	Name         string   // "otiai10"
	Targets      []Target // AppleとかAndroidとかがTokenとともに入る
}
