package suite

import (
    "context"

    "k8s.io/cloud-provider"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func RunRoutesSuite(provider cloudprovider.Interface, clusterName string) {
    Describe("Cloud Provider Routes Contract", func() {
        It("should list routes", func() {
            routes, err := provider.Routes().ListRoutes(context.TODO(), clusterName)
            Expect(err).NotTo(HaveOccurred())
            Expect(routes).NotTo(BeNil())
        })

        // Additional tests for CreateRoute, DeleteRoute, etc.
    })
}

