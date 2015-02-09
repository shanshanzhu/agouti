package dsl_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sclevine/agouti/dsl"
	"github.com/sclevine/agouti/dsl/internal/mocks"
)

var _ = Describe("Selection", func() {
	var (
		selection   *mocks.Selection
		failMessage string
	)

	BeforeEach(func() {
		failMessage = ""
		InjectFail(func(message string, callerSkip ...int) {
			failMessage = message
			ExpectWithOffset(3, callerSkip[0]).To(Equal(2))
		})
		selection = &mocks.Selection{}
	})

	AfterEach(func() {
		InjectFail(Fail)
	})

	Describe(".SwitchToFrame", func() {
		It("should call selection.SwitchToFrame", func() {
			SwitchToFrame(selection)
			Expect(selection.SwitchToFrameCall.Called).To(BeTrue())
		})

		It("should fail when page.SwitchToFrame returns an error", func() {
			selection.SwitchToFrameCall.Err = errors.New("some error")
			SwitchToFrame(selection)
			Expect(failMessage).To(Equal("Agouti failure: some error"))
		})
	})

	Describe(".Click", func() {
		It("should call selection.Click", func() {
			Click(selection)
			Expect(selection.ClickCall.Called).To(BeTrue())
		})

		It("should fail when page.Click returns an error", func() {
			selection.ClickCall.Err = errors.New("some error")
			Click(selection)
			Expect(failMessage).To(Equal("Agouti failure: some error"))
		})
	})

	Describe(".DoubleClick", func() {
		It("should call selection.DoubleClick", func() {
			DoubleClick(selection)
			Expect(selection.DoubleClickCall.Called).To(BeTrue())
		})

		It("should fail when page.DoubleClick returns an error", func() {
			selection.DoubleClickCall.Err = errors.New("some error")
			DoubleClick(selection)
			Expect(failMessage).To(Equal("Agouti failure: some error"))
		})
	})

	Describe(".Fill", func() {
		It("should call selection.Fill", func() {
			Fill(selection, "some text")
			Expect(selection.FillCall.Text).To(Equal("some text"))
		})

		It("should fail when page.Fill returns an error", func() {
			selection.FillCall.Err = errors.New("some error")
			Fill(selection, "some text")
			Expect(failMessage).To(Equal("Agouti failure: some error"))
		})
	})

	Describe(".Check", func() {
		It("should call selection.Check", func() {
			Check(selection)
			Expect(selection.CheckCall.Called).To(BeTrue())
		})

		It("should fail when page.Check returns an error", func() {
			selection.CheckCall.Err = errors.New("some error")
			Check(selection)
			Expect(failMessage).To(Equal("Agouti failure: some error"))
		})
	})

	Describe(".Uncheck", func() {
		It("should call selection.Uncheck", func() {
			Uncheck(selection)
			Expect(selection.UncheckCall.Called).To(BeTrue())
		})

		It("should fail when page.Uncheck returns an error", func() {
			selection.UncheckCall.Err = errors.New("some error")
			Uncheck(selection)
			Expect(failMessage).To(Equal("Agouti failure: some error"))
		})
	})

	Describe(".Select", func() {
		It("should call selection.Select", func() {
			Select(selection, "some text")
			Expect(selection.SelectCall.Text).To(Equal("some text"))
		})

		It("should fail when page.Select returns an error", func() {
			selection.SelectCall.Err = errors.New("some error")
			Select(selection, "some text")
			Expect(failMessage).To(Equal("Agouti failure: some error"))
		})
	})

	Describe(".Submit", func() {
		It("should call selection.Submit", func() {
			Submit(selection)
			Expect(selection.SubmitCall.Called).To(BeTrue())
		})

		It("should fail when page.Submit returns an error", func() {
			selection.SubmitCall.Err = errors.New("some error")
			Submit(selection)
			Expect(failMessage).To(Equal("Agouti failure: some error"))
		})
	})
})
