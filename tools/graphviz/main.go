package main

import (
	"os/user"
	"time"
)

type A struct {
	B
	C map[string]D
}

type B struct {
	E, F  string
	G     user.User
	Timer H
}

type D struct {
	I uint64
}

type H struct {
	Timer time.Timer
	J     chan D
}
