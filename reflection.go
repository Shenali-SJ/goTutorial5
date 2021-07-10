package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId int
	customerId int
}

type employee struct {
	name string
	id int
	address string
	salary int
	country string
}

func check(q interface{}) {
	//t := reflect.TypeOf(q)   //concrete type
	//k := t.Kind()  //specific type
	//
	//fmt.Println("Type ", t)
	//fmt.Println("Kind ", k)

	fmt.Println()

	//if struct
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)   //underlying value
		fmt.Println("Number of fields ", v.NumField())  //number of fields in struct
		for i := 0; i < v.NumField(); i++ {
			//Field() - value of the field
			fmt.Printf("Field: %d, type: %T, value:%v \n", i, v.Field(i), v.Field(i))
		}
	}

	//if number
	if reflect.TypeOf(q).Kind() == reflect.Int {
		t1 := reflect.ValueOf(q).Int()
		fmt.Printf("type: %T, value: %v \n", t1, t1)
	}

	if reflect.TypeOf(q).Kind() == reflect.String {
		t2 := reflect.ValueOf(q).String()
		fmt.Printf("type: %T, value: %v \n", t2, t2)
	}
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {   //making sure that it is a struct
		concreteType := reflect.TypeOf(q).Name()   //get the concrete type of the struct
		query := fmt.Sprintf("insert into %s values(", concreteType)
		value := reflect.ValueOf(q)  //value of the struct -- {21 40}

		for i := 0; i < value.NumField(); i++ {
			switch value.Field(i).Kind() {  //get the specific type of each field
				//int
			case reflect.Int:
				if i == 0 {
					//if it is the first value
					//insert value of the field as a int
					query = fmt.Sprintf("%s%d", query, value.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, value.Field(i).Int())
				}
				//string
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, value.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, value.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("Unsupported type")
}

func main() {
	order1 := order{
		orderId:    1,
		customerId: 3129,
	}
	check(order1)
	check(12)
	check("hello")

	fmt.Println()

	order2 := order{
		orderId:    2,
		customerId: 520,
	}

	employee1 := employee{
		name:    "mindy",
		id:      231,
		address: "Sun Avenue, Texas",
		salary:  2900,
		country: "USA",
	}

	createQuery(employee1)
	createQuery(order2)

}
