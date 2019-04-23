// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package g

import "github.com/gogf/gf/g/container/gvar"

// Universal variable type, like generics.
//
// 动态变量类型，可以用该类型替代interface{}类型
type Var        = gvar.Var

// Frequently-used map type alias.
//
// 常用map数据结构(使用别名)
type Map        = map[string]interface{}
type MapAnyAny  = map[interface{}]interface{}
type MapAnyStr  = map[interface{}]string
type MapAnyInt  = map[interface{}]int
type MapStrAny  = map[string]interface{}
type MapStrStr  = map[string]string
type MapStrInt  = map[string]int
type MapIntAny  = map[int]interface{}
type MapIntStr  = map[int]string
type MapIntInt  = map[int]int

// Frequently-used slice type alias.
//
// 常用list数据结构(使用别名)
type List       = []Map
type ListAnyStr = []map[interface{}]string
type ListAnyInt = []map[interface{}]int
type ListStrAny = []map[string]interface{}
type ListStrStr = []map[string]string
type ListStrInt = []map[string]int
type ListIntAny = []map[int]interface{}
type ListIntStr = []map[int]string
type ListIntInt = []map[int]int

// Frequently-used slice type alias.
//
// 常用slice数据结构(使用别名)
type Slice      = []interface{}
type SliceAny   = []interface{}
type SliceStr   = []string
type SliceInt   = []int
type Array      = []interface{}
type ArrayAny   = []interface{}
type ArrayStr   = []string
type ArrayInt   = []int
