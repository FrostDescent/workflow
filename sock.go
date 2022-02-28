package workflow

import (
	"fmt"
	"time"
)

type WearSocksAction struct {
}

func (a *WearSocksAction) Run(i interface{}) {
	fmt.Println("我正在穿袜子...")
	time.Sleep(1 * time.Second)
}
