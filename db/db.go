package db

import "encoding/json"

type Entry map[string]string
type Db []Entry

func FromJson(text []byte) (*Db, error) {
	v := Db{}

	err := json.Unmarshal(text, &v)
	return &v, err
}

func (data *Db) Hosts(host_field string) []string {
	result := make([]string, 0)
	for i := range *data {
		h, ok := (*data)[i][host_field]
		if ok {
			result = append(result, h)
		}
	}

	return result
}
