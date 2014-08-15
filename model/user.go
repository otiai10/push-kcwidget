package model

import "github.com/otiai10/rodeo"
import "github.com/otiai10/push-kcwidget/common"

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

var vaquero *rodeo.Vaquero

func init() {
	host, port := common.GetRedisHostAndPort()
	vaquero, _ = rodeo.NewVaquero(host, port)
}

func CreaetOrMergeUserWithRegisterParams(username, twitterIdStr, deviceToken, service string) (user User) {
	vaquero.Cast(common.Prefix()+"user."+twitterIdStr, &user)
	user.TwitterIdStr = twitterIdStr
	user.Name = username
	return user.UpdateWithPushService(deviceToken, service)
}

func (user User) UpdateWithPushService(deviceToken, serviceName string) User {
	s := Service{
		GetPushTypeByName(serviceName),
		deviceToken,
	}
	updated := false
	for i, srvc := range user.Services {
		if s.Type != srvc.Type {
			continue
		}
		user.Services[i] = s
		updated = true
	}
	if !updated {
		user.Services = append(user.Services, s)
	}
	return user
}

func (user User) Save() (e error) {
	return vaquero.Store(common.Prefix()+"user."+user.TwitterIdStr, user)
}
