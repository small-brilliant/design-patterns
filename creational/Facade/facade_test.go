package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	hp := &HomePlayerFacade{}
	hp.back()
	fmt.Println("------------")
	hp.leave()
}
