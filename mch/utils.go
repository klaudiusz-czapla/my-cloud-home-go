package mch

import "encoding/json"

func FromJson[T any](j string) (*any, error) {

	bytesArr := []byte(j)

	var v any
	err := json.Unmarshal(bytesArr, &v)

	if err != nil {
		return nil, err
	}

	return &v, nil
}

func ToJson[T any](v any) (string, error) {

	m, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(m), nil
}
