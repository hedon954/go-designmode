package dom

import (
	"bytes"
	"fmt"
)

// Node 一个文档对象节点
type Node interface {
	String() string
	Parent() Node
	SetParent(node Node)
	Children() []Node
	AddChild(child Node)
	Clone() Node
}

// Element 代表文档对象中的一个元素
type Element struct {
	text     string
	parent   Node
	children []Node
}

func NewElement(text string) *Element {
	return &Element{
		text:     text,
		parent:   nil,
		children: make([]Node, 0),
	}
}

func (e Element) Parent() Node {
	return e.parent
}

func (e Element) SetParent(node Node) {
	e.parent = node
}

func (e Element) Children() []Node {
	return e.children
}

// AddChild 添加孩子节点
func (e Element) AddChild(child Node) {

	// 拷贝一个新的
	cp := child.Clone()
	cp.SetParent(e)
	e.children = append(e.children, cp)
}

// Clone 拷贝
func (e Element) Clone() Node {
	cp := Element{
		text:     e.text,
		parent:   nil,
		children: make([]Node, 0),
	}

	// 这里是深拷贝，因为每一个 child 都是 Element
	for _, child := range e.children {
		cp.AddChild(child)
	}
	return cp
}

// String 输出 DOM 树
func (e Element) String() string {
	buffer := bytes.NewBufferString(e.text)

	for _, c := range e.Children() {
		text := c.String()
		fmt.Fprintf(buffer, "\n %s", text)
	}

	return buffer.String()
}
