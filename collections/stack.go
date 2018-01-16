package collections

type Stack struct {
	first *stackNode
	size  int
}

type stackNode struct {
	value interface{}
	next  *stackNode
}

func (ref *Stack) Count() int {
	return ref.size
}

func (ref *Stack) Push(value interface{}) {
	ref.first = &stackNode{value, ref.first}
	ref.size += 1
}

func (ref *Stack) Pop() interface{} {
	if ref.size > 0 {
		result := ref.first.value
		ref.first = ref.first.next
		ref.size--
		return result
	}
	return nil
}

func (ref *Stack) Peek() interface{} {
	if ref.size > 0 {
		return ref.first.value
	}
	return nil
}
