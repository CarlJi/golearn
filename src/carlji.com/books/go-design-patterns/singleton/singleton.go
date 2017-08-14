package singleton

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count += 1
	return s.count
}

// Another singleton implementation
var Instance2 *singleton

func init() {
	Instance2 = new(singleton)
}
