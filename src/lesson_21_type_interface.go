package main

import "fmt"

// Simple interface with one method
type Initer interface {
	Init()
}

// Embedding of another interface
type Server interface {
	Initer
	Start()

	SetModule(string, Initer)
}

type Module struct {}

// Initialisation of module
func (s *Module) Init() {
	fmt.Println("Module: Init")
}

type SomeServer struct {
	modules map[string]Initer
}

// Initialisation of server
func (s *SomeServer) Init() {
	s.modules = make(map[string]Initer, 3)
	fmt.Println("SomeServer: Init")
}

// Add module to server
func (s *SomeServer) SetModule(name string, module Initer) {
	fmt.Println("SetModule:", name)
	s.modules[name] = module
}

// Starting server
func (s *SomeServer) Start() {
	for _, module := range s.modules {
		module.Init()
	}
	fmt.Println("Start")
}

// Base function for building and starting server
func ServerRun(server Server, modules ...string) {
	server.Init()

	for _, moduleName := range modules {
		module := new(Module)
		server.SetModule(moduleName, module)
	}

	server.Start()
}

func main() {
	ServerRun(new(SomeServer), "config", "log", "blog")
}
