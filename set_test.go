package set

import (
	"testing"
	"reflect"
)

func Test_Union(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "4", "5")
	x := New("5", "6", "7")

	u := Union(s, r, x)

	if u.Size() != 7 {
		t.Error("Union: the merged set doesn't have all items in it.")
	}

	if !u.Has("1", "2", "3", "4", "5", "6", "7") {
		t.Error("Union: merged items are not availabile in the set.")
	}

	y := Union()
	if y.Size() != 0 {
		t.Error("Union: should have zero items because nothing is passed")
	}

	z := Union(x)
	if z.Size() != 3 {
		t.Error("Union: the merged set doesn't have all items in it.")
	}

}

func Test_Difference(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "4", "5")
	x := New("5", "6", "7")
	u := Difference(s, r, x)

	if u.Size() != 2 {
		t.Error("Difference: the set doesn't have all items in it.")
	}

	if !u.Has("1", "2") {
		t.Error("Difference: items are not availabile in the set.")
	}

	y := Difference()
	if y.Size() != 0 {
		t.Error("Difference: size should be zero")
	}

	z := Difference(s)
	if z.Size() != 3 {
		t.Error("Difference: size should be four")
	}
}

func Test_Intersection(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "5")
	u := Intersection(s, r)

	if u.Size() != 1 {
		t.Error("Intersection: the set doesn't have all items in it.")
	}

	if !u.Has("3") {
		t.Error("Intersection: items after intersection are not availabile in the set.")
	}
}

func Test_SymmetricDifference(t *testing.T) {
	s := New("1", "2", "3")
	r := New("3", "4", "5")
	u := SymmetricDifference(s, r)

	if u.Size() != 4 {
		t.Error("SymmetricDifference: the set doesn't have all items in it.")
	}

	if !u.Has("1", "2", "4", "5") {
		t.Error("SymmetricDifference: items are not availabile in the set.")
	}
}

func Test_StringSlice(t *testing.T) {
	s := New("san francisco", "istanbul", 3.14, 1321, "ankara")
	u := StringSlice(s)

	if len(u) != 3 {
		t.Error("StringSlice: slice should only have three items")
	}

	for _, item := range u {
		r := reflect.TypeOf(item)
		if r.Kind().String() != "string" {
			t.Error("StringSlice: slice item should be a string")
		}
	}
}

func Test_IntSlice(t *testing.T) {
	s := New("san francisco", "istanbul", 3.14, 1321, "ankara", 8876)
	u := IntSlice(s)

	if len(u) != 2 {
		t.Error("IntSlice: slice should only have two items")
	}

	for _, item := range u {
		r := reflect.TypeOf(item)
		if r.Kind().String() != "int" {
			t.Error("Intslice: slice item should be a int")
		}
	}
}

func BenchmarkSetEquality(b *testing.B) {
	s := New()
	u := New()

	for i := 0; i < b.N; i++ {
		s.Add(i)
		u.Add(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.IsEqual(u)
	}
}

func BenchmarkSubset(b *testing.B) {
	s := New()
	u := New()

	for i := 0; i < b.N; i++ {
		s.Add(i)
		u.Add(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.IsSubset(u)
	}
}
