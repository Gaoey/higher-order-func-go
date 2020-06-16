package main

import (
	"errors"
	"fmt"
	"reflect"
)

var count = 0

func Test1(a string) (string, error) {
	return a, nil
}

func Test2(a int) (int, error) {
	return a, nil
}

func Test3(a int) (string, error) {
	return "0", errors.New("Fuckkk")
}


func Count() {
	count += 1
	fmt.Println("count >> ", count)
}

func Monad(callback interface{}, args ...interface{}) (interface{}, error) {
	v := reflect.ValueOf(callback)
	if v.Kind() == reflect.Func {
		vargs := make([]reflect.Value, len(args))
		for i, arg := range args {
			vargs[i] = reflect.ValueOf(arg)
		}
		
		vrets := v.Call(vargs)

		// side effect
		Count()

		switch {
		case len(vrets) == 2 && vrets[1].IsNil():
			return vrets[0].Interface(), nil
		case len(vrets) == 2 && !vrets[1].IsNil():
			return nil, vrets[1].Interface().(error)
		default:
			return vrets[0].Interface(), nil
		}
	}

	return nil, errors.New("Not function")
}

func main() {
	a, err := Monad(Test1, "xxx")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	
	b, err := Monad(Test2, 2)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	c, err := Monad(Test3, 3)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	fmt.Printf("a = %v, b = %v , c = %v\n", a, b, c)
}

