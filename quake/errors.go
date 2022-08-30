/**------------------------------------------------------------**
 * @filename quake/errors.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/17 15:49
 * @desc     go-quake - errors def
 **------------------------------------------------------------**/
package quake

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

var (
	ErrInvalidQuery = errors.New("query is invalid")
	ErrBodyRead     = errors.New("could not read error response")
)

func GetErrorFromResponse(r *http.Response) error {
	errorResponse := new(struct {
		Error string `json:"error"`
	})
	message, err := io.ReadAll(r.Body)
	if err == nil {
		if err := json.Unmarshal(message, errorResponse); err == nil {
			return errors.New(errorResponse.Error)
		}

		return errors.New(strings.TrimSpace(string(message)))
	}

	return ErrBodyRead
}
