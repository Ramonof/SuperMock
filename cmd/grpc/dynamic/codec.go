package dynamic

import (
	"encoding/json"
)

const CodecName = "proto"

//const CodecName = "proto"

//func init() {
//	encoding.RegisterCodec(&Codec{})
//}

type Codec struct{}

func (c *Codec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *Codec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (c *Codec) Name() string {
	return CodecName
}
