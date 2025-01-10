package stack

import "testing"

func TestIntStack(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)

	expected_results := []int{3, 2, 1}

	for i := 0; i < len(expected_results); i++ {
		val, error := s.Pop()
		if error != nil {
			t.Fatal("Stack was emptied out before all expected results could be verified")
		}
		t.Logf("Pop value: %v", *val)
	}

}

func CheckStackEmptyReportsError(t *testing.T) {

}

func TestStringStack(t *testing.T) {
	var s Stack[string]
	s.Push("world")
	s.Push("hello")

	expectedResults := []string{"hello", "world"}

	for i := 0; i < len(expectedResults); i++ {
		val, error := s.Pop()
		if error != nil {
			t.Fatal("Stack was emptied out before all expected results could be verified")
		}
		t.Logf("Pop value: %v", *val)
	}
}
