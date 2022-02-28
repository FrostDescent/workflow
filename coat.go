package workflow

import (
	"fmt"
	"time"
)

type WearCoatNodeAction struct {
}

func (a *WearCoatNodeAction) Run(i interface{}) {
	fmt.Println("我正在外套...")
	time.Sleep(1 * time.Second)
}
