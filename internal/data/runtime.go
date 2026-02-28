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

	unquoted, err := strconv.Unquote(jsonString)
	if err == nil {
		value := strings.TrimSuffix(unquoted, " mins")
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return ErrParsingRuntime
		}
		*r = Runtime(valueInt)
		return nil
	}

	valueInt, err := strconv.Atoi(jsonString)
	if err != nil {
		return ErrInvalidRuntime
	}

	*r = Runtime(valueInt)
	return nil
}
