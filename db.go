package main

import "encoding/json"

type db []map[string]string

func new_db(text []byte) (db, error) {
	v := db{}

	err := json.Unmarshal(text, &v)
	return v, err
}

func (data *db) hosts(host_field string) []string {
	result := make([]string, 0)
	for i := range *data {
		h, ok := (*data)[i][host_field]
		if ok {
			result = append(result, h)
		}
	}

	return result
}