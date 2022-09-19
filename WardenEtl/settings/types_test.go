package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("check json load local bridge settings", Label("settings"), func() {
	When("local", func() {
		Context("json load local bridge settings", func() {
			It("local", func() {

				spaceClient := http.Client{
					Timeout: time.Second * 2,
				}
				url := "http://127.0.0.1:8050/static/heco_bridge_settings_1.json"
				req, err := http.NewRequest(http.MethodGet, url, nil)
				if err != nil {
					log.Fatal(err)
				}
				res, err := spaceClient.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				if res.Body != nil {
					defer res.Body.Close()
				}
				body, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				var bridgeSettings Settings
				err = json.Unmarshal(body, &bridgeSettings)
				logrus.Infof("%+v\n", bridgeSettings)
				Expect(err).To(BeNil())
				Expect(bridgeSettings.NonCritical.ChainlinkDxUsdFeedAddress.Hex()).To(Equal("0xE1329B3f6513912CAf589659777b66011AEE5880"))
				Expect(bridgeSettings.NonCritical.MinimumConfirmations.Dxchain).To(BeEquivalentTo(1))
			})
		})
	})
})
