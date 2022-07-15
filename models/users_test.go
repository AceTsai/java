package models

import (
	"fmt"
	"testing"
)

func TestGetOne(t *testing.T) {
	x := new(Users)
	fmt.Println(x.Find(1))
}
