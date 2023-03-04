// https://golangbyexample.com/singleton-design-pattern-go/
// Singleton is a creational design pattern, which ensures that only one object of its kind exists
// and provides a single point of access to it for any other code.
// Some of the cases where the singleton object is applicable:
// DB instance – we only want to create only one instance of DB object and that instance will be used throughout the application.
// Logger instance – again only one instance of the logger should be created and it should be used throughout the application.
package creational

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creting Single Instance Now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}
