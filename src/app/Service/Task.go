package Service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"time"
	"web-srceenshot-service/lib/util"
)

type Task struct {
	Id    string   `json:"id"`
	Query CapQuery `json:"query"`
}

type TaskResp struct {
	TaskId        string        `json:"task_id"`
	Task          Task          `json:"task"`
	Finish        bool          `json:"finish"`
	Status        string        `json:"status"`
	Step          string        `json:"step"`
	Message       string        `json:"message"`
	ScreenshotRes ScreenshotRes `json:"screenshot_res"`
}

// 任务结果map
var taskRespMap map[string]TaskResp

// 任务队列最大支持10000个任务等待
var waitQueue chan Task

//
var consumer chan int

var waitQueueSize int

var initStatus = false

//InitTask checkout 获取消费者数量， waitQueueSize 等待
func InitTask(checkout int, waitQueueSize int) {
	taskRespMap = make(map[string]TaskResp)
	waitQueue = make(chan Task, waitQueueSize)
	consumer = make(chan int, checkout)
	initStatus = true
	go ConsumeQueue()
}

func checkInit() {
	if initStatus == false {
		InitTask(20, 10000)
	}
}

// 生成任务id
func genTaskId(url string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return util.GetMd5(url + timestamp)
}

// 创建任务
func CreateTask(query CapQuery) (TaskResp, error) {
	log.Infof("start create screen task:%v", query)
	checkInit()
	// 生成任务id
	taskId := genTaskId(query.Url)
	var task Task
	task.Id = taskId
	task.Query = query
	// 判断是否可以加入队列
	var taskResp TaskResp
	taskResp.Task = task
	taskResp.TaskId = taskId
	taskResp.Step = "queue"
	taskResp.Finish = false
	if len(waitQueue) >= cap(waitQueue)-10 {
		taskResp.Status = "failed"
		taskResp.Finish = true
		log.Errorf("create screen task failed: %v", query)
		return taskResp, errors.New("任务执行等待队列待执行任务数即将达到上限")
	}
	taskResp.Status = "wait"
	// 将任务加入到任务队列
	taskRespMap[taskId] = taskResp
	waitQueue <- task
	return taskResp, nil
}

// 处理任务
func ConsumeQueue() {
	for {
		consumer <- 1
		task := <-waitQueue
		go handlerTask(task)
	}
}

func handlerTask(task Task) {
	log.Infof("start handler task(%v)", task)
	resp := taskRespMap[task.Id]
	resp.Status = "process"
	resp.Step = "runing"
	taskRespMap[task.Id] = resp
	screen, err := doScreenshotPlus(task.Query, task.Id)
	<-consumer
	resp.Finish = true
	if err != nil {
		log.Warnf("task %s runing failed:%s", task.Id, err.Error())
		resp.Status = "failed"
		resp.Message = err.Error()
		taskRespMap[task.Id] = resp
		return
	}
	resp.ScreenshotRes = screen
	resp.Status = "finish"
	resp.Step = "complete"
	taskRespMap[task.Id] = resp
	log.Infof("task(%v) handler success", task)
}

func GetTaskResp(taskId string) (TaskResp, bool) {
	resp, ok := taskRespMap[taskId]
	return resp, ok
}
