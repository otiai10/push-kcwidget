package observer_test

import "push-kcwidget/observer"
import "testing"
import . "github.com/otiai10/mint"

func TestObserver(t *testing.T) {
	Expect(t, observer.New()).TypeOf("*observer.Observer")
}
