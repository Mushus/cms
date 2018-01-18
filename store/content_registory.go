package store

import (
	"reflect"
)

type (
	ContentRegistry struct {
		registry map[string]registry
	}

	registry struct {
		data interface{}
		typ  reflect.Type
	}
)

func NewContentRegistry() *ContentRegistry {
	return &ContentRegistry{}
}

func (c *ContentRegistry) Register(name string, data interface{}) {
	typ := reflect.TypeOf(data)
	c.registry[name] = registry{
		data: data,
		typ:  typ,
	}
}

func (c *ContentRegistry) GetType(data interface{}) string {
	typ := reflect.TypeOf(data)
	for name, r := range c.registry {
		if r.typ == typ {
			return name
		}
	}
	return ""
}
