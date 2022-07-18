package task

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"google.golang.org/api/tasks/v1"
)

type Task struct {
	Name   string
	Due    time.Duration
	Client *tasks.Service
	task   *tasks.Task
}

func (t *Task) Done() (bool, error) {
	if t.task == nil {
		return false, errors.New("task not found")
	}

	if strings.Contains(t.task.Status, "completed") {
		return true, nil
	}

	return false, nil
}

func (t *Task) Late() (bool, error) {
	if t.task == nil {
		return false, errors.New("task not found")
	}

	dd, _ := time.Parse(time.RFC3339, t.task.Due)
	dd = dd.Add(t.Due).Local()
	dd = dd.Add(time.Duration(-1*dd.Hour()) * time.Hour)
	if dd.Add(t.Due).Before(time.Now()) {
		return true, nil
	}

	return false, nil
}

func (t *Task) Find() error {
	tl, err := t.Client.Tasklists.List().Do()
	if err != nil {
		return errors.New("no tasklists found")
	}
	id := tl.Items[0].Id
	tasks, err := t.Client.Tasks.List(id).ShowCompleted(true).ShowHidden(true).ShowDeleted(true).Do()
	if err != nil {
		return err
	}

	for _, task := range tasks.Items {
		fmt.Println(task.Title)
		if strings.Contains(task.Title, t.Name) {
			t.task = task
			return nil
		}
	}

	return errors.New("task not found")
}
