package simpleCache

import "fmt"

type Cashier interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
	Delete(key string)
}

type Cashe struct {
	storage map[string]interface{}
	Cashier
}

func NewCache() *Cashe {
	return &Cashe{storage: make(map[string]interface{})}
}

func (c *Cashe) Set(key string, value interface{}) {
	c.storage[key] = value
}

func (c *Cashe) Get(key string) (interface{}, error) {
	if val, ok := c.storage[key]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("no data found")
}

func (c *Cashe) Delete(key string) {
	delete(c.storage, key)
}
