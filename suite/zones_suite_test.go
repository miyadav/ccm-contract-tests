package suite

import (
    "context"

    "k8s.io/cloud-provider"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func RunZonesSuite(provider cloudprovider.Interface) {
    Describe("Cloud Provider Zones Contract", func() {
        It("should retrieve zone information", func() {
            zones, err := provider.Zones().GetZone(context.TODO())
            Expect(err).NotTo(HaveOccurred())
            Expect(zones.FailureDomain).NotTo(BeEmpty())
            Expect(zones.Region).NotTo(BeEmpty())
        })
    })
}

