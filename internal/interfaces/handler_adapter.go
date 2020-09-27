package interfaces

import "github.com/Satssuki/CleanArchSimplified/pkg/transform"

// Serializer struct exist because Serializable interface
// will bring this type to Serialize struct object
type Serializer struct {
	data *[]byte
}

// CreateSerializer just function to initialize serializer type
// which save your time withour repeating create Serializer
func CreateSerializer(d *[]byte) *Serializer {
	return &Serializer{
		data: d,
	}
}

// Transform bytes which brought by serialize to target object
func (s *Serializer) Transform(v interface{}) error {
	return transform.JSONUnmarshal(*s.data, v)
}

// define action wich domain rule can use Compose method
// to construct their functionality
type interfaceToDomainAdapter interface {
	Compose(*Serializer) error
}

// Serializable this interface link current rule to higher level rule (usecase)
// defined action to serialize usecase model
type Serializable interface {
	Serialize(*Serializer) error
}

// ComposeToConcreteObject this function link rule between interfaces, domain
// handler to call compose domain logic
func ComposeToConcreteObject(v interfaceToDomainAdapter, d *[]byte) error {
	serializer := CreateSerializer(d)
	return v.Compose(serializer)
}
