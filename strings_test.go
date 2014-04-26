package strings2_test

import (
	. "github.com/101loops/bdd"
	"github.com/101loops/strings2"
)

var _ = Describe("String Utility", func() {

	It("deletes all spaces", func() {
		Check(strings2.DeleteWhitespace(""), Equals, "")
		Check(strings2.DeleteWhitespace(" a  b "), Equals, "ab")
		Check(strings2.DeleteWhitespace("a \n \t b"), Equals, "ab")
	})

	It("trims all spaces", func() {
		Check(strings2.ShrinkWhitespace(""), Equals, "")
		Check(strings2.ShrinkWhitespace(" "), Equals, "")
		Check(strings2.ShrinkWhitespace(" test "), Equals, "test")
		Check(strings2.ShrinkWhitespace(" te  st "), Equals, "te st")
		Check(strings2.ShrinkWhitespace(" te\tst "), Equals, "te st")
		Check(strings2.ShrinkWhitespace(" te\nst "), Equals, "te st")
	})

	It("reverses string", func() {
		Check(strings2.Reverse(""), Equals, "")
		Check(strings2.Reverse("abc"), Equals, "cba")
	})

	It("checks if empty", func() {
		Check(strings2.IsEmpty(""), IsTrue)

		Check(strings2.IsEmpty(" "), IsFalse)
		Check(strings2.IsEmpty("abc"), IsFalse)
	})

	It("checks if blank", func() {
		Check(strings2.IsBlank(""), IsTrue)
		Check(strings2.IsBlank("\n"), IsTrue)
		Check(strings2.IsBlank("\t"), IsTrue)
		Check(strings2.IsBlank(" "), IsTrue)

		Check(strings2.IsBlank("abc"), IsFalse)
		Check(strings2.IsBlank(" abc "), IsFalse)
	})

	It("returns index after substring", func() {
		Check(strings2.IndexAfter("", ""), Equals, 0)
		Check(strings2.IndexAfter("", "a"), Equals, -1)
		Check(strings2.IndexAfter("abc", "ab"), Equals, 2)
		Check(strings2.IndexAfter("abc", "bc"), Equals, 3)
		Check(strings2.IndexAfter("abc", "xyz"), Equals, -1)
	})

	It("returns lines", func() {
		Check(strings2.Lines(""), Equals, []string{""})
		Check(strings2.Lines("a"), Equals, []string{"a"})
		Check(strings2.Lines("a\nb\nc"), Equals, []string{"a", "b", "c"})
	})

	It("wraps text to fixed width", func() {
		Check(strings2.WordWrap(``, 5), Equals, ``)
		//Check(WordWrap(`abcdef`, 5), Equals, `abcdef`)
		//Check(WordWrap(`abcdefghijklmno`, 5), Equals, `abcdef`)
	})

	It("surrounds a string", func() {
		Check(strings2.Surround("Go", "'"), Equals, "'Go'")
	})
})
