package workflow

import (
	"fmt"
	"time"
)

type WearTrouserNodeAction struct {
}

func (a *WearTrouserNodeAction) Run(i interface{}) {
	fmt.Println("我正在穿长裤...")
	time.Sleep(1 * time.Second)
}
