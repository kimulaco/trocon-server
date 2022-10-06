package httputil

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadBody[T any](res *http.Response) (T, error) {
	var body T
	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return body, err
	}

	err = json.Unmarshal(bytes, &body)

	return body, err
}
