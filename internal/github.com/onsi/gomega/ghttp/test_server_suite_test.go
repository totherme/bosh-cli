package ghttp_test

import (
	. "github.com/cloudfoundry/bosh-init/internal/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry/bosh-init/internal/github.com/onsi/gomega"

	"testing"
)

func TestGHTTP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GHTTP Suite")
}
