package strutil_test

import (
	. "github.com/101loops/bdd"
	"github.com/101loops/strutil"
)

var _ = Describe("String Utility", func() {

	It("deletes all spaces", func() {
		Check(strutil.DeleteWhitespace(""), Equals, "")
		Check(strutil.DeleteWhitespace("abc"), Equals, "abc")
		Check(strutil.DeleteWhitespace(" a  b "), Equals, "ab")
		Check(strutil.DeleteWhitespace("a \n \t b"), Equals, "ab")
	})

	It("trims all spaces", func() {
		Check(strutil.ShrinkWhitespace(""), Equals, "")
		Check(strutil.ShrinkWhitespace(" "), Equals, "")
		Check(strutil.ShrinkWhitespace(" test "), Equals, "test")
		Check(strutil.ShrinkWhitespace(" te  st "), Equals, "te st")
		Check(strutil.ShrinkWhitespace(" te\tst "), Equals, "te st")
		Check(strutil.ShrinkWhitespace(" te\nst "), Equals, "te st")
	})

	It("reverses string", func() {
		Check(strutil.Reverse(""), Equals, "")
		Check(strutil.Reverse("abc"), Equals, "cba")
	})

	It("checks if empty", func() {
		Check(strutil.IsEmpty(""), IsTrue)

		Check(strutil.IsEmpty(" "), IsFalse)
		Check(strutil.IsEmpty("abc"), IsFalse)
	})

	It("checks if blank", func() {
		Check(strutil.IsBlank(""), IsTrue)
		Check(strutil.IsBlank("\n"), IsTrue)
		Check(strutil.IsBlank("\t"), IsTrue)
		Check(strutil.IsBlank(" "), IsTrue)

		Check(strutil.IsBlank("abc"), IsFalse)
		Check(strutil.IsBlank(" abc "), IsFalse)
	})

	It("returns index after substring", func() {
		Check(strutil.IndexAfter("", ""), Equals, 0)
		Check(strutil.IndexAfter("", "a"), Equals, -1)
		Check(strutil.IndexAfter("abc", "ab"), Equals, 2)
		Check(strutil.IndexAfter("abc", "bc"), Equals, 3)
		Check(strutil.IndexAfter("abc", "xyz"), Equals, -1)
	})

	It("returns lines", func() {
		Check(strutil.Lines(""), Equals, []string{""})
		Check(strutil.Lines("a"), Equals, []string{"a"})
		Check(strutil.Lines("a\nb\nc"), Equals, []string{"a", "b", "c"})
	})

	It("wraps text to fixed width", func() {
		Check(strutil.WordWrap(``, 5), Equals, ``)
		//Check(WordWrap(`abcdef`, 5), Equals, `abcdef`)
		//Check(WordWrap(`abcdefghijklmno`, 5), Equals, `abcdef`)
	})

	It("surrounds a string", func() {
		Check(strutil.Surround("Go", "'"), Equals, "'Go'")
	})
})
