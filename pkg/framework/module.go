package framework

import "go.uber.org/dig"

type Module interface {
	Register(container *dig.Container) error
}
