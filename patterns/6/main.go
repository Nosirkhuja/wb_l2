package pattern

import "fmt"

type Class interface {
	getClassName() string
}

type Assassin struct {
	className string
	charName  string
}

func newAssasin(name string) *Assassin {
	return &Assassin{
		className: "Assassin",
		charName:  name,
	}
}

func (a *Assassin) getClassName() string {
	return a.className
}

type Warrior struct {
	className string
	charName  string
}

func newWarrior(name string) *Warrior {
	return &Warrior{
		className: "Warrior",
		charName:  name,
	}
}

func (w *Warrior) getClassName() string {
	return w.className
}

func getClass(className, charName string) (Class, error) {
	switch className {
	case "Warrior":
		return newWarrior(charName), nil
	case "Assassin":
		return newAssasin(charName), nil
	}
	return nil, fmt.Errorf("Wrong class name passed")
}

/*
 *	warrior, _ := getClass("Warrior", "Some name")
 *	assassin, _ := getClass("Assassin", "new Name")
 *
 *	fmt.Printf("Class name: %s\n", warrior.className)
 *	fmt.Printf("Class name: %s\n", assassin.className)
 */
