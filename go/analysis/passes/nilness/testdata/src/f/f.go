package f

import "fmt"

func TriggerBug(a *struct{}) {
	hasA := a != nil
	if hasA {
		fmt.Println("A is not nil")
		return
	}
	if !hasA {
		fmt.Println("A is nil")
	}
}

func TriggerTautological(a *struct{}) {
	if a != nil {
		fmt.Println("A is not nil")
		return
	}
	if a == nil { // want "tautological condition: nil == nil"
		fmt.Println("A is nil")
	}
}
