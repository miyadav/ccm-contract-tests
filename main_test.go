package main

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestCCMContractTests(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "CCM Contract Test Suite")
}

