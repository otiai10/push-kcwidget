package model

import "time"
import "sort"
import "github.com/otiai10/rodeo"
import "github.com/otiai10/push-kcwidget/common"

type User struct {
	TwitterIdStr string    // "140021552"
	Name         string    // "otiai10"
	Services     []Service // AppleとかAndroidとかがTokenとともに入る
	Events       Events
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
	if username != "" {
		user.Name = username
	}
	return user.UpdateWithPushService(deviceToken, service)
}

func (user User) UpdateWithPushService(deviceToken, serviceName string) User {
	if deviceToken == "" {
		return user
	}
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

func FindUserByTwitterIdStr(twitterIdStr string) (user User, ok bool) {
	vaquero.Cast(common.Prefix()+"user."+twitterIdStr, &user)
	if user.TwitterIdStr == "" {
		return user, false
	}
	return user, true
}

func (user User) SetEvent(newEvent Event) User {
	var updated bool
	for i, ev := range user.Events {
		if ev.Kind != newEvent.Kind {
			continue
		}
		if ev.Identifier != newEvent.Identifier {
			continue
		}
		user.Events[i] = newEvent
		updated = true
		break
	}
	if !updated {
		user.Events = append(user.Events, newEvent)
	}
	return user
}

func (user User) FindReadyEvents() (events []Event) {
	for _, ev := range user.Events {
		if ev.Finish <= time.Now().Unix() {
			events = append(events, ev)
		}
	}
	return
}

func (user User) CleanUpEvents(checked time.Time) (e error) {
	var events []Event
	for _, ev := range user.Events {
		if ev.Finish < checked.Unix() {
			continue
		}
		events = append(events, ev)
	}
	user.Events = events
	return user.Save()
}

func (user User) FilterPrivateInfo() User {
	for i, _ := range user.Services {
		user.Services[i].Token = ""
	}
	return user
}

func (user User) SortEvents() User {
	sort.Sort(user.Events)
	return user
}
