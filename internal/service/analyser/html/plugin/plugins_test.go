package plugin

import "testing"

func Temp(url string) ProcessorInterface {
	return nil
}
func TestRegisterProcessorUnique(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function paniced for unique keys")
		}
	}()
	RegisterProcessor("here", Temp)
	RegisterProcessor("here2", Temp)
	if len(Processors) != 2 {
		t.Error("Processors were not registerd")
	}
}

func TestRegisterProcessorNotUnique(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Function paniced for unique keys")
		}
	}()
	RegisterProcessor("here", Temp)
	RegisterProcessor("here", Temp)
}
