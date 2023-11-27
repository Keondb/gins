package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1 //本地
	QiNiu ImageType = 2 //QiNiu
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "Local"
	case QiNiu:
		str = "QiNiu"
	default:
		str = "其他"
	}
	return str
}
