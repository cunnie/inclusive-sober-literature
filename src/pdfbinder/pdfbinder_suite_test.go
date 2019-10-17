package pdfbinder_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPdfbinder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pdfbinder Suite")
}
