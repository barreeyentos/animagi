package animagi_test

import (
	"github.com/barreeyentos/animagi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Animagi", func() {

	mlFactor := 5
	lFactor := 1

	Context("single depth strings", func() {
		It("Should handle 1st param as empty strings", func() {
			rank := animagi.SimilarityRank("", "hello")
			Expect(rank).To(BeNumerically("==", mlFactor*5))
		})

		It("Should handle 2nd param as empty strings", func() {
			rank := animagi.SimilarityRank("hello", "")
			Expect(rank).To(BeNumerically("==", mlFactor*5))
		})

		It("Should handle exactly same strings", func() {

			rank := animagi.SimilarityRank("SuperCalifragilisticExpialidocious", "SuperCalifragilisticExpialidocious")
			Expect(rank).To(BeZero())
		})

		It("Should handle extra letters in second string", func() {

			rank := animagi.SimilarityRank("oneextra", "oneextra1")
			Expect(rank).To(BeNumerically("==", mlFactor*1))
		})

		It("Should handle extra letters in first string", func() {
			rank := animagi.SimilarityRank("oneextra1", "oneextra")
			Expect(rank).To(BeNumerically("==", mlFactor*1))
		})

		It("Should handle wrong letters", func() {
			rank := animagi.SimilarityRank("onewrong", "on3wrong")
			Expect(rank).To(BeNumerically("==", lFactor*1))
		})

		It("Should handle many wrong letters", func() {
			rank := animagi.SimilarityRank("th3s3ar3wr0ng", "thesearewrong")
			Expect(rank).To(BeNumerically("==", lFactor*4))
		})

		It("Should handle many wrong letters and extras", func() {
			rank := animagi.SimilarityRank("th3s3ar3wr0ng", "thesearewrong31")
			Expect(rank).To(BeNumerically("==", lFactor*4+mlFactor*2))
		})
	})

	FContext("Invalid strings", func() {
		It("Should return MaxRank for '.'", func() {
			rank := animagi.SimilarityRank(".", "valid")
			Expect(rank).To(BeNumerically("==", animagi.MaxRank))
		})

		It("Should return MaxRank for ' '", func() {
			rank := animagi.SimilarityRank(" ", "valid")
			Expect(rank).To(BeNumerically("==", animagi.MaxRank))
		})

		It("Should return MaxRank for consecutive '.'", func() {
			rank := animagi.SimilarityRank("something..invalid", "valid")
			Expect(rank).To(BeNumerically("==", animagi.MaxRank))
		})
	})

	Context("Strings of same depth", func() {

	})
})
