package relax

type lifo []Vertex

func (l *lifo) Empty() bool {
	return len(*l) == 0
}

func (l *lifo) Top() Vertex {
	if l.Empty() {
		return nil
	}
	return (*l)[len(*l)-1]
}

func (l *lifo) Pop() Vertex {
	top := l.Top
	if top != nil {
		*l = (*l)[:len(*l)-1]
	}
	return top
}

func (l *lifo) Push(v Vertex) {
	*l = append(*l, v)
}
