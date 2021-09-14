package interpreter

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Interpreter", func() {
	Context("evalE called with (Plus {x, y})", func() {
		It("Should return 6", func(){
			Expect(evalE(Plus{Variable{"x"},
				Variable{"y"},
			},
				Store{map[string]int{
					"x": 2,
					"y": 4,
				}})).To(Equal(6))
		})
	})

	Context("Run called with Assignment (x := 4)", func() {
		It("Should return a Store with x: 4", func(){
			Expect(Run(Assignment{"x", Const{Value{MyInt,4,false}}},
				Store{map[string]int{}})).To(Equal(Store{innerMap: map[string]int{"x": 4}}))
		})
	})

	Context("Run called with SequentialComposition (x := 4) (y := 12)", func() {
		It("Should return a Store with x:4, y:12", func(){
			Expect(Run(SequentialComposition{
				Assignment{"x", Const{Value{MyInt,4, false}}},
				Assignment{"y", Const{Value{MyInt,12, false}}}},
				Store{map[string]int{}})).To(Equal(Store{innerMap: map[string]int{"x": 4, "y":12}}))
		})
	})

	Context("Run called with Conditional (0 then x := 4 else y := 12)", func() {
		It("Should return a Store with x:4", func() {
			Expect(Run(Conditional{
				Const{Value{MyInt,0, false}},
				Assignment{"x", Const{Value{MyInt,4, false}}},
				Assignment{"y", Const{Value{MyInt,12, false}}}},
				Store{map[string]int{}})).To(Equal(Store{innerMap: map[string]int{"x": 4}}))
		})
	})

})
