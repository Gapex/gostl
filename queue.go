package gostl

type QueueT interface{}

type Queue struct {
	elements []QueueT
}

func (self *Queue) Size() uint32 {
	return uint32(len(self.elements))
}

func (self *Queue) Empty() bool {
	return self.Size() == 0
}

func (self *Queue) Push(v QueueT) {
	self.elements = append(self.elements, v)
}

func (self *Queue) Pop() QueueT {
	if self.Empty() {
		panic("empty queue")
	}
	res := self.elements[0]
	self.elements = self.elements[1:]
	return res
}

func (self *Queue) Front() *QueueT {
	if self.Empty() {
		panic("empty queue")
	}
	res := &self.elements[0]
	return res
}
