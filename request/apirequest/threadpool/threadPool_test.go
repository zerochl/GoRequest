package threadpool

import (
	"testing"
	"log"
	"time"
)

func TestThreadPool_AddTask(t *testing.T) {
	threadPool := NewThreadPool(20, 3)
	//for i := 0; i < 10;i++ {
	//	log.Println("before add task i:", strconv.Itoa(i))
	//	threadPool.AddTask(func() error {
	//		log.Println("in task i:", strconv.Itoa(i))
	//		return nil
	//	})
	//}
	quit := make(chan int)
	threadPool.AddTask(func() error {
		time.Sleep(time.Second)
		log.Println("in task 1:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 2)
		log.Println("in task 2:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 5)
		log.Println("in task 3:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 8)
		log.Println("in task 4:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 1)
		log.Println("in task 5:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 2)
		log.Println("in task 6:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 1)
		log.Println("in task 7:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 12)
		log.Println("in task 8:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 3)
		log.Println("in task 9:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 2)
		log.Println("in task 10:")
		return nil
	})
	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 4)
		log.Println("in task 11:")
		return nil
	})

	threadPool.AddTask(func() error {
		time.Sleep(time.Second * 3)
		log.Println("in task 12:")
		return nil
	})
	log.Println("任务全部添加")
	<- quit
}
