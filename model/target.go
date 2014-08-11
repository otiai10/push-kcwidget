package model

const (
	PushTypeApple = iota
	// PushServiceAndroid = iota
)

type Target struct {
	Type  int
	Token string
}
