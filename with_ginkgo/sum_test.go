package testing_ginkgo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sajanjswl/testing-go/with_ginkgo"
)

var _ = Describe("Sum", func() {

	var (
		p, q, m, n, sum1, sum2 int
	)
	BeforeEach(func() {
		p, q, sum1 = 5, 6, 11
		// Putting wrong value of sum2 intentionally
		m, n, sum2 = 8, 7, 16
	})
	Context("Addition of two digits", func() {
		It("should return sum of the two digits", func() {
			addition_of_two_digits := Sum(p, q)
			Expect(addition_of_two_digits).Should(Equal(sum1))
		})
		It("should not return the sum provided", func() {
			addition_of_two_digits := Sum(m, n)
			Expect(addition_of_two_digits).ShouldNot(Equal(sum2))
		})
	})

	var (
		a, b, c, d, sub1, sub2 int
	)

	BeforeEach(func() {
		a, b, sub1 = 6, 5, 1
		d, c, sub2 = 7, 2, 2

	})

	Context("Subtracton of two digits", func() {

		It("should return subtraction of two digits", func() {
			subtraction_of_two_digits := Sub(a, b)
			Expect(subtraction_of_two_digits).Should(Equal(sub1))
		})

		It("should not return the subtraction provide", func() {
			subtraction_of_two_digits := Sub(d, c)
			Expect(subtraction_of_two_digits).ShouldNot(Equal(sub2))

		})

	})

})
