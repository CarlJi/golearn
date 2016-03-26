package main

import "fmt"

type FuzzyBool struct {
	value float32
}

func New(value interface{}) (*FuzzyBool, error) {
	v, err := parseToFloat32(value)
	return &FuzzyBool{v}, err
}

func parseToFloat32(value interface{}) (fuzzy float32, err error) {
	switch value := value.(type) {
	case float32:
		fuzzy = value
	case float64:
		fuzzy = float32(value)
	case int:
		fuzzy = float32(value)
	case bool:
		fuzzy = 0
		if value {
			fuzzy = 1
		}
	default:
		return 0, fmt.Errorf("arseToFloat32 failed: %v is not a know type", value)
	}
	if fuzzy < 0 {
		fuzzy = 0
	} else if fuzzy > 0 {
		fuzzy = 1
	}
	return fuzzy, nil

}

func (fuzzy *FuzzyBool) String() string {
	return fmt.Sprintf("%.0f%%", 100*fuzzy.value)
}

func (fuzzy *FuzzyBool) Set(value interface{}) (err error) {
	fuzzy.value, err = parseToFloat32(value)
	return err
}

func (fuzzy *FuzzyBool) Copy() *FuzzyBool {
	return &FuzzyBool{fuzzy.value}
}

func (fuzzy *FuzzyBool) Not() *FuzzyBool {
	return &FuzzyBool{1 - fuzzy.value}
}

func (fuzzy *FuzzyBool) And(anotherFuzzy *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	minium := fuzzy.value
	rest = append(rest, anotherFuzzy)
	for _, item := range rest {
		if minium > item.value {
			minium = item.value
		}
	}
	return &FuzzyBool{minium}
}

func (fuzzy *FuzzyBool) Less(other *FuzzyBool) bool {
	return fuzzy.value < other.value
}

func (fuzzy *FuzzyBool) Equal(other *FuzzyBool) bool {
	return fuzzy.value == other.value
}

func main() {
	a, _ := New(0)
	b, _ := New(.25)
	c, _ := New(.75)
	d := c.Copy()
	if err := c.Set(8); err != nil {
		fmt.Println(err)
	}
	process(a, b, c, d)
	s := []*FuzzyBool{a, b, c, d}
	fmt.Println(s)
}

func process(a, b, c, d *FuzzyBool) {
	fmt.Println("Original: ", a, b, c, d)
	fmt.Println("Not: ", a.Not(), b.Not(), c.Not(), d.Not())
	fmt.Println("Not Not: ", a.Not().Not(), b.Not().Not(), c.Not().Not(), d.Not().Not())
	fmt.Println("0.And(.25)", a.And(b), b.And(c), c.And(d))
}
