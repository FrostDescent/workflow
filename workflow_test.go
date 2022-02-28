package workflow

import (
	"context"
	"fmt"
	"testing"
)

func TestNewWorkFlow(t *testing.T) {
	wf := NewWorkFlow()

	//构建节点
	UnderpantsNode := NewNode(&WearUnderpantsAction{})
	SocksNode := NewNode(&WearSocksAction{})
	ShirtNode := NewNode(&WearShirtNodeAction{})
	WatchNode := NewNode(&WearWatchNodeAction{})
	TrousersNode := NewNode(&WearTrouserNodeAction{})
	ShoesNode := NewNode(&WearShoesNodeAction{})
	CoatNode := NewNode(&WearCoatNodeAction{})

	//构建节点之间的关系
	wf.AddStartNode(UnderpantsNode)
	wf.AddStartNode(SocksNode)
	wf.AddStartNode(ShirtNode)
	wf.AddStartNode(WatchNode)

	wf.AddEdge(UnderpantsNode, TrousersNode)
	wf.AddEdge(TrousersNode, ShoesNode)
	wf.AddEdge(SocksNode, ShoesNode)
	wf.AddEdge(ShirtNode, CoatNode)
	wf.AddEdge(WatchNode, CoatNode)

	wf.ConnectToEnd(ShoesNode)
	wf.ConnectToEnd(CoatNode)

	var completedAction []string

	wf.StartWithContext(context.Background(), completedAction)
	wf.WaitDone()

	fmt.Println("执行其他逻辑")
}
