package observer_test

import "github.com/otiai10/push-kcwidget/observer"
import "testing"
import . "github.com/otiai10/mint"

import "time"

func TestObserver(t *testing.T) {
	obs := observer.New()
	Expect(t, obs).TypeOf("*observer.Observer")
}

func TestObserver_Start(t *testing.T) {
	obs := observer.New()
	obs.Start()

	time.Sleep(10 * time.Second)
}
