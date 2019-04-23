// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gp.

// Package gparser provides a flexible and easy way for accessing/converting variable and JSON/XML/YAML/TOML contents.
package gparser

import (
    "github.com/gogf/gf/g/encoding/gjson"
    "time"
)

type Parser struct {
    json *gjson.Json
}

// 将变量转换为Parser对象进行处理，该变量至少应当是一个map或者array，否者转换没有意义
// value可以传递nil, 表示创建一个空的Parser对象
func New (value interface{}, unsafe...bool) *Parser {
    return &Parser{gjson.New(value, unsafe...)}
}

// 非并发安全Parser对象
func NewUnsafe (value...interface{}) *Parser {
    if len(value) > 0 {
        return &Parser{gjson.New(value[0], false)}
    }
    return &Parser{gjson.New(nil, false)}
}

func Load (path string) (*Parser, error) {
    if j, e := gjson.Load(path); e == nil {
        return &Parser{j}, nil
    } else {
        return nil, e
    }
}

// 支持的数据内容格式：json(默认), xml, yaml/yml, toml
func LoadContent (data []byte, dataType...string) (*Parser, error) {
    if j, e := gjson.LoadContent(data, dataType...); e == nil {
        return &Parser{j}, nil
    } else {
        return nil, e
    }
}

// 设置自定义的层级分隔符号
func (p *Parser) SetSplitChar(char byte) {
    p.json.SetSplitChar(char)
}

// 设置是否执行层级冲突检查，当键名中存在层级符号时需要开启该特性，默认为关闭。
// 开启比较耗性能，也不建议允许键名中存在分隔符，最好在应用端避免这种情况。
func (p *Parser) SetViolenceCheck(check bool) {
    p.json.SetViolenceCheck(check)
}

// 将指定的json内容转换为指定结构返回，查找失败或者转换失败，目标对象转换为nil
// 注意第二个参数需要给的是变量地址
func (p *Parser) GetToVar(pattern string, v interface{}) error {
    return p.json.GetToVar(pattern, v)
}

// 获得一个键值对关联数组/哈希表，方便操作，不需要自己做类型转换
// 注意，如果获取的值不存在，或者类型与json类型不匹配，那么将会返回nil
func (p *Parser) GetMap(pattern string) map[string]interface{} {
    return p.json.GetMap(pattern)
}

// 获得一个数组[]interface{}，方便操作，不需要自己做类型转换
// 注意，如果获取的值不存在，或者类型与json类型不匹配，那么将会返回nil
func (p *Parser) GetArray(pattern string) []interface{} {
    return p.json.GetArray(pattern)
}

// 返回指定json中的string
func (p *Parser) GetString(pattern string) string {
    return p.json.GetString(pattern)
}

func (p *Parser) GetStrings(pattern string) []string {
    return p.json.GetStrings(pattern)
}

func (p *Parser) GetInterfaces(pattern string) []interface{} {
    return p.json.GetInterfaces(pattern)
}

func (p *Parser) GetTime(pattern string, format ... string) time.Time {
    return p.json.GetTime(pattern, format...)
}

func (p *Parser) GetTimeDuration(pattern string) time.Duration {
    return p.json.GetTimeDuration(pattern)
}

// 返回指定json中的bool(false:"", 0, false, off)
func (p *Parser) GetBool(pattern string) bool {
    return p.json.GetBool(pattern)
}

func (p *Parser) GetInt(pattern string) int {
    return p.json.GetInt(pattern)
}

func (p *Parser) GetInt8(pattern string) int8 {
    return p.json.GetInt8(pattern)
}

func (p *Parser) GetInt16(pattern string) int16 {
    return p.json.GetInt16(pattern)
}

func (p *Parser) GetInt32(pattern string) int32 {
    return p.json.GetInt32(pattern)
}

func (p *Parser) GetInt64(pattern string) int64 {
    return p.json.GetInt64(pattern)
}

func (p *Parser) GetInts(pattern string) []int {
    return p.json.GetInts(pattern)
}

func (p *Parser) GetUint(pattern string) uint {
    return p.json.GetUint(pattern)
}

func (p *Parser) GetUint8(pattern string) uint8 {
    return p.json.GetUint8(pattern)
}

func (p *Parser) GetUint16(pattern string) uint16 {
    return p.json.GetUint16(pattern)
}

func (p *Parser) GetUint32(pattern string) uint32 {
    return p.json.GetUint32(pattern)
}

func (p *Parser) GetUint64(pattern string) uint64 {
    return p.json.GetUint64(pattern)
}

func (p *Parser) GetFloat32(pattern string) float32 {
    return p.json.GetFloat32(pattern)
}

func (p *Parser) GetFloat64(pattern string) float64 {
    return p.json.GetFloat64(pattern)
}

func (p *Parser) GetFloats(pattern string) []float64 {
    return p.json.GetFloats(pattern)
}

// 将指定变量转换为struct对象(对象属性赋值)
func (p *Parser) GetToStruct(pattern string, objPointer interface{}) error {
    return p.json.GetToStruct(pattern, objPointer)
}

// 根据pattern查找并设置数据
// 注意：写入的时候"."符号只能表示层级，不能使用带"."符号的键名
func (p *Parser) Set(pattern string, value interface{}) error {
    return p.json.Set(pattern, value)
}

// 计算指定pattern的元素长度(pattern对应数据类型为map[string]interface{}/[]interface{}时有效)
func (p *Parser) Len(pattern string) int {
    return p.json.Len(pattern)
}

// 指定pattern追加元素
func (p *Parser) Append(pattern string, value interface{}) error {
    return p.json.Append(pattern, value)
}

// 动态删除变量节点
func (p *Parser) Remove(pattern string) error {
    return p.json.Remove(pattern)
}

// 根据约定字符串方式访问json解析数据，参数形如： "items.name.first", "list.0"; 当pattern为空时，表示获取所有数据
// 返回的结果类型的interface{}，因此需要自己做类型转换;
// 如果找不到对应节点的数据，返回nil;
func (p *Parser) Get(pattern...string) interface{} {
    return p.json.Get(pattern...)
}

// 转换为map[string]interface{}类型,如果转换失败，返回nil
func (p *Parser) ToMap() map[string]interface{} {
    return p.json.ToMap()
}

// 转换为[]interface{}类型,如果转换失败，返回nil
func (p *Parser) ToArray() []interface{} {
    return p.json.ToArray()
}

/* 以下为数据文件格式转换，支持类型：xml, json, yaml/yml, toml */

func (p *Parser) ToXml(rootTag...string) ([]byte, error) {
    return p.json.ToXml(rootTag...)
}

func (p *Parser) ToXmlIndent(rootTag...string) ([]byte, error) {
    return p.json.ToXmlIndent(rootTag...)
}

func (p *Parser) ToJson() ([]byte, error) {
    return p.json.ToJson()
}

func (p *Parser) ToJsonIndent() ([]byte, error) {
    return p.json.ToJsonIndent()
}

func (p *Parser) ToYaml() ([]byte, error) {
    return p.json.ToYaml()
}

func (p *Parser) ToToml() ([]byte, error) {
    return p.json.ToToml()
}

// 打印Json对象
func (p *Parser) Dump() error {
    return p.json.Dump()
}

// 将变量解析为对应的struct对象，注意传递的参数为struct对象指针
func (p *Parser) ToStruct(o interface{}) error {
    return p.json.ToStruct(o)
}

func VarToXml(value interface{}, rootTag...string) ([]byte, error) {
    return New(value).ToXml(rootTag...)
}

func VarToXmlIndent(value interface{}, rootTag...string) ([]byte, error) {
    return New(value).ToXmlIndent(rootTag...)
}

func VarToJson(value interface{}) ([]byte, error) {
    return New(value).ToJson()
}

func VarToJsonIndent(value interface{}) ([]byte, error) {
    return New(value).ToJsonIndent()
}

func VarToYaml(value interface{}) ([]byte, error) {
    return New(value).ToYaml()
}

func VarToToml(value interface{}) ([]byte, error) {
    return New(value).ToToml()
}

// 将变量解析为对应的struct对象，注意传递的参数为struct对象指针
func VarToStruct(value interface{}, obj interface{}) error {
    return New(value).ToStruct(obj)
}

