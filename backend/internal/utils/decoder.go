package utils

import (
	"encoding/json"
	"io"
)

func GatherBodyDetails(body io.ReadCloser, v any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	if jsonError := decoder.Decode(v); jsonError != nil {
		return jsonError
	}

	return nil
}
