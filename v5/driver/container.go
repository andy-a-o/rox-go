package main

import "github.com/andy-a-o/rox-go/v5/server"

type Container struct {
	BoolDefaultFalse server.RoxFlag
	BoolDefaultTrue  server.RoxFlag
}

var container = &Container{
	BoolDefaultFalse: server.NewRoxFlag(false),
	BoolDefaultTrue:  server.NewRoxFlag(true),
}
