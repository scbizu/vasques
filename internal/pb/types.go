package pb

import "fmt"

type Keyword string
type MessageType string

var (
	TypeMessage MessageType = "message"
	TypeEnum    MessageType = "enum"
)

func (mt MessageType) print() string {
	return fmt.Sprintf("%s ", mt)
}

var (
	KeywordRepeated   Keyword = "repeated"
	KeywordTypeString Keyword = "string"
	KeywordTypeInt64  Keyword = "int64"
	KeywordTypeInt32  Keyword = "int32"
	KeywordTypeDouble Keyword = "double"
	KeywordTypeFloat  Keyword = "float"
	KeywordTypeBool   Keyword = "bool"
	KeywordTypeBytes  Keyword = "bytes"
)

func (k Keyword) print() string {
	return fmt.Sprintf("%s ", k)
}

func (k Keyword) join(k2 Keyword) Keyword {
	if k.isJoinAble() {
		return Keyword(fmt.Sprintf("%s%s", k.print(), k2.print()))
	}
	return k
}

func (k Keyword) isJoinAble() bool {
	return k == KeywordRepeated
}

func NewType(tname string) Keyword {
	return Keyword(tname)
}
