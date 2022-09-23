package chain

import (
	"math/big"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utils Suite")
}

var _ = Describe("check utils", Label("utils"), func() {
	When("string", func() {
		Context("gwei string to wei", func() {
			It("string", func() {
				amount := ToWei("9", 9)
				log.Info(amount)
				Expect(amount).To(Equal(big.NewInt(9000000000)))
			})
		})

		Context("wei string to wei", func() {
			It("string", func() {
				amount := ToWei("9", 0)
				log.Info(amount)
				Expect(amount).To(Equal(big.NewInt(9)))
			})
		})
	})
})
