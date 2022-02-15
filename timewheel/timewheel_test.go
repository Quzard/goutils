package timewheel

import (
	"fmt"
	"testing"
	"time"
)

func TimeWheelDefaultJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a timewheel job with key: %v", time.Now().Format(time.RFC3339), key))
}

func TaskJob(key interface{}) {
	fmt.Println(fmt.Sprintf("%v This is a task job with key: %v", time.Now().Format(time.RFC3339), key))
}

func TestName(t *testing.T) {
	//初始化一个时间间隔是1s，一共有60个齿轮的时间轮盘，默认轮盘转动一圈的时间是60s
	tw := CreateTimeWheel(1*time.Second, 5, TimeWheelDefaultJob)

	// 启动时间轮盘
	tw.Start()

	if tw.IsRunning() {
		// 添加一个task
		// 每隔10s执行一次
		// task名字叫task1
		// task的创建时间是time.Now()
		// task执行的任务设置为nil，所以默认执行timewheel的Job，也就是example.TimeWheelDefaultJob
		fmt.Println(fmt.Sprintf("%v Add task task1-5s", time.Now().Format(time.RFC3339)))
		err := tw.AddTask(5*time.Second, "task1-5s", time.Now(), -1, nil)
		if err != nil {
			panic(err)
		}
		// 该Task执行example.TaskJob
		fmt.Println(fmt.Sprintf("%v Add task task2-5s", time.Now().Format(time.RFC3339)))
		err = tw.AddTask(5*time.Second, "task2-5s", time.Now(), -1, TaskJob)
		if err != nil {
			panic(err)
		}

	} else {
		panic("TimeWheel is not running")
	}
	time.Sleep(5 * time.Second)

	// 删除task
	fmt.Println("Remove task task1-5s")
	err := tw.RemoveTask("task1-5s")
	if err != nil {
		panic(err)
	}

	// 删除task
	fmt.Println("Remove task task2-5s")
	err = tw.RemoveTask("task2-5s")
	if err != nil {
		panic(err)
	}
	time.Sleep(6 * time.Second)

	// 关闭时间轮盘
	tw.Stop()
}
