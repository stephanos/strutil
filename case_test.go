package strings2_test

import (
	. "github.com/101loops/bdd"
	"github.com/101loops/strings2"
)

var _ = Describe("casing", func() {

	It("lowers first char", func() {
		Check(strings2.LowerFirst(""), Equals, "")
		Check(strings2.LowerFirst("GoLang"), Equals, "goLang")
		Check(strings2.LowerFirst("!golang"), Equals, "!golang")
	})

	It("uppers first char", func() {
		Check(strings2.UpperFirst(""), Equals, "")
		Check(strings2.UpperFirst("goLang"), Equals, "GoLang")
		Check(strings2.UpperFirst("!golang"), Equals, "!golang")
	})
})
