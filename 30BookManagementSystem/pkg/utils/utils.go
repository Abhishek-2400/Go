package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	// json.Unmarshal([]byte(body), &x)
	json.Unmarshal(body, x)
}
