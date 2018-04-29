package docker

import (
	"fmt"
	"testing"
)

func TestCreateContainer(t *testing.T) {
	containerId, err := CreateContainer("test", []string{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(containerId)
}
