package strutil_test

import (
	. "github.com/101loops/bdd"
	"github.com/101loops/strutil"
)

var _ = Describe("conjuction", func() {

	It("creates a conjuction with 'and'", func() {
		Check(strutil.And(), Equals, "")
		Check(strutil.And("one"), Equals, "one")
		Check(strutil.And("one", "two"), Equals, "one and two")
		Check(strutil.And("one", "two", "three"), Equals, "one, two and three")
		Check(strutil.And("one", "two", "three", "four"), Equals, "one, two, three and four")
	})

	It("converts a conjuction with 'or'", func() {
		Check(strutil.Or(), Equals, "")
		Check(strutil.Or("one"), Equals, "one")
		Check(strutil.Or("one", "two"), Equals, "one or two")
		Check(strutil.Or("one", "two", "three"), Equals, "one, two or three")
		Check(strutil.Or("one", "two", "three", "four"), Equals, "one, two, three or four")
	})
})
