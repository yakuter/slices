package slices

import (
	"math/rand"
	"testing"
)

func TestDelete(t *testing.T) {
	t.Run("Delete from slice", func(t *testing.T) {
		s := []int{0, 1, 2, 3}
		s = Delete(s, 0)
		if len(s) != 3 {
			t.Fatalf("could not delete data")
		}
		if s[0] != 1 {
			t.Fatalf("could not delete data")
		}
	})

	t.Run("Delete out of index", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		s := []int{0, 1, 2, 3}
		s = Delete(s, 4)
		if len(s) != 4 {
			t.Fatalf("could not delete data")
		}
		if s[0] != 0 {
			t.Fatalf("could not delete data")
		}
	})

	t.Run("On empty slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		var s []struct{}
		_ = Delete(s, 0)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Insert another slice from beginning", func(t *testing.T) {
		s1 := []int{0, 1, 2, 3}
		s2 := []int{4, 5}
		s1 = Insert(s1, 0, s2...)
		if len(s1) != 6 {
			t.Fatalf("should've insert slice")
		}
		if s1[1] != 5 {
			t.Fatalf("should've insert slice correctly")
		}
	})

	t.Run("Insert another slice from out of index", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		s1 := []int{0, 1, 2, 3}
		s2 := []int{4, 5}
		s1 = Insert(s1, 5, s2...)
		if len(s1) != 6 {
			t.Fatalf("should've insert slice")
		}
		if s1[3] != 4 {
			t.Fatalf("should've insert slice correctly")
		}
	})

	t.Run("On empty slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		var s1 []int
		s2 := []int{4, 5}
		s1 = Insert(s1, 5, s2...)
		if len(s1) != 2 {
			t.Fatalf("should've insert slice")
		}
		if s1[1] != 5 {
			t.Fatalf("should've insert slice correctly")
		}
	})
}

func TestDuplicate(t *testing.T) {
	t.Run("Duplicate a slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		s1 := []int{0, 1, 2, 3}
		s2 := Duplicate(s1)
		if len(s1) != len(s2) {
			t.Fatalf("should've duplicate slice")
		}
		for i, e := range s1 {
			if s2[i] != e {
				t.Fatalf("should've duplicated in correct order")
			}
		}
	})

	t.Run("Duplicate an empty slice", func(t *testing.T) {
		var s1 []int
		s2 := Duplicate(s1)
		if len(s1) != len(s2) {
			t.Fatalf("should've duplicate slice")
		}
	})
}

func TestPush(t *testing.T) {
	t.Run("Push to empty slice", func(t *testing.T) {
		var s1 []int
		s1 = Push(s1, 1)
		if len(s1) != 1 {
			t.Fatalf("should've pushed element into the slice")
		}
	})

	t.Run("Push to a slice", func(t *testing.T) {
		s1 := []int{0}
		s1 = Push(s1, 1)
		if len(s1) != 2 {
			t.Fatalf("should've pushed element into the slice")
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Pop from a slice", func(t *testing.T) {
		s1 := []int{0}
		var e int
		e, s1 = Pop(s1)
		if len(s1) != 0 {
			t.Fatalf("should've popped element from the slice")
		}
		if e != 0 {
			t.Fatalf("should've popped correct element from the slice")
		}
	})

	t.Run("Pop from an empty slice", func(t *testing.T) {
		s1 := []int{}
		var e int
		e, s1 = Pop(s1)
		if len(s1) != 0 {
			t.Fatalf("should've popped element from the slice")
		}
		if e != 0 {
			t.Fatalf("should've popped zero value of the element from the slice")
		}
	})
}

func TestShuffle(t *testing.T) {
	t.Run("Shuffle an empty slice", func(t *testing.T) {
		s1 := []int{}
		s1 = Shuffle(s1)
	})

	rand.Seed(42) // get default Source to a deterministic state.

	t.Run("Shuffle a slice", func(t *testing.T) {
		s1 := []int{0, 1}
		s2 := Duplicate(s1)
		s1 = Shuffle(s1)
		if len(s1) != 2 {
			t.Fatalf("should've shuffled slice, not changing its length")
		}

		if s1[0] != s2[1] || s1[1] != s2[0] {
			t.Fatalf("should've shuffled")
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("Reverse an empty slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		s1 := []int{}
		s1 = Reverse(s1)
	})

	t.Run("Reverse a slice", func(t *testing.T) {
		s1 := []int{0, 1}
		s2 := Duplicate(s1)
		s1 = Reverse(s1)
		if len(s1) != 2 {
			t.Fatalf("should've shuffled slice, not changing its length")
		}

		if s1[0] != s2[1] || s1[1] != s2[0] {
			t.Fatalf("should've shuffled")
		}
	})
}
