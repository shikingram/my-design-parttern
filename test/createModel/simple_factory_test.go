package createModel

import (
	"my-design-parttern/createModel/simple_factory"
	"testing"
)

func TestNewTesla(t *testing.T) {
	var expect = "tesla"
	var teslaNo1 = simple_factory.NewCar(expect)
	if teslaNo1 == nil {
		t.Fatal("TestNewTesla() failed,teslaNo1 is nil")
		return
	}
	var typeTeslaNo1 = teslaNo1.GetCarType()
	if expect != typeTeslaNo1 {
		t.Fatalf("TestNewTesla() failed,expect : %s ,actually : %s", expect, typeTeslaNo1)
	}
}
