package common_test

import "github.com/otiai10/push-kcwidget/common"
import "testing"
import . "github.com/otiai10/mint"

func TestGetRedisHostAndPort(t *testing.T) {
	tHost, tPort := common.GetRedisHostAndPort(-1)
	Expect(t, tHost).ToBe("localhost")
	Expect(t, tPort).ToBe("6379")
}
