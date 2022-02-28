package workflow

import (
	"fmt"
	"time"
)

type WearShoesNodeAction struct{}

func (a *WearShoesNodeAction) Run(i interface{}) {
	fmt.Println("我正在穿鞋子...")
	time.Sleep(1 * time.Second)
}
