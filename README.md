# conv
Simple, easy-to-use type convertion library for Go 1.18+. Replacement of the existing [convert repository](https://github.com/vence722/convert).

Requirements
-----
```
Go 1.18+
```

Installation
-----
```
go get github.com/vence722/conv
```

Function List
-----

Type Aliases:
```
Int => int | int8 | int16 | int32 | int64
Uint => uint | uint8 | uint16 | uint32 | uint64
Float => float32 | float64
```

| From Type | To Type | Function  | Function With Default value |
| --------- | ------- | --------- | --------------------------- |
| Int       | string  | Int2Str   | -                           |
| Uint      | string  | Uint2Str  | -                           |
| Float     | string  | Float2Str | -                           | 
| Bool      | string  | Bool2Str  | -                           |
| string    | Int     | Str2Int   | Str2IntOrElse               |
| string    | Uint    | Str2Uint  | Str2UintOrElse              |
| string    | Float   | Str2Float | Str2FloatOrElse             |
| string    | Bool    | Str2Bool  | Str2BoolOrElse              |

Examples
-----

Please refer to [conv_test.go](conv_test.go)