package store

import (
	"fmt"
	"io/ioutil"

	"github.com/Satssuki/CleanArchSimplified/pkg/system"
	"github.com/Satssuki/CleanArchSimplified/pkg/transform"
)

type listForDeserializeType struct {
	List interface{} `json:"list"`
}

// Deserialize function that attach json seed to object
// v is reference of specified struct type
// seedName is name of file that contain byte of data
// eg. user = user.json
func Deserialize(v interface{}, seedName string) error {
	bytes, err := ioutil.ReadFile(system.CrossPlatformPath(fmt.Sprintf("./store/%s.json", seedName)))
	if err == nil {
		list := listForDeserializeType{List: v}
		err = transform.JSONUnmarshal(bytes, &list)
	}
	return err
}
