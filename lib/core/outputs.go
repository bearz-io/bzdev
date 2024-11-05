package core

type Outputs struct {
	data map[string]interface{}
}

func NewOutputs() *Outputs {
	return &Outputs{
		data: make(map[string]interface{}),
	}
}

func (o *Outputs) Set(key string, value interface{}) {
	o.data[key] = value
}

func (o *Outputs) Get(key string) interface{} {
	return o.data[key]
}

func (o *Outputs) Has(key string) bool {
	_, ok := o.data[key]
	return ok
}
