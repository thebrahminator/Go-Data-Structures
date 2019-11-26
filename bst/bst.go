package main

import (
	"errors"
	"fmt"
	"log"
)

type Node struct{
	Value string
	Data string
	Left *Node
	Right *Node
}

func (n *Node) Insert(value, data string) error {
	if  n == nil {
		return errors.New("Cannot insert a value into a nil tree.")
	}

	switch {
		case value == n.value:
			return nil
		case value <=
	}
}
