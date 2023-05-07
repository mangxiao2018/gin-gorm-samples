package util

import (
	"encoding/json"
	"log"
	"reflect"
)

// 通过Json方式进行结构体间数据拷贝
func StructUtils(from, target interface{}) interface{} {
	fromObj, _ := json.Marshal(from)
	_ = json.Unmarshal(fromObj, target)
	log.Println(target)
	return target
}

// 通过反射方式进行结构体间数据拷贝
func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value) //这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
	}
}
