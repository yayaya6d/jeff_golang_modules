package set

type void struct{}
type val interface{}

var member void

type Set interface {
	Values() *map[val]void
	Exist(val) bool
	Insert(val)
	Delete(val)
	Size() int
}

type set struct {
	ms map[val]void
}

func NewSet() Set {
	s := &set{
		ms: make(map[val]void),
	}
	return s
}

func (s *set) Values() *map[val]void {
	return &s.ms
}

func (s *set) Exist(v val) bool {
	_, exist := s.ms[v]
	return exist
}

func (s *set) Insert(v val) {
	s.ms[v] = member
}

func (s *set) Delete(v val) {
	if s.Exist(v) {
		delete(s.ms, v)
	}
}

func (s *set) Size() int {
	return len(s.ms)
}
