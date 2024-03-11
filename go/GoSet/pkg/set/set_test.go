package set

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func defaultTestStringer(val string) string {
	return fmt.Sprint(val)
}

type SuperCustom struct {
	a string
	b int
}

func Test_NewSet(t *testing.T) {
	t.Run("init a new Set with given single string", func(t *testing.T) {
		got := NewSet[string]("hallo")

		want := &Set[string]{
			container: []string{
				"hallo",
			},
			Stringer: defaultTestStringer,
		}

		require.Equal(t, want.container, got.container)
	})

	t.Run("init a new set with a number", func(t *testing.T) {
		set := NewSet[int](1, 2, 3, 4)

		expected := &Set[int]{
			container: []int{1, 2, 3, 4},
		}
		require.Equal(t, expected.container, set.container)
	})

	t.Run("init a new Set with given multiple strings and one duplicate", func(t *testing.T) {
		got := NewSet("hallo", "du", "depp", "depp")

		want := &Set[string]{
			container: []string{
				"hallo",
				"du",
				"depp",
			},
		}

		require.Equal(t, want.container, got.container)
	})
}

func TestSet_Add(t *testing.T) {

	t.Run("Adds a value to an existing set", func(t *testing.T) {
		set := NewSet[string]()
		set.Add("test")

		expected := &Set[string]{
			container: []string{"test"},
		}
		require.Equal(t, expected.container, set.container)
	})

	t.Run("add value twice with custom type", func(t *testing.T) {

		set := NewSet[SuperCustom](SuperCustom{"test", 1})
		set.Add(SuperCustom{"test", 1})

		expected := &Set[SuperCustom]{
			container: []SuperCustom{
				SuperCustom{"test", 1},
			},
		}

		require.Equal(t, expected.container, set.container)
		require.Equal(t, 1, set.ItemsCount())
	})

	t.Run("add value twice", func(t *testing.T) {
		set := NewSet("hallo")
		set.Add("hallo")

		expected := &Set[string]{
			container: []string{
				"hallo",
			},
		}

		require.Equal(t, expected.container, set.container)
	})
}

func TestSet_Remove(t *testing.T) {
	t.Run("remove from set", func(t *testing.T) {
		set := NewSet("test1")
		set.Add("hallo")
		set.Add("hallo1")
		set.Remove("test1")
		set.Remove("hallo1")

		expected := &Set[string]{
			container: []string{
				"hallo",
			},
		}

		require.Equal(t, expected.container, set.container)
	})

	t.Run("remove from set with custom type", func(t *testing.T) {

		sc1 := SuperCustom{a: "one", b: 1}
		sc2 := SuperCustom{a: "two", b: 2}

		set := NewSet(sc1)
		set.Add(sc2)

		set.Remove(sc1)

		expected := &Set[SuperCustom]{
			container: []SuperCustom{
				sc2,
			},
		}

		require.Equal(t, expected.container, set.container)
	})
}

func TestSet_String(t *testing.T) {
	t.Run("stringify a string set", func(t *testing.T) {
		set := NewSet("hello", "beautiful!")
		got := set.String()

		want := "hello\nbeautiful!"

		require.Equal(t, want, got, "strings do not match")
	})

	t.Run("stringify a int set", func(t *testing.T) {
		set := NewSet(1, 2)
		got := set.String()

		want := "1\n2"

		require.Equal(t, want, got, "strings do not match")
	})

	t.Run("stringify a custom type set", func(t *testing.T) {

		set := NewSet[SuperCustom](SuperCustom{a: "one", b: 1}, SuperCustom{a: "two", b: 2})
		set.Stringer = func(val SuperCustom) string {
			return fmt.Sprintf("(%s,%d)", val.a, val.b)
		}
		got := set.String()

		want := "(one,1)\n(two,2)"

		require.Equal(t, want, got, "strings do not match")
	})
}

func TestSet_IsValueInSet(t *testing.T) {
	t.Run("is value in set", func(t *testing.T) {
		set := NewSet("banana")
		expectTrue := set.IsValueInSet("banana")
		expectFalse := set.IsValueInSet("banana1")

		require.True(t, expectTrue)
		require.False(t, expectFalse)
	})

	t.Run("is value in set with custom type ", func(t *testing.T) {

		set := NewSet[SuperCustom](SuperCustom{a: "one", b: 1}, SuperCustom{a: "two", b: 2})
		set.Stringer = func(val SuperCustom) string {
			return fmt.Sprintf("(%s,%d)", val.a, val.b)
		}

		require.True(t, set.IsValueInSet(SuperCustom{a: "one", b: 1}))

	})
}

func TestSet_ItemsCount(t *testing.T) {

	t.Run("counts items", func(t *testing.T) {
		set := NewSet("hallo")
		itemsCount := set.ItemsCount()

		require.Equal(t, 1, itemsCount)
	})

}

func TestSet_Clear(t *testing.T) {

	t.Run("clear set", func(t *testing.T) {
		set := NewSet("one")
		set.Clear()

		emptySet := NewSet[string]()

		require.Equal(t, emptySet.container, set.container)

	})

}

func TestSet_index(t *testing.T) {

}
