package settings

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

func TestTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("check json load local bridge settings", Label("settings"), func() {
	When("local", func() {
		Context("json load local bridge settings", func() {
			It("local", func() {
				var bridgeSettings Settings
				file, err := ioutil.ReadFile("bridge_settings_1.json")
				Expect(err).To(BeNil())
				err = json.Unmarshal([]byte(file), &bridgeSettings)
				log.Infof("%+v\n", bridgeSettings)
				Expect(err).To(BeNil())
				Expect(bridgeSettings.NonCritical.ChainlinkDxUsdFeedAddress.Hex()).To(Equal("0xE1329B3f6513912CAf589659777b66011AEE5880"))
				Expect(bridgeSettings.NonCritical.MinimumConfirmations.Dxchain).To(BeEquivalentTo(1))
			})
		})
	})
})
