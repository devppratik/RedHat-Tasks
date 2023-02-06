package main_test

import (
	strman "strman/pkg"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Concat", func() {
	var (
		strs []string
		res  string
	)
	BeforeEach(func() {
		strs = append(strs, "test1", "test2")
		res = "test1+test2"
	})
	Context("Concat of 2 Strings", func() {
		It("should return sum of the two digits", func() {
			concatenated_string := strman.ConcatStrings(strs, "+")
			Expect(concatenated_string).Should(Equal(res))
		})
	})
	Context("Concat of 2 Strings", func() {
		It("should return sum of the two digits", func() {
			concatenated_string := strman.ConcatStrings(strs, "")
			Expect(concatenated_string).ShouldNot(Equal(res))
		})
	})
})
