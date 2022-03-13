package config

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GlobalConfig", func() {
	It("Global config should be defaulted when no flags passed", func() {
		ProcessArgs()
		Expect(GlobalConfig.Namespace).To(Equal("redhat-nvidia-gpu"))
		Expect(GlobalConfig.GpuPrefix).To(Equal("nvidia-gpu-addon"))
		Expect(GlobalConfig.NfdPrefix).To(Equal("node-feature-discovery-operator"))
	})
})

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}
