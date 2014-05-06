package strutil_test

import (
	. "github.com/101loops/bdd"
	"github.com/101loops/strutil"
)

var _ = Describe("substring", func() {

	It("returns tail", func() {
		Check(strutil.Tail(""), Equals, "")
		Check(strutil.Tail("a"), Equals, "")
		Check(strutil.Tail("abc"), Equals, "bc")
	})

	It("returns init", func() {
		Check(strutil.Init(""), Equals, "")
		Check(strutil.Init("a"), Equals, "")
		Check(strutil.Init("abc"), Equals, "ab")
	})

	It("returns head", func() {
		Check(strutil.Head(""), Equals, "")
		Check(strutil.Head("!"), Equals, "!")
		Check(strutil.Head("abc"), Equals, "a")
	})

	It("returns last", func() {
		Check(strutil.Last(""), Equals, "")
		Check(strutil.Last("!"), Equals, "!")
		Check(strutil.Last("abc"), Equals, "c")
	})

	It("truncates", func() {
		Check(strutil.Truncate("abcdef", 0), Equals, "")
		Check(strutil.Truncate("abcdef", 1), Equals, "a...")
		Check(strutil.Truncate("abcdef", 2), Equals, "ab...")
		Check(strutil.Truncate("abcdef", 3), Equals, "abcdef")
		Check(strutil.Truncate("abcdef", 4), Equals, "abcdef")
		Check(strutil.Truncate("abcdef", 5), Equals, "abcdef")
		Check(strutil.Truncate("abcdef", 6), Equals, "abcdef")

		Check(strutil.Truncate("a b c d e f", 6), Equals, "a b c...")
		Check(strutil.Truncate("ab\ncdef", 3), Equals, "ab...")

		Check(strutil.Truncate("abcdef", 3, ".."), Equals, "abc..")
	})

	Context("take chars", func() {

		It("from left", func() {
			Check(strutil.Take("", 1), Equals, "")
			Check(strutil.Take("Hello World!", 5), Equals, "Hello")
			Check(strutil.Take("Hello World!", 100), Equals, "Hello World!")
		})

		It("from right", func() {
			Check(strutil.TakeRight("", 1), Equals, "")
			Check(strutil.TakeRight("Hello World!", 6), Equals, "World!")
			Check(strutil.TakeRight("Hello World!", 100), Equals, "Hello World!")
		})

		It("before pattern from left", func() {
			Check(strutil.TakeBefore(":8080", ":"), Equals, "")
			Check(strutil.TakeBefore("localhost", ":"), Equals, "localhost")
			Check(strutil.TakeBefore("http://localhost:8080", ":"), Equals, "http")
		})

		It("after pattern from left", func() {
			Check(strutil.TakeAfter(":8080", ":"), Equals, "8080")
			Check(strutil.TakeAfter("localhost", ":"), Equals, "localhost")
			Check(strutil.TakeAfter("http://localhost:8080", "://"), Equals, "localhost:8080")
		})

		It("before pattern from right", func() {
			Check(strutil.TakeRightBefore(":8080", ":"), Equals, "")
			Check(strutil.TakeRightBefore("localhost", ":"), Equals, "localhost")
			Check(strutil.TakeRightBefore("http://localhost:8080", ":"), Equals, "http://localhost")
		})

		It("after pattern from right", func() {
			Check(strutil.TakeRightAfter("localhost:", ":"), Equals, "")
			Check(strutil.TakeRightAfter("localhost", ":"), Equals, "localhost")
			Check(strutil.TakeRightAfter("http://localhost:8080", ":"), Equals, "8080")
		})
	})
})
