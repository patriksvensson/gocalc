package collections

type Queue struct {
	in   Stack
	out  Stack
	size int
}

func (ref *Queue) Count() int {
	return ref.in.Count() + ref.out.Count()
}

func (ref *Queue) Enqueue(value interface{}) {
	ref.in.Push(value)
}

func (ref *Queue) Dequeue() interface{} {
	if ref.out.Count() == 0 {
		for ref.in.Count() > 0 {
			ref.out.Push(ref.in.Pop())
		}
	}
	return ref.out.Pop()
}

func (ref *Queue) DequeueAll() []interface{} {
	var result []interface{}
	for ref.Count() > 0 {
		result = append(result, ref.Dequeue())
	}
	return result
}
