package basemongo

import (
	"fmt"
	"testing"
)

func TestGetClient(t *testing.T) {
	InitTest()
	getClient := GetClient()
	if getClient != nil {
		fmt.Println("success")
	}
}
