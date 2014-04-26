package strings2_test

import (
	. "github.com/101loops/bdd"
	"github.com/101loops/strings2"
)

var _ = Describe("conversion", func() {

	It("converts string to int", func() {
		i, err := strings2.ToInt("32")
		Check(err, IsNil)
		Check(i, Equals, 32)

		i, err = strings2.ToInt("32.5")
		Check(err, NotNil)

		i, err = strings2.ToInt("a")
		Check(err, NotNil)
	})

	It("converts int to string", func() {
		Check(strings2.FromInt(32), Equals, "32")
	})
})
