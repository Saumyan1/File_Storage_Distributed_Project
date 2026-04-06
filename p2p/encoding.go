package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface{
	Decode(io.Reader,any) error
}

type GOBDecoder struct{}
//We are reading encoded (binary) data from conn and decoding it into v
func (dec GOBDecoder) Decode(r io.Reader, v any) error{
	return gob.NewDecoder(r).Decode(v)
} 