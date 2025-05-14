package suite

import (
    "context"
    "fmt"

    v1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/cloud-provider"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func RunInstancesV2Suite(provider cloudprovider.Interface, nodeName string) {
    Describe("Cloud Provider InstancesV2 Contract", func() {
        var node *v1.Node

        BeforeEach(func() {
            node = &v1.Node{
                ObjectMeta: metav1.ObjectMeta{
                    Name: nodeName,
                },
            }
        })

        It("should confirm instance exists", func() {
            exists, err := provider.InstancesV2().InstanceExists(context.TODO(), node)
            Expect(err).NotTo(HaveOccurred())
            Expect(exists).To(BeTrue())
        })

        It("should retrieve instance metadata", func() {
            metadata, err := provider.InstancesV2().InstanceMetadata(context.TODO(), node)
            Expect(err).NotTo(HaveOccurred())
            Expect(metadata.ProviderID).NotTo(BeEmpty())
            Expect(metadata.NodeAddresses).NotTo(BeEmpty())
        })

        It("should confirm instance is not shutdown", func() {
            shutdown, err := provider.InstancesV2().InstanceShutdown(context.TODO(), node)
            Expect(err).NotTo(HaveOccurred())
            Expect(shutdown).To(BeFalse())
        })
    })
}

