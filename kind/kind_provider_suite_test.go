package kind

import (
    "testing"

    "github.com/yourname/ccm-contract-tests/suite"
    kindprovider "sigs.k8s.io/cloud-provider-kind/pkg/cloudprovider"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestKindProvider(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Kind Cloud Provider Suite")
}

var _ = Describe("Kind Cloud Provider", func() {
    var provider *kindprovider.Cloud

    BeforeSuite(func() {
        var err error
        provider, err = kindprovider.NewCloud("kind", nil)
        Expect(err).NotTo(HaveOccurred())
    })

    suite.RunInstancesV2Suite(provider, "kind-control-plane")
})

