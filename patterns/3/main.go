package main

import "fmt"

func main() {

	backend := Developer{"Baggins", "Bilbo", 1000, 32}
	boss := Director{"Alexandr", "Konovalov", 2000, 25}

	backend.FullName()
	backend.Accept(CalculIncome{20})
	backend.Accept(AddingCaptainAge{0})

	boss.FullName()
	boss.Accept(CalculIncome{10})
	boss.Accept(AddingCaptainAge{1})

}

type Employee interface {
	FullName()
	Accept(Visitor)
}

type Developer struct {
	FirstName string
	LastName  string
	Income    int
	Age       int
}

func (d Developer) FullName() {
	fmt.Println("Developer ", d.FirstName, " ", d.LastName)
}
func (d Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}

type Director struct {
	FirstName string
	LastName  string
	Income    int
	Age       int
}

func (d Director) FullName() {
	fmt.Println("Director ", d.FirstName, " ", d.LastName)
}
func (d Director) Accept(v Visitor) {
	v.VisitDirector(d)
}

type Visitor interface {
	VisitDeveloper(d Developer)
	VisitDirector(d Director)
}

type CalculIncome struct {
	bonusRate int
}

func (c CalculIncome) VisitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}

func (c CalculIncome) VisitDirector(d Director) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}

type AddingCaptainAge struct {
	captainAge int
}

func (c AddingCaptainAge) VisitDeveloper(d Developer) {
	fmt.Println(d.Age + c.captainAge)
}

func (c AddingCaptainAge) VisitDirector(d Director) {
	fmt.Println(d.Age + c.captainAge)
}
