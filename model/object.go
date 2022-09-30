package model

type Object struct {
	Name     string
	Capacity int
}

func (o *Object) GetName() string {
	return o.Name
}

func (o *Object) GetCapacity() int {
	return o.Capacity
}

func (o *Object) SetName(name string) {
	o.Name = name
}

func (o *Object) SetCapacity(capacity int) {
	o.Capacity = capacity
}
