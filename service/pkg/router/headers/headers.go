package headers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/jgfranco17/home-network-api/core/pkg/errors"
)

type OriginInfo struct {
	Origin  string `json:"origin"`
	Version string `json:"version"`
}

// Gets the origin info based on the received headers
func CreateOriginInfoHeader(c *gin.Context) (OriginInfo, error) {
	header := c.Request.Header["X-Origin-Info"]

	json_header := OriginInfo{}

	if len(header) == 0 {
		return json_header, errors.NewInputError(c, "X-Origin-Info header not found.")
	}

	error := json.Unmarshal([]byte(header[0]), &json_header)
	if error != nil {
		return json_header, errors.NewInputError(c, "Header schema validation: %s", error.Error())
	}
	return json_header, nil
}
