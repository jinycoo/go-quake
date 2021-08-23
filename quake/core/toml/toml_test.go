// +build go1.16

package toml_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/jinycoo/go-quake/quake/core/toml"
	"github.com/jinycoo/go-quake/quake/core/toml/internal/tag"
)

// Test if the error message matches what we want for invalid tests. Every slice
// entry is tested with strings.Contains.
//
// Filepaths are glob'd
var errorTests = map[string][]string{
	"encoding-bad-utf8*":            {"invalid UTF-8 byte"},
	"encoding-utf16*":               {"files cannot contain NULL bytes; probably using UTF-16"},
	"string-multiline-escape-space": {`invalid escape: '\ '`},
}

// Test metadata; all keys listed as "keyname: type".
var metaTests = map[string]string{
	// TODO: this probably should have albums as a Hash as well?
	"table-array-implicit": `
			albums.songs: ArrayHash
			albums.songs.name: String
		`,
}

type parser struct{}

func (p parser) Encode(input string) (output string, outputIsError bool, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			switch rr := r.(type) {
			case error:
				retErr = rr
			default:
				retErr = fmt.Errorf("%s", rr)
			}
		}
	}()

	var tmp interface{}
	err := json.Unmarshal([]byte(input), &tmp)
	if err != nil {
		return "", false, err
	}

	buf := new(bytes.Buffer)
	err = toml.NewEncoder(buf).Encode(tag.Remove(tmp))
	if err != nil {
		return err.Error(), true, retErr
	}

	return buf.String(), false, retErr
}

func (p parser) Decode(input string) (output string, outputIsError bool, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			switch rr := r.(type) {
			case error:
				retErr = rr
			default:
				retErr = fmt.Errorf("%s", rr)
			}
		}
	}()

	var d interface{}
	if _, err := toml.Decode(input, &d); err != nil {
		return err.Error(), true, retErr
	}

	j, err := json.MarshalIndent(tag.Add("", d), "", "  ")
	if err != nil {
		return "", false, err
	}
	return string(j), false, retErr
}
