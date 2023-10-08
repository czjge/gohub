package tasks

import "github.com/czjge/gohub/pkg/crontab"

type Task struct {
}

// 让编译器检查 *Task 是否实现了 接口 crontab.TaskInterface
// 如果没有实现，编译时候会报错
// var name type = value
var _ crontab.TaskInterface = (*Task)(nil)

func New() crontab.TaskInterface {
	return &Task{}
}

func (*Task) Tasks() crontab.Tasks {
	return []crontab.Interface{}
}
