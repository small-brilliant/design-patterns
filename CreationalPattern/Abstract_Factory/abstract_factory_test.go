package abstractfactory

import "testing"

func TestNewSimpleLuchFactory(t *testing.T) {
	appfc := new(AppleFactory)
	appfc.CreateFruit().Show()
}
