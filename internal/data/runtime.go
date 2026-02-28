package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var (
	ErrInvalidRuntime = errors.New("invalid runtime")
	ErrParsingRuntime = errors.New("error parsing runtime")
)

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJsonValue := strconv.Quote(jsonValue)
	return []byte(quotedJsonValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonData []byte) error {
	jsonString := string(jsonData)
	jsonUnquoted, err := strconv.Unquote(jsonString)
	if err != nil || jsonUnquoted == "" {
		return ErrInvalidRuntime
	}

	value := strings.TrimSuffix(jsonUnquoted, " mins")
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return ErrParsingRuntime
	}

	*r = Runtime(valueInt)

	return nil
}
