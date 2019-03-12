package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Before callFunc1")
	callFunc1()
	fmt.Println("After callFunc1")
}

func callFunc1() {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", RecoverExceptionDetails(FunctionName()), " and recovered in callFunc1 function, Error Info: ", errD)
		}
	}()
	// Exception occurred at  main.callFunc3:48 << main.callFunc2:36 << main.callFunc1:24  and recovered in callFunc1 function, Error Info:  runtime error: index out of range
	fmt.Println("Before callFunc2")
	callFunc2()
	fmt.Println("After callFunc2")
}

func callFunc2() {
	/*defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", RecoverExceptionDetails(FunctionName()), " and recovered in callFunc2 function, Error Info: ", errD)
		}
	}()*/
	// Exception occurred at  main.callFunc3:48 << main.callFunc2:36  and recovered in callFunc2 function, Error Info:  runtime error: index out of range
	fmt.Println("Before callFunc3")
	callFunc3()
	fmt.Println("After callFunc3")
}

func callFunc3() {
	/*defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred at ", RecoverExceptionDetails(FunctionName()), " and recovered in callFunc3 function, Error Info: ", errD)
		}
	}()*/
	// Exception occurred at  main.callFunc3:48  and recovered in callFunc3 function, Error Info:  runtime error: index out of range
	var ai []int
	ai[0] = 0
	panic("Panic in callFunc3")
	fmt.Println("After Panic, inside callFunc3")
}

// FunctionName - should return name of calling function
func FunctionName() string {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in FunctionName(), Error Info: ", errD)
		}
	}()
	pc, _, _, _ := runtime.Caller(1)
	fmt.Println("****runtime.FuncForPC(pc).Name():", runtime.FuncForPC(pc).Name())
	funcName := strings.TrimSuffix(runtime.FuncForPC(pc).Name(), ".func1") // This is for defer function
	funcName = strings.TrimSuffix(funcName, ".1")                          // This is for go runtine function
	fmt.Println("****funcName:", funcName)
	return funcName
}

// RecoverExceptionDetails will take one formal parameter as function name - should return exeception detail formated as: packageName.functionName:lineNumber
// Each format detail appended  with "<<" if it is multiple stack frames(LIFO). For example: packageName.functionName:lineNumber << packageName.functionName:lineNumber
func RecoverExceptionDetails(strfuncName string) string {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in RecoverExceptionDetails(), Error Info: ", errD)
		}
	}()
	var output string
	flag := false
	for skip := 1; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		strfunctionName := runtime.FuncForPC(pc).Name()
		if strings.Contains(file, "/runtime/") && strings.Contains(strfunctionName, "runtime.") {
			flag = true
			continue
		}
		if flag && strings.HasSuffix(file, ".go") {
			output += strfunctionName + ":" + strconv.Itoa(line) + " << "
			if strfuncName == strfunctionName {
				output = strings.TrimSuffix(output, " << ")
				fmt.Println("****output:", output)
				break
			}
		}
	}
	return output
}

/*
Before callFunc1
Before callFunc2
Before callFunc3
****runtime.FuncForPC(pc).Name(): main.callFunc1.func1
****funcName: main.callFunc1
****output: main.callFunc3:48 << main.callFunc2:36 << main.callFunc1:24
Exception occurred at  main.callFunc3:48 << main.callFunc2:36 << main.callFunc1:24  and recovered in callFunc1 function, Error Info:  runtime error: index out of range
After callFunc1
*/
