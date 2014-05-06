package strutil_test

import (
	. "github.com/101loops/bdd"
	"github.com/101loops/strutil"
)

var _ = Describe("conversion", func() {

	It("converts string to int", func() {
		i, err := strutil.ToInt("32")
		Check(err, IsNil)
		Check(i, Equals, 32)

		i, err = strutil.ToInt("32.5")
		Check(err, NotNil)

		i, err = strutil.ToInt("a")
		Check(err, NotNil)
	})

	It("converts int to string", func() {
		Check(strutil.FromInt(32), Equals, "32")
	})
})
