package json

import (
	js "encoding/json"
	"errors"
	"fmt"

	"google.golang.org/grpc/encoding"
)

const Name = "share_json"

type Codes struct{}

func init() {
	encoding.RegisterCodec(Codes{})
}

//c codes

func (Codes) Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return []byte(""), nil
	}
	// if proto.Message
	if vv, ok := v.([]byte); ok {
		return vv, nil
	}

	return js.Marshal(v)
}
func (Codes) Unmarshal(data []byte, v interface{}) error {

	if len(data) == 0 {
		return nil
	}
	if js.Valid(data) {
		err := js.Unmarshal(data, v)
		fmt.Println(string(data))
		fmt.Println(v)
		return err
		return js.Unmarshal(data, v)
	}
	return errors.New("response data is not json")
}
func (Codes) Name() string {
	return Name
}

func (Codes) String() string {
	return Name
}
