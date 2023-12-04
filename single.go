package gomod

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error){

	if(name == "") {
		return name, errors.New("empty name")
	}
	
	message := fmt.Sprintf("my name is %s" , name)
	return message, nil
}