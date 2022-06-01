package workerpool

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testTask struct {
	Name          string
	TaskProcessor func(...interface{})
}

func (t testTask) Run() {
	t.TaskProcessor(t.Name)
}

func TestDispatcher(t *testing.T) {
	pool := NewGoroutinePool(25)
	taskSize := 50
	taskCounter := 0

	//wait for jobs to finish
	wg := &sync.WaitGroup{}
	wg.Add(taskSize)

	//specific task
	sampleStringTaskFn := func(dm ...interface{}) {
		if myinput, ok := dm[0].(string); ok {
			time.Sleep(time.Second)
			if myinput != "" {
				fmt.Printf("Finished %s\n", myinput)
			}
			taskCounter++
			wg.Done()
		}
	}

	var tasks []testTask
	for v := 0; v < taskSize; v++ {
		tasks = append(tasks, testTask{
			Name:          fmt.Sprintf("task %d", v),
			TaskProcessor: sampleStringTaskFn,
		})
	}

	for _, task := range tasks {
		pool.ScheduleWork(task)
	}
	pool.Close()

	wg.Wait()

	assert.NotNil(t, pool)
	assert.EqualValues(t, taskCounter, taskSize)
}
