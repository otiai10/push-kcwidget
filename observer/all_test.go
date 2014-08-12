package observer_test

import "github.com/otiai10/push-kcwidget/observer"
import "github.com/otiai10/push-kcwidget/common"
import "github.com/otiai10/push-kcwidget/model"
import "testing"
import . "github.com/otiai10/mint"

import "time"
import "strconv"

func init() {
	common.SetPrefix("test.")
	var u model.User
	for i := 0; i < 5; i++ {
		dur, _ := time.ParseDuration("-" + strconv.Itoa(i*3) + "s")
		timestamp := time.Now().Add(dur).Unix()
		model.Enqueue(timestamp, u)
	}
}

func TestObserver(t *testing.T) {
	obs := observer.New()
	Expect(t, obs).TypeOf("*observer.Observer")
}

func TestObserver_Start(t *testing.T) {
	obs := observer.New()
	obs.Start()

	time.Sleep(10 * time.Second)

	obs.Kill("Test ended")
}
