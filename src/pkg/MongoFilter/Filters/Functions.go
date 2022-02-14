package Filters

import (
	"strconv"
	"time"
)

//TYPE CONVERSION FUNCTIONS

func DateTime(layout string) func(StringValue string) interface{} {
	return func(StringValue string) interface{} {
		dateValue, _ := time.Parse(layout, StringValue)
		return dateValue
	}
}

func Integer(StringValue string) interface{} {
	intValue, _ := strconv.Atoi(StringValue)
	return intValue
}
