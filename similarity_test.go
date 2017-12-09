package animagi_test

import (
	"github.com/barreeyentos/animagi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Animagi", func() {

	mlFactor := 5
	lFactor := 1
	dFactor := 3

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

	Context("Invalid strings", func() {
		It("Should return MaxRank for '.'", func() {
			rank := animagi.SimilarityRank(".", "valid")
			Expect(rank).To(BeNumerically("==", animagi.MaxRank))
		})

		It("Should return MaxRank for ' ' as first string", func() {
			rank := animagi.SimilarityRank(" ", "valid")
			Expect(rank).To(BeNumerically("==", animagi.MaxRank))
		})

		It("Should return MaxRank for ' ' as second string", func() {
			rank := animagi.SimilarityRank("valid", " ")
			Expect(rank).To(BeNumerically("==", animagi.MaxRank))
		})

		It("Should return MaxRank for consecutive '.'", func() {
			rank := animagi.SimilarityRank("something..invalid", "valid")
			Expect(rank).To(BeNumerically("==", animagi.MaxRank))
		})

		It("Should return MaxRank for string with space", func() {
			rank := animagi.SimilarityRank("foo. .baar", "valid")
			Expect(rank).To(BeNumerically("==", animagi.MaxRank))
		})
	})

	Context("Strings of same depth", func() {
		It("Should return 0 for same strings", func() {
			rank := animagi.SimilarityRank("com.barreeyentos.animagi", "com.barreeyentos.animagi")
			Expect(rank).To(BeZero())
		})

		It("Should return 1 for strings off by one letter", func() {
			rank := animagi.SimilarityRank("com.barreeyentos.animag1", "com.barreeyentos.animagi")
			Expect(rank).To(BeNumerically("==", 1*lFactor))
		})
	})

	Context("Strings with different depths", func() {
		It("Should return dFactor*2 for same string off by 2 depths", func() {
			rank := animagi.SimilarityRank("animagi", "com.barreeyentos.animagi")
			Expect(rank).To(BeNumerically("==", 2*dFactor))
		})

		It("Should find smallest rank for same string off by 1 depths", func() {
			rank := animagi.SimilarityRank("com.animagi", "com.barreeyentos.animagi")
			Expect(rank).To(BeNumerically("==", 1*dFactor))
		})

		It("Should find smallest rank for same string off by 2 depths", func() {
			rank := animagi.SimilarityRank("user.employer.manager.name", "user.name")
			Expect(rank).To(BeNumerically("==", 2*dFactor))
		})

		It("Should find smallest rank for same string off by 3 depths", func() {
			rank := animagi.SimilarityRank("user.employer.manager.details.name", "manager.name")
			Expect(rank).To(BeNumerically("==", 3*dFactor))
		})

		It("Should find smallest rank for same string off by 3 depths and wrong letters", func() {
			rank := animagi.SimilarityRank("user.employer.maneger.details.name", "manager.name")
			Expect(rank).To(BeNumerically("==", 3*dFactor+1*lFactor))
		})

		It("Should find smallest rank for same string off by 3 depths and wrong letter and extra letter", func() {
			rank := animagi.SimilarityRank("user.employer.manegers.details.name", "manager.name")
			Expect(rank).To(BeNumerically("==", 3*dFactor+1*lFactor+1*mlFactor))
		})
	})
})
