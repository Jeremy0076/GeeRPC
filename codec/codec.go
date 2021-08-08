package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"time"
)

const (
	GobType  = "application/gob"
	JsonType = "application/json"

	MagicNumber = 0x3bef5c
)

type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}

var (
	NewCodecFuncMap map[Type]NewCodecFunc
	_               Codec = (*GobCodec)(nil)
	DefaultOption         = &Option{
		MagicNumber:    MagicNumber,
		CodecType:      GobType,
		ConnectTimeout: time.Second * 10,
	}
)

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

// Option | Option{MagicNumber: xxx, CodecType: xxx} | Header{ServiceMethod ...} | Body interface{} |
//| <------      固定 JSON 编码      ------>  | <-------   编码方式由 CodeType 决定   ------->|
//| Option | Header1 | Body1 | Header2 | Body2 | ...
type Option struct {
	MagicNumber    int  // MagicNumber marks this is a geerpc req
	CodecType      Type // client may choose diff Codec to encode body
	ConnectTimeout time.Duration
	HandleTimeout  time.Duration
}
