package conv

import (
	"strconv"
)

func Int2Str[T SignedInterger](input T) string {
	return strconv.FormatInt(int64(input), 10)
}

func Uint2Str[T UnsignedInteger](input T) string {
	return strconv.FormatUint(uint64(input), 10)
}

func Str2Int[T SignedInterger](input string) (T, error) {
	if res, err := strconv.ParseInt(input, 10, 64); err != nil {
		return 0, err
	} else {
		return T(res), nil
	}
}

func Str2IntOrElse[T SignedInterger](input string, defaultValue T) T {
	if res, err := Str2Int[T](input); err != nil {
		return defaultValue
	} else {
		return res
	}
}
