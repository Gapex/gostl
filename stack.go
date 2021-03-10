package gostl

type StackT interface {
}

type Stack struct {
	elements []StackT
}

func (self *Stack) Push(val StackT) {
	self.elements = append(self.elements, val)
}

func (self *Stack) Pop() StackT {
	if self.Empty() {
		panic("empty stack")
	}
	n := self.Size()
	res := self.elements[n-1]
	self.elements = self.elements[:n-1]
	return res
}

func (self *Stack) Top() *StackT {
	if self.Empty() {
		panic("empty stack")
	}
	n := self.Size()
	res := &self.elements[n-1]
	return res
}

func (self *Stack) Empty() bool {
	return len(self.elements) == 0
}

func (self *Stack) Size() int {
	n := len(self.elements)
	return n
}
