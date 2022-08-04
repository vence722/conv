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

func Str2Uint[T UnsignedInteger](input string) (T, error) {
	if res, err := strconv.ParseUint(input, 10, 64); err != nil {
		return 0, err
	} else {
		return T(res), nil
	}
}

func Str2UintOrElse[T UnsignedInteger](input string, defaultValue T) T {
	if res, err := Str2Uint[T](input); err != nil {
		return defaultValue
	} else {
		return res
	}
}

func Float2Str[T Float](input T, prec ...int) string {
	if len(prec) > 0 {
		return strconv.FormatFloat(float64(input), 'f', prec[0], 64)
	}
	return strconv.FormatFloat(float64(input), 'g', 6, 64)
}

func Str2Float[T Float](input string) (T, error) {
	if res, err := strconv.ParseFloat(input, 64); err != nil {
		return 0, err
	} else {
		return T(res), nil
	}
}

func Str2FloatOrElse[T Float](input string, defaultValue T) T {
	if res, err := Str2Float[T](input); err != nil {
		return defaultValue
	} else {
		return res
	}
}
