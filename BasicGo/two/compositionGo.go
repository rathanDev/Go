// Go does not have classical inheritance like Java or C++, 
// where a subclass extends a parent class and inherits its methods and fields. 
// Instead, Go achieves similar behavior using composition and interfaces.

// 1. Composition (Struct Embedding)
// Instead of inheritance, Go allows you to embed one struct inside another. 
// The embedded struct's methods and fields are accessible as if they were part of the outer struct.

package two

import "fmt"

// Base struct
type Animal struct {
    Name string
}

// Method for Animal
func (a Animal) Speak() {
    fmt.Println("Some generic animal sound")
}

// Derived struct using embedding
type Dog struct {
    Animal // Embedded struct (Composition)
    Breed  string
}

// Overriding behavior
func (d Dog) Speak() {
    fmt.Println("Woof!")
}

func main() {
    dog := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }

    dog.Speak()  // Calls Dog's Speak() → "Woof!"
    dog.Animal.Speak() // Calls Animal's Speak() → "Some generic animal sound"
}

// 2. Interfaces (Polymorphism)
// Go uses interfaces to achieve polymorphism, allowing different types to implement the same behavior.

// package mainC2

// import "fmt"

// Define an interface
type Speaker interface {
    Speak()
}

// Implementing the interface
type Cat struct{}

func (c Cat) Speak() {
    fmt.Println("Meow!")
}

func mainNot() {
    var s Speaker
    s = Cat{}
    s.Speak() // Output: Meow!
}

// Summary:
// - No traditional class-based inheritance
// - Composition (struct embedding) is used instead of subclassing
// - Interfaces enable polymorphism

// This design avoids deep inheritance trees and promotes explicit over implicit behavior, making Go code easier to maintain and understand. 🚀