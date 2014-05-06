package strutil_test

import (
	. "github.com/101loops/bdd"
	"github.com/101loops/strutil"
)

var _ = Describe("casing", func() {

	It("lowers first char", func() {
		Check(strutil.LowerFirst(""), Equals, "")
		Check(strutil.LowerFirst("GoLang"), Equals, "goLang")
		Check(strutil.LowerFirst("!golang"), Equals, "!golang")
	})

	It("uppers first char", func() {
		Check(strutil.UpperFirst(""), Equals, "")
		Check(strutil.UpperFirst("goLang"), Equals, "GoLang")
		Check(strutil.UpperFirst("!golang"), Equals, "!golang")
	})
})
