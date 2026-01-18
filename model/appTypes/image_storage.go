package appTypes

import "encoding/json"

type Storage int

const (
	Local Storage = iota
	Qiniu
)

func (s Storage) MarshallJson() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s *Storage) UnmarshalJson(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = ToStorage(str)
	return nil
}

func ToStorage(str string) Storage {
	switch str {
	case "本地":
		return Local
	case "七牛云":
		return Qiniu
	default:
		return -1
	}
}

func (s Storage) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case Qiniu:
		str = "七牛云"
	default:
		str = "未知存储"
	}
	return str
}
