package workflow

import (
	"context"
	"sync"
	"sync/atomic"
)

type WorkFlow struct {
	done        chan struct{} //结束标识,该标识由结束节点写入
	doneOnce    *sync.Once    //保证并发时只写入一次
	alreadyDone bool          //有节点出错时终止流程标记
	root        *Node         //开始节点
	End         *Node         //结束节点
	edges       []*Edge       //所有经过的边，边连接了节点
}

type Edge struct {
	FromNode *Node
	ToNode   *Node
}

type Node struct {
	Dependency   []*Edge  //依赖的边
	DepCompleted int32    //表示依赖的边有多少个已执行完成，用于判断该节点是否可以执行了
	Task         Runnable //任务执行
	Children     []*Edge  //节点的字边
}

type Runnable interface {
	Run(i interface{})
}

func NewNode(Task Runnable) *Node {
	return &Node{
		Task: Task,
	}
}

func AddEdge(from *Node, to *Node) *Edge {
	edge := &Edge{
		FromNode: from,
		ToNode:   to,
	}
	// from节点的出边，to节点的入边
	from.Children = append(from.Children, edge)
	to.Dependency = append(to.Dependency, edge)
	return edge
}

type EndWorkFlowAction struct {
	done chan struct{} //节点执行完成，往该done写入消息，和workflow中的done共用
	s    *sync.Once    //并发控制，确保只往done中写入一次
}

//结束节点的具体执行任务
func (end *EndWorkFlowAction) Run(i interface{}) {
	end.s.Do(func() { end.done <- struct{}{} })
}

func NewWorkFlow() *WorkFlow {
	wf := &WorkFlow{
		root:     &Node{Task: nil},
		done:     make(chan struct{}, 1),
		doneOnce: &sync.Once{},
	}
	EndNode := NewNode(&EndWorkFlowAction{done: wf.done, s: wf.doneOnce})
	wf.End = EndNode
	return wf
}

func (wf *WorkFlow) AddEdge(from *Node, to *Node) {
	wf.edges = append(wf.edges, AddEdge(from, to))
}

func (wf *WorkFlow) AddStartNode(node *Node) {
	wf.edges = append(wf.edges, AddEdge(wf.root, node))
}

func (wf *WorkFlow) ConnectToEnd(node *Node) {
	wf.edges = append(wf.edges, AddEdge(node, wf.End))
}

func (wf *WorkFlow) StartWithContext(ctx context.Context, i interface{}) {
	wf.root.ExecuteWithContext(ctx, wf, i)
}

func (wf *WorkFlow) WaitDone() {
	<-wf.done
	close(wf.done)
}

func (wf *WorkFlow) Interupt() {
	wf.alreadyDone = true
	wf.doneOnce.Do(func() { wf.done <- struct{}{} })
}

func (n *Node) ExecuteWithContext(ctx context.Context, wf *WorkFlow, i interface{}) {
	if !n.DependencyHasDone() {
		return
	}

	if ctx.Err() != nil {
		return
	}

	if n.Task != nil {
		n.Task.Run(i)
	}

	if len(n.Children) > 0 {
		for id := 1; id < len(n.Children); id++ {
			go func(child *Edge) {
				child.ToNode.ExecuteWithContext(ctx, wf, i)
			}(n.Children[id])
		}
		n.Children[0].ToNode.ExecuteWithContext(ctx, wf, i)
	}
}

func (n *Node) DependencyHasDone() bool {
	if n.Dependency == nil {
		return true
	}

	if len(n.Dependency) == 1 {
		return true
	}

	atomic.AddInt32(&n.DepCompleted, 1)

	return n.DepCompleted == int32(len(n.Dependency))
}
