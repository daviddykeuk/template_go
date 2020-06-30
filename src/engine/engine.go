package engine

import "fmt"

type Engine struct {
	dependency Dependency
}

// New...Returns a new Engine
func New(factory Factory) *Engine {
	return &Engine{
		dependency: factory.GetConcreteStruct("dependency").(Dependency),
	}
}

// Run...Runs the engine code with provided RunArguments
func (this *Engine) Run(args RunArguments) error {
	withThis := WithThis{
		property1: args.argument1,
	}
	produceThis, err := this.dependency.DoThis(withThis)

	fmt.Println(produceThis.output1)

	return err
}

// Factory...The factory interferface must return concrete structs of provided types
type Factory interface {
	GetConcreteStruct(string) interface{}
}

// RunRunArguments...The arguments required to run the function
type RunArguments struct {
	argument1 string
}

// Dependency...The expected dependency interface to perform a required task
type Dependency interface {
	DoThis(WithThis) (ToProduceThis, error)
}

// WithThis...The struct that the dependency will need to use
type WithThis struct {
	property1 string
}

// ToToProduceThis...The struct that the dependency will need to return
type ToProduceThis struct {
	output1 string
}
