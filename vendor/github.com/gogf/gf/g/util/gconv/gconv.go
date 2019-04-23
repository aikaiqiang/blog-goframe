// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gconv implements powerful and easy-to-use converting functionality for any types of variables.
// 
// 类型转换, 
// 内部使用了bytes作为底层转换类型，效率很高。
package gconv

import (
    "encoding/json"
    "github.com/gogf/gf/g/encoding/gbinary"
    "reflect"
    "strconv"
    "strings"
)

// 转换为string类型的接口
type apiString interface {
    String() string
}

var (
    // 为空的字符串
    emptyStringMap = map[string]struct{}{
        ""      : struct {}{},
        "0"     : struct {}{},
        "off"   : struct {}{},
        "false" : struct {}{},
    }
)

// 将变量i转换为字符串指定的类型t，非必须参数extraParams用以额外的参数传递
func Convert(i interface{}, t string, extraParams...interface{}) interface{} {
    switch t {
        case "int":             return Int(i)
        case "int8":            return Int8(i)
        case "int16":           return Int16(i)
        case "int32":           return Int32(i)
        case "int64":           return Int64(i)
        case "uint":            return Uint(i)
        case "uint8":           return Uint8(i)
        case "uint16":          return Uint16(i)
        case "uint32":          return Uint32(i)
        case "uint64":          return Uint64(i)
        case "float32":         return Float32(i)
        case "float64":         return Float64(i)
        case "bool":            return Bool(i)
        case "string":          return String(i)
        case "[]byte":          return Bytes(i)
        case "[]int":           return Ints(i)
        case "[]string":        return Strings(i)
        case "time.Time":
            if len(extraParams) > 0 {
                return Time(i, String(extraParams[0]))
            }
            return Time(i)
        case "gtime.Time":
            if len(extraParams) > 0 {
                return GTime(i, String(extraParams[0]))
            }
            return *GTime(i)
        case "*gtime.Time":
            if len(extraParams) > 0 {
                return GTime(i, String(extraParams[0]))
            }
            return GTime(i)
        case "time.Duration":   return TimeDuration(i)
        default:
            return i
    }
}

// 转换为二进制[]byte
func Bytes(i interface{}) []byte {
    if i == nil {
        return nil
    }
    switch value := i.(type) {
        case string:  return []byte(value)
        case []byte:  return value
        default:
            return gbinary.Encode(i)
    }
}

// 基础的字符串类型转换
func String(i interface{}) string {
    if i == nil {
        return ""
    }
    switch value := i.(type) {
        case int:     return strconv.FormatInt(int64(value), 10)
        case int8:    return strconv.Itoa(int(value))
        case int16:   return strconv.Itoa(int(value))
        case int32:   return strconv.Itoa(int(value))
        case int64:   return strconv.FormatInt(int64(value), 10)
        case uint:    return strconv.FormatUint(uint64(value), 10)
        case uint8:   return strconv.FormatUint(uint64(value), 10)
        case uint16:  return strconv.FormatUint(uint64(value), 10)
        case uint32:  return strconv.FormatUint(uint64(value), 10)
        case uint64:  return strconv.FormatUint(uint64(value), 10)
        case float32: return strconv.FormatFloat(float64(value), 'f', -1, 32)
        case float64: return strconv.FormatFloat(value, 'f', -1, 64)
        case bool:    return strconv.FormatBool(value)
        case string:  return value
        case []byte:  return string(value)
        default:
            if f, ok := value.(apiString); ok {
                // 如果变量实现了String()接口，那么使用该接口执行转换
                return f.String()
            } else {
                // 默认使用json进行字符串转换
                jsonContent, _ := json.Marshal(value)
                return string(jsonContent)
            }
    }
}

//false: false, "", 0, "false", "off", empty slice/map
func Bool(i interface{}) bool {
    if i == nil {
        return false
    }
    if v, ok := i.(bool); ok {
        return v
    }
    if s, ok := i.(string); ok {
        if _, ok := emptyStringMap[s]; ok {
            return false
        }
        return true
    }
    rv := reflect.ValueOf(i)
    switch rv.Kind() {
        case reflect.Ptr:    return !rv.IsNil()
        case reflect.Map:    fallthrough
        case reflect.Array:  fallthrough
        case reflect.Slice:  return rv.Len() != 0
        case reflect.Struct: return true
        default:
            s := String(i)
            if _, ok := emptyStringMap[s]; ok {
                return false
            }
            return true

    }
}

func Int(i interface{}) int {
    if i == nil {
        return 0
    }
    if v, ok := i.(int); ok {
        return v
    }
    return int(Int64(i))
}

func Int8(i interface{}) int8 {
    if i == nil {
        return 0
    }
    if v, ok := i.(int8); ok {
        return v
    }
    return int8(Int64(i))
}

func Int16(i interface{}) int16 {
    if i == nil {
        return 0
    }
    if v, ok := i.(int16); ok {
        return v
    }
    return int16(Int64(i))
}

func Int32(i interface{}) int32 {
    if i == nil {
        return 0
    }
    if v, ok := i.(int32); ok {
        return v
    }
    return int32(Int64(i))
}

func Int64(i interface{}) int64 {
    if i == nil {
        return 0
    }
    if v, ok := i.(int64); ok {
        return v
    }
    switch value := i.(type) {
        case int:     return int64(value)
        case int8:    return int64(value)
        case int16:   return int64(value)
        case int32:   return int64(value)
        case int64:   return value
        case uint:    return int64(value)
        case uint8:   return int64(value)
        case uint16:  return int64(value)
        case uint32:  return int64(value)
        case uint64:  return int64(value)
        case float32: return int64(value)
        case float64: return int64(value)
        case bool:
            if value {
                return 1
            }
            return 0
        default:
            s := String(value)
            // 按照十六进制解析
            if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
                if v, e := strconv.ParseInt(s[2 : ], 16, 64); e == nil {
                    return v
                }
            }
            // 按照八进制解析
            if len(s) > 1 && s[0] == '0' {
                if v, e := strconv.ParseInt(s[1 : ], 8, 64); e == nil {
                    return v
                }
            }
            // 按照十进制解析
            if v, e := strconv.ParseInt(s, 10, 64); e == nil {
                return v
            }
            // 按照浮点数解析
            return int64(Float64(value))
    }
}

func Uint(i interface{}) uint {
    if i == nil {
        return 0
    }
    if v, ok := i.(uint); ok {
        return v
    }
    return uint(Uint64(i))
}

func Uint8(i interface{}) uint8 {
    if i == nil {
        return 0
    }
    if v, ok := i.(uint8); ok {
        return v
    }
    return uint8(Uint64(i))
}

func Uint16(i interface{}) uint16 {
    if i == nil {
        return 0
    }
    if v, ok := i.(uint16); ok {
        return v
    }
    return uint16(Uint64(i))
}

func Uint32(i interface{}) uint32 {
    if i == nil {
        return 0
    }
    if v, ok := i.(uint32); ok {
        return v
    }
    return uint32(Uint64(i))
}

func Uint64(i interface{}) uint64 {
    if i == nil {
        return 0
    }
    switch value := i.(type) {
        case int:     return uint64(value)
        case int8:    return uint64(value)
        case int16:   return uint64(value)
        case int32:   return uint64(value)
        case int64:   return uint64(value)
        case uint:    return uint64(value)
        case uint8:   return uint64(value)
        case uint16:  return uint64(value)
        case uint32:  return uint64(value)
        case uint64:  return value
        case float32: return uint64(value)
        case float64: return uint64(value)
        case bool:
            if value {
                return 1
            }
            return 0
        default:
            s := String(value)
            // 按照十六进制解析
            if len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
                if v, e := strconv.ParseUint(s[2 : ], 16, 64); e == nil {
                    return v
                }
            }
            // 按照八进制解析
            if len(s) > 1 && s[0] == '0' {
                if v, e := strconv.ParseUint(s[1 : ], 8, 64); e == nil {
                    return v
                }
            }
            // 按照十进制解析
            if v, e := strconv.ParseUint(s, 10, 64); e == nil {
                return v
            }
            // 按照浮点数解析
            return uint64(Float64(value))
    }
}

func Float32 (i interface{}) float32 {
    if i == nil {
        return 0
    }
    if v, ok := i.(float32); ok {
        return v
    }
    v, _ := strconv.ParseFloat(strings.TrimSpace(String(i)), 64)
    return float32(v)
}

func Float64 (i interface{}) float64 {
    if i == nil {
        return 0
    }
    if v, ok := i.(float64); ok {
        return v
    }
    v, _ := strconv.ParseFloat(strings.TrimSpace(String(i)), 64)
    return v
}

