package model

import "github.com/otiai10/push-kcwidget/common"
import "github.com/otiai10/rodeo"

var vaquero *rodeo.Vaquero

func init() {
	host, port := common.GetRedisHostAndPort()
	vaq, e := rodeo.NewVaquero(host, port)
	if e != nil {
		panic(e)
	}
	vaquero = vaq
}
