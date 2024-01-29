package main

import "fmt"

type PolicyError struct {
	File    string
	Line    int
	Column  int
	Message string
}

func (e *PolicyError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", e.File, e.Line, e.Column, e.Message)
}
