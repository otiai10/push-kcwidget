package main

import "github.com/otiai10/push-kcwidget/observer"
import "fmt"

func main() {
	// ここでobserverをスタートさせる
	obsrvr := observer.New()
	closer := obsrvr.Start()
	e := <-closer
	fmt.Printf("Error!! %+v\n", e)
}
