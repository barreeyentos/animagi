package animagi_test

import (
	"github.com/barreeyentos/animagi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type myint int
type mystring string

type SimpleSingleDepth struct {
	T_int   int
	T_int8  int8
	T_int16 int16
	T_int32 int32
	T_myint myint

	T_string   string
	T_mystring mystring
}

type TheSameSimpleSingleDepth struct {
	T_int   int
	T_int8  int8
	T_int16 int16
	T_int32 int32
	T_myint myint

	T_string   string
	T_mystring mystring

	T_extraInt int
	T_extra    string
}
type SimpleWithDepthOfTwo struct {
	Description     string
	SameSingleDepth TheSameSimpleSingleDepth
}
type TheSameSimpleWithDepthOfTwo struct {
	Description           string
	ExtraDescription      string
	SameSingleDepth       TheSameSimpleSingleDepth
	SameTypeDifferentName TheSameSimpleSingleDepth
}

var _ = Describe("AnimagiBasicMapping", func() {

	Context("Single Depth Structs", func() {
		It("Should Map Simple Structs", func() {
			src := SimpleSingleDepth{1, 127, 32767, 512, 1024, "animagi", "animato"}
			var dst TheSameSimpleSingleDepth
			err := animagi.Transform(src, &dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst).NotTo(BeEquivalentTo(src))

			Expect(dst.T_extra).To(BeEmpty())
			Expect(dst.T_extraInt).To(BeZero())
			Expect(dst.T_int).To(Equal(src.T_int))
			Expect(dst.T_int8).To(Equal(src.T_int8))
			Expect(dst.T_int16).To(Equal(src.T_int16))
			Expect(dst.T_int32).To(Equal(src.T_int32))
			Expect(dst.T_string).To(Equal(src.T_string))
			Expect(dst.T_mystring).To(Equal(src.T_mystring))

		})

		It("Should Map Structs with Fields of Same Kind", func() {
			src := struct {
				A int
				B mystring
			}{420, "i'm a string"}
			var dst struct {
				A myint
				B string
			}

			err := animagi.Transform(src, &dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.A).To(BeNumerically("==", src.A))
			Expect(dst.B).To(BeEquivalentTo(src.B))

		})
	})

	Context("Depth of Two Structs", func() {
		It("Should Map Simple Structs", func() {
			src := SimpleWithDepthOfTwo{"two", TheSameSimpleSingleDepth{1, 127, 32767, 512, 1024, "animagi", "animato", 42, "extra"}}
			var dst TheSameSimpleWithDepthOfTwo
			err := animagi.Transform(src, &dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst).NotTo(BeEquivalentTo(src))

			Expect(dst.Description).To(Equal(src.Description))
			Expect(dst.ExtraDescription).To(BeEmpty())

			Expect(dst.SameSingleDepth.T_extra).To(Equal(src.SameSingleDepth.T_extra))
			Expect(dst.SameSingleDepth.T_extraInt).To(Equal(src.SameSingleDepth.T_extraInt))
			Expect(dst.SameSingleDepth.T_int).To(Equal(src.SameSingleDepth.T_int))
			Expect(dst.SameSingleDepth.T_int8).To(Equal(src.SameSingleDepth.T_int8))
			Expect(dst.SameSingleDepth.T_int16).To(Equal(src.SameSingleDepth.T_int16))
			Expect(dst.SameSingleDepth.T_int32).To(Equal(src.SameSingleDepth.T_int32))
			Expect(dst.SameSingleDepth.T_string).To(Equal(src.SameSingleDepth.T_string))
			Expect(dst.SameSingleDepth.T_mystring).To(Equal(src.SameSingleDepth.T_mystring))

			Expect(dst.SameTypeDifferentName.T_extra).To(BeZero())
			Expect(dst.SameTypeDifferentName.T_extraInt).To(BeZero())
			Expect(dst.SameTypeDifferentName.T_int).To(BeZero())
			Expect(dst.SameTypeDifferentName.T_int8).To(BeZero())
			Expect(dst.SameTypeDifferentName.T_int16).To(BeZero())
			Expect(dst.SameTypeDifferentName.T_int32).To(BeZero())
			Expect(dst.SameTypeDifferentName.T_string).To(BeEmpty())
			Expect(dst.SameTypeDifferentName.T_mystring).To(BeEmpty())

		})
	})
})
