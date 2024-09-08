package main

// Mediator some step can be done concurrently, for the purpose of showing how mediator works, we make it step by step
type Mediator struct {
	ProductManager *ProductManager
	FrontEnd       *Developer
	BackEnd        *Developer
	Test           *TestDeveloper

	frontFinishTecDoc bool
	backFinishTecDoc  bool

	frontFinishedDevelop bool
	backFinishedDevelop  bool

	frontReleasedToTest bool
	backReleasedToTest  bool

	frontReleasedToOnline bool
	backReleasedToOnline  bool
}

func (m *Mediator) RequirementReview() {
	m.Test.ReceiveRequirementDocument()
	m.FrontEnd.ReceiveRequirementDocument()
	m.BackEnd.ReceiveRequirementDocument()
}

func (m *Mediator) ReviewTechnicalDevelopDocument(isFront bool) {
	if isFront {
		m.frontFinishTecDoc = true
	} else {
		m.backFinishTecDoc = true
	}

	if m.frontFinishTecDoc && m.backFinishTecDoc {
		m.Test.WriteTestPlan()
		m.ProductManager.ReviewTechnicalDevelopDocument()
	}
}

func (m *Mediator) BeginDevelop() {
	m.FrontEnd.MakeDevelopPlan()
	m.BackEnd.MakeDevelopPlan()

	m.FrontEnd.Develop()
	m.BackEnd.Develop()
}

func (m *Mediator) FinishedDevelop(isFront bool) {
	if isFront {
		m.frontFinishedDevelop = true
	} else {
		m.backFinishedDevelop = true
	}

	if m.frontFinishedDevelop && m.backFinishedDevelop {
		m.FrontEnd.ReleaseToTest()
		m.BackEnd.ReleaseToTest()
	}

}

func (m *Mediator) ReleasedToTest(isFront bool) {
	if isFront {
		m.frontReleasedToTest = true
	} else {
		m.backReleasedToTest = true
	}

	if m.frontReleasedToTest && m.backReleasedToTest {
		m.Test.Test()
	}
}

func (m *Mediator) TestPassed() {
	m.ProductManager.Showcase()
}

func (m *Mediator) ShowcasePassed() {
	m.FrontEnd.ReleaseToOnline()
	m.BackEnd.ReleaseToOnline()
}

func (m *Mediator) ReleasedToOnline(isFront bool) {
	if isFront {
		m.frontReleasedToOnline = true
	} else {
		m.backReleasedToOnline = true
	}

	if m.frontReleasedToOnline && m.backReleasedToOnline {
		m.Test.RegressionTest()
	}
}

func (m *Mediator) RegressionTestPassed() {
	m.FrontEnd.RegressionTestPassed()
	m.BackEnd.RegressionTestPassed()
	m.ProductManager.Summary()
}
