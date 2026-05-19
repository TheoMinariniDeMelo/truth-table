package queue

type Queue[T any] []T


func (s *Queue[T]) Add(value T) {
	*s = append(*s, value)
}

func (s *Queue[T]) Poll() T{
	v := (*s)[0]
	*s = (*s)[1:]
	return v
}
