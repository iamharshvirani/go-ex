package patterns

import (
	"fmt"
	"sync"
)

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("Creating single instance now.")
		instance = &singleton{}
	})
	fmt.Println("Single instance already created.")
	return instance
}

func (s *singleton) DoSomething() {
	fmt.Println("Singleton is doing something.")
}

func RunSingleton() {
	fmt.Println("--- Singleton Pattern ---")
	for i := 0; i < 3; i++ {
		go func() {
			instance := GetInstance()
			instance.DoSomething()
		}()
	}
	fmt.Scanln()
}
