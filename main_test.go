package main

import (
    "testing"
    "reflect"
)

func TestReverse(t *testing.T){
    defer func() {
        if err := recover();err!=nil{
            t.Error(err)
        }
    }()

    type DTO1 struct {
        BigNumber int64
        Number int32
        Text string
    }

    value :=  DTO1{BigNumber: 123, Number: 1212333, Text:"Text"}
    v := reflect.ValueOf(value)
    for i := 0; i < v.NumField(); i++ {
        //fmt.Println("\nInput:",v.Field(i))
        //println("Output:",string(reverse(v.Field(i))))
        reverse(v.Field(i))
    }
}
