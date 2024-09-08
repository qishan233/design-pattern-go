package main

import "fmt"

// =================== ProductManager ===================
type ProductManager struct {
	Name     string
	Mediator *Mediator
}

func GiveMeProductManager() *ProductManager {
	pm := &ProductManager{Name: "PM"}
	test := &TestDeveloper{Name: "Test", Mediator: nil}
	frontEnd := &Developer{Name: "FrontEnd", IsFront: true, Mediator: nil}
	backEnd := &Developer{Name: "BackEnd", IsFront: false, Mediator: nil}

	mediator := &Mediator{
		ProductManager: pm,
		Test:           test,
		FrontEnd:       frontEnd,
		BackEnd:        backEnd,
	}

	pm.Mediator = mediator
	test.Mediator = mediator
	frontEnd.Mediator = mediator
	backEnd.Mediator = mediator

	return pm
}

func (p *ProductManager) WriteRequirementDocument() {
	fmt.Println(p.Name, " write the requirement document")
	p.RequirementReview()
}

func (p *ProductManager) RequirementReview() {
	fmt.Println(p.Name, " reviewed the requirement document")
	p.Mediator.RequirementReview()
}

func (p *ProductManager) ReviewTechnicalDevelopDocument() {
	fmt.Println(p.Name, " reviewed the technical develop document")
	p.Mediator.BeginDevelop()
}

func (p *ProductManager) Showcase() {
	fmt.Println(p.Name, " project showcase passed")
	p.Mediator.ShowcasePassed()
}

func (p *ProductManager) Summary() {
	fmt.Println(p.Name, " summarized the project, it had made lots of money")
}

// =================== Developer ===================
type Developer struct {
	Name     string
	IsFront  bool
	Mediator *Mediator
}

func (d *Developer) ReceiveRequirementDocument() {
	fmt.Println(d.Name, " received the requirement document")
	d.FinishTechnicalDevelopDocument()
}

func (d *Developer) FinishTechnicalDevelopDocument() {
	fmt.Println(d.Name, " finished the technical develop document")
	d.Mediator.ReviewTechnicalDevelopDocument(d.IsFront)
}

func (d *Developer) MakeDevelopPlan() {
	fmt.Println(d.Name, " make the development plan")
}

func (d *Developer) Develop() {
	fmt.Println(d.Name, " developed the project")
	d.FinishedDevelop()
}

func (d *Developer) FinishedDevelop() {
	fmt.Println(d.Name, " finished the development work")
	d.Mediator.FinishedDevelop(d.IsFront)
}

func (d *Developer) ReleaseToTest() {
	fmt.Println(d.Name, " release the project to test")
	d.Mediator.ReleasedToTest(d.IsFront)
}

func (d *Developer) ReleaseToOnline() {
	fmt.Println(d.Name, " release the project to online")
	d.Mediator.ReleasedToOnline(d.IsFront)
}

func (d *Developer) RegressionTestPassed() {
	fmt.Println(d.Name, " finished all work, project is online")
}

// =================== TestDeveloper ===================

type TestDeveloper struct {
	Name     string
	Mediator *Mediator
}

func (t *TestDeveloper) ReceiveRequirementDocument() {
	fmt.Println(t.Name, " received the requirement document")
}

func (t *TestDeveloper) WriteTestPlan() {
	fmt.Println(t.Name, " write the test plan")
}

func (t *TestDeveloper) Test() {
	fmt.Println(t.Name, " tested the project")
	t.TestPassed()
}

func (t *TestDeveloper) TestPassed() {
	fmt.Println(t.Name, " finished the test work, result is passed")
	t.Mediator.TestPassed()
}
func (t *TestDeveloper) RegressionTest() {
	fmt.Println(t.Name, " doing the regression test")
	t.RegressionTestPassed()
}

func (t *TestDeveloper) RegressionTestPassed() {
	fmt.Println(t.Name, " finished the regression test work, result is passed")
	t.Mediator.RegressionTestPassed()
}
