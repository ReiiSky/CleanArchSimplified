package interfaces

import "github.com/Satssuki/CleanArchSimplified/pkg/transform"

// Serializer ..
type Serializer struct {
	data *[]byte
}

// CreateSerializer ..
func CreateSerializer(d *[]byte) *Serializer {
	return &Serializer{
		data: d,
	}
}

// Transform ..
func (s *Serializer) Transform(v interface{}) error {
	return transform.JSONUnmarshal(*s.data, v)
}

type interfaceToDomainAdapter interface {
	Compose(*Serializer) error
}

// Serializable ..
type Serializable interface {
	Serialize(*Serializer) error
}

// ComposeToConcreteObject ..
func ComposeToConcreteObject(v interfaceToDomainAdapter, d *[]byte) error {
	serializer := CreateSerializer(d)
	return v.Compose(serializer)
}
