package main

import "fmt"

// Person encapsulates personal data fields like name and age.
type Person struct {
	name string
	age  int
}

// NewPerson is a constructor function to create a new Person instance.
func NewPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

// GetName returns the name of the person.
func (p *Person) GetName() string {
	return p.name
}

// SetName sets the name of the person.
func (p *Person) SetName(name string) {
	p.name = name
}

// GetAge returns the age of the person.
func (p *Person) GetAge() int {
	return p.age
}

// SetAge sets the age of the person.
func (p *Person) SetAge(age int) {
	p.age = age
}

// Worker defines an interface for worker behaviors.
type Worker interface {
	Work() string
}

// Engineer represents an engineer with a specialization.
type Engineer struct {
	Person         // Embeds Person struct (is-a relationship).
	specialization string
}

// Work returns the type of engineering work the engineer is doing.
func (e Engineer) Work() string {
	return e.specialization + " engineering work in progress..."
}

// Manager represents a manager with a department to manage.
type Manager struct {
	Person     // Embeds Person struct (is-a relationship).
	department string
}

// Work returns the type of work the manager is doing.
func (m Manager) Work() string {
	return "Managing the " + m.department + " department..."
}

func main() {
	// Creating instances of Engineer and Manager.
	engineer := Engineer{
		Person:         *NewPerson("Alice", 30),
		specialization: "Software",
	}

	manager := Manager{
		Person:     *NewPerson("Bob", 45),
		department: "IT",
	}

	// Using polymorphism: both Engineer and Manager implement the Worker interface.
	var worker Worker

	worker = engineer
	fmt.Println(worker.Work()) // Output: Software engineering work in progress...

	worker = manager
	fmt.Println(worker.Work()) // Output: Managing the IT department...
}