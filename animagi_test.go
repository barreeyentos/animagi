package animagi_test

import (
	. "github.com/barreeyentos/animagi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MySimpleStruct struct {
	aString string
}

type MyTestStruct struct {
	anInteger     int
	aString       string
	aSimpleStruct MySimpleStruct
}

type MyInt32Kind int32

var _ = Describe("Animagi", func() {

	Context("Unsettable destinations", func() {
		It("Should return error when dst is not a ref", func() {
			src := 31
			var dst int
			err := Transform(src, dst)
			Expect(err).Should(HaveOccurred())
		})

		It("Should return error when dst is nil", func() {
			src := 23
			err := Transform(src, nil)
			Expect(err).Should(HaveOccurred())
		})

		It("Should return error when dst is a literal", func() {
			src := 353
			err := Transform(src, "nil")
			Expect(err).Should(HaveOccurred())
		})

		It("Should return error when dst is a struct", func() {
			src := MyTestStruct{42, "someting", MySimpleStruct{"somethingElse"}}
			dst := MyTestStruct{}
			err := Transform(src, dst)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("Types are exactly the same", func() {
		It("Should copy int64", func() {
			var src int64 = 332
			var dst int64
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy int32", func() {
			var src int32 = -193
			var dst int32
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy int16", func() {
			var src int16 = 83
			var dst int16
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy int8", func() {
			var src int8 = 87
			var dst int8
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy uint64", func() {
			var src uint64 = 6
			var dst uint64
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy uint32", func() {
			var src uint32 = 998
			var dst uint32
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy uint16", func() {
			var src uint16 = 420
			var dst uint16
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy uint8", func() {
			var src uint8 = 31
			var dst uint8
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy float64", func() {
			var src float64 = 3.14312
			var dst float64
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy float32", func() {
			var src float32 = 3.14
			var dst float32
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy bools", func() {
			var src = true
			var dst bool
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy strings", func() {
			var src = "a string"
			var dst string
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy array of strings", func() {
			var src = []string{"a string", "and another"}
			var dst []string
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy simple structs", func() {
			var src = MySimpleStruct{"a string"}
			var dst MySimpleStruct
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy nested structs", func() {
			var src = MyTestStruct{42, "a string", MySimpleStruct{"inner struct"}}
			var dst MyTestStruct
			_ = Transform(src, &dst)
			Expect(dst).To(Equal(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})
	})

	Context("Types are same Kind", func() {
		It("Should copy kind of int", func() {
			var src MyInt32Kind = 332
			var dst int32
			_ = Transform(src, &dst)
			Expect(dst).To(BeEquivalentTo(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})

		It("Should copy kind of MyInt32Kind", func() {
			var src int32 = 332
			var dst MyInt32Kind
			_ = Transform(src, &dst)
			Expect(dst).To(BeEquivalentTo(src))
			Expect(&dst).NotTo(BeIdenticalTo(&src))
		})
	})
})
