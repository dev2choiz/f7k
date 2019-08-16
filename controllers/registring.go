package controllers

import (
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/pkg/annotation"
	"go/ast"
)

var registryInstance *ControllerRegistry

type ControllerRegistry struct {
	Imports            []string
	ControllersSpecs   map[string]*ast.TypeSpec
	Controllers        map[string]interfaces.ControllerInterface
	Annotations        map[string]map[string]annotation.IAnnotation
}

func RegistryInstance() *ControllerRegistry {
	if nil == registryInstance {
		registryInstance = &ControllerRegistry{
			*new([]string),
			make(map[string]*ast.TypeSpec),
			make(map[string]interfaces.ControllerInterface),
			make(map[string]map[string]annotation.IAnnotation),
		}
	}

	return registryInstance
}
