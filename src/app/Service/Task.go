package Service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
	"web-screenshot-service/lib"
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

var TaskQueueKey string = "screen_task_queue"

// 任务结果map
var taskRespMap map[string]TaskResp

// 任务队列最大支持10000个任务等待
var waitQueue chan Task

// 消费者任务队列
var consumer chan int
var submitQueue chan int

var waitQueueSize int

var initStatus = false

//InitTask checkout 获取消费者数量， waitQueueSize 等待
func InitTask(checkout int, waitQueueSize int) {
	taskRespMap = make(map[string]TaskResp)
	waitQueue = make(chan Task, waitQueueSize)
	consumer = make(chan int, checkout)
	submitQueue = make(chan int, 2)
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
	prefix := time.Now().Format("20060102")
	uuid := uuid.Must(uuid.NewV4(), nil)
	return fmt.Sprintf("%s:%s", prefix, uuid)
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
	taskResp.Step = "submitQueue"
	taskResp.Finish = false
	//　判断 是否满足队列推送条件
	if isNotAppend() {
		taskResp.Status = "failed"
		taskResp.Finish = true
		log.Errorf("create screen task failed: %v", query)
		return taskResp, errors.New("任务执行等待队列待执行任务数即将达到上限")
	}
	taskResp.Status = "wait"
	// 将任务加入到任务队列
	appendTaskToQueue(taskResp)
	return taskResp, nil
}

func appendTaskToQueue(taskResp TaskResp) {
	content := encode(taskResp)
	log.Infof("append task to submitQueue (%s)", content)
	syncTaskStatus(taskResp)
	lib.RPush(TaskQueueKey, content)
}

// 同步任务状态
func syncTaskStatus(taskResp TaskResp) int64 {
	content := encode(taskResp)
	log.Infof("sync task status info(%s)", content)
	return lib.HashSet(parseDataId(taskResp.TaskId), taskResp.TaskId, content)
}

func parseDataId(taskId string) string {
	split := strings.Split(taskId, ":")
	return fmt.Sprintf("%s:SCREEN_TASK_INFO", split[0])
}

func isNotAppend() bool {
	return false
}

// 处理任务
func ConsumeQueue() {
	for {
		consumer <- 1
		go consumeQueue()
	}
}

func consumeQueue() {
	submitQueue <- 1
	content := lib.LPop(TaskQueueKey)
	if content == "" {
		log.Info("get task is empty , continue")
		time.Sleep(time.Second)
		<-consumer
		<-submitQueue
		return
	}
	<-submitQueue

	log.Infof("start handler task(%s)", content)
	resp, err := decode(content)
	if err == nil {
		handlerTask(resp)
	}

}

func handlerTask(resp TaskResp) {
	resp.Status = "process"
	resp.Step = "runing"
	task := resp.Task
	log.Infof("start handler task(%v)", task)

	syncTaskStatus(resp)
	screen, err := doScreenshotPlus(task.Query, task.Id)

	<-consumer
	resp.Finish = true
	if err != nil {
		log.Warnf("task %s runing failed:%s", task.Id, err.Error())
		resp.Status = "failed"
		resp.Message = err.Error()
		syncTaskStatus(resp)
		return
	}
	resp.ScreenshotRes = screen
	resp.Status = "finish"
	resp.Step = "complete"
	syncTaskStatus(resp)
	log.Infof("task(%v) handler success", task)
}

func GetTaskResp(taskId string) (TaskResp, bool) {
	content := lib.HashGet(parseDataId(taskId), taskId)
	resp, err := decode(content)
	return resp, err == nil
}

func encode(resp TaskResp) string {
	marshal, err := json.Marshal(resp)
	if err != nil {
		log.Error("encode task resp to json error:%s", err.Error())
		panic(err)
	}
	content := fmt.Sprintf("%s", marshal)
	return content
}

func decode(content string) (TaskResp, error) {
	var taskResp TaskResp
	err := json.Unmarshal([]byte(content), &taskResp)
	if err != nil {
		log.Errorf("encode task resp to json error:%s,json=(%s)", err.Error(), content)
		return taskResp, err
	}
	return taskResp, nil
}
