package pdfbinder_test

import (
	"github.com/alcoano/gender-neutral-aa-lit/src/pdfbinder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("pdfbinder", func() {
	// Bind(io.File[] inputMarkdownFiles) io.File

	Describe("Bind", func() {
		It("returns hello", func() {

			Expect(pdfbinder.Bind()).To(Equal("yellow"))
		})
	})
})
