package api

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	nullStr      = "null"
	boolTrueStr  = "true"
	boolFalseStr = "false"
)

func createResponse(data any) (string, error) {
	if data == nil {
		return nullStr, nil
	}

	if res, ok := data.(string); ok {
		return fmt.Sprintf("\"%s\"", res), nil
	}

	if res, ok := data.(bool); ok {
		if res {
			return boolTrueStr, nil
		} else {
			return boolFalseStr, nil
		}
	}

	if res, ok := data.(int); ok {
		return formatInt(res), nil
	}

	if res, ok := data.(int8); ok {
		return formatInt(res), nil
	}

	if res, ok := data.(int16); ok {
		return formatInt(res), nil
	}

	if res, ok := data.(int32); ok {
		return formatInt(res), nil
	}

	if res, ok := data.(int64); ok {
		return formatInt(res), nil
	}

	if res, ok := data.(uint); ok {
		return formatUint(res), nil
	}

	if res, ok := data.(uint8); ok {
		return formatUint(res), nil
	}

	if res, ok := data.(uint16); ok {
		return formatUint(res), nil
	}

	if res, ok := data.(uint32); ok {
		return formatUint(res), nil
	}

	if res, ok := data.(uint64); ok {
		return formatUint(res), nil
	}

	if res, ok := data.(uintptr); ok {
		return formatUint(res), nil
	}

	if res, ok := data.(float32); ok {
		return formatFloat(res), nil
	}

	if res, ok := data.(float64); ok {
		return formatFloat(res), nil
	}

	return "", fmt.Errorf("unhandled data type in response")
}

func formatInt[I interface {
	int | int8 | int16 | int32 | int64
}](i I) string {
	return strconv.FormatInt(int64(i), 10)
}

func formatUint[I interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}](i I) string {
	return strconv.FormatUint(uint64(i), 10)
}

func formatFloat[I interface{ float32 | float64 }](i I) string {
	return strconv.FormatFloat(float64(i), 'f', 10, 64)
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
