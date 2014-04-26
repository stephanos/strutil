package strings2_test

import (
	. "github.com/101loops/bdd"
	"github.com/101loops/strings2"
)

var _ = Describe("substring", func() {

	It("returns tail", func() {
		Check(strings2.Tail(""), Equals, "")
		Check(strings2.Tail("a"), Equals, "")
		Check(strings2.Tail("abc"), Equals, "bc")
	})

	It("returns init", func() {
		Check(strings2.Init(""), Equals, "")
		Check(strings2.Init("a"), Equals, "")
		Check(strings2.Init("abc"), Equals, "ab")
	})

	It("returns head", func() {
		Check(strings2.Head(""), Equals, "")
		Check(strings2.Head("!"), Equals, "!")
		Check(strings2.Head("abc"), Equals, "a")
	})

	It("returns last", func() {
		Check(strings2.Last(""), Equals, "")
		Check(strings2.Last("!"), Equals, "!")
		Check(strings2.Last("abc"), Equals, "c")
	})

	It("truncates", func() {
		Check(strings2.Truncate("abcdef", 0), Equals, "")
		Check(strings2.Truncate("abcdef", 1), Equals, "a...")
		Check(strings2.Truncate("abcdef", 2), Equals, "ab...")
		Check(strings2.Truncate("abcdef", 3), Equals, "abcdef")
		Check(strings2.Truncate("abcdef", 4), Equals, "abcdef")
		Check(strings2.Truncate("abcdef", 5), Equals, "abcdef")
		Check(strings2.Truncate("abcdef", 6), Equals, "abcdef")

		Check(strings2.Truncate("a b c d e f", 6), Equals, "a b c...")
		Check(strings2.Truncate("ab\ncdef", 3), Equals, "ab...")

		Check(strings2.Truncate("abcdef", 3, ".."), Equals, "abc..")
	})

	Context("take chars", func() {

		It("from left", func() {
			Check(strings2.Take("", 1), Equals, "")
			Check(strings2.Take("Hello World!", 5), Equals, "Hello")
			Check(strings2.Take("Hello World!", 100), Equals, "Hello World!")
		})

		It("from right", func() {
			Check(strings2.TakeRight("", 1), Equals, "")
			Check(strings2.TakeRight("Hello World!", 6), Equals, "World!")
			Check(strings2.TakeRight("Hello World!", 100), Equals, "Hello World!")
		})

		It("before pattern from left", func() {
			Check(strings2.TakeBefore(":8080", ":"), Equals, "")
			Check(strings2.TakeBefore("localhost", ":"), Equals, "localhost")
			Check(strings2.TakeBefore("http://localhost:8080", ":"), Equals, "http")
		})

		It("after pattern from left", func() {
			Check(strings2.TakeAfter(":8080", ":"), Equals, "8080")
			Check(strings2.TakeAfter("localhost", ":"), Equals, "localhost")
			Check(strings2.TakeAfter("http://localhost:8080", "://"), Equals, "localhost:8080")
		})

		It("before pattern from right", func() {
			Check(strings2.TakeRightBefore(":8080", ":"), Equals, "")
			Check(strings2.TakeRightBefore("localhost", ":"), Equals, "localhost")
			Check(strings2.TakeRightBefore("http://localhost:8080", ":"), Equals, "http://localhost")
		})

		It("after pattern from right", func() {
			Check(strings2.TakeRightAfter("localhost:", ":"), Equals, "")
			Check(strings2.TakeRightAfter("localhost", ":"), Equals, "localhost")
			Check(strings2.TakeRightAfter("http://localhost:8080", ":"), Equals, "8080")
		})
	})
})
