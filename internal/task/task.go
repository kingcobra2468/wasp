package task

import (
	"errors"
	"strings"
	"time"

	"google.golang.org/api/tasks/v1"
)

var (
	errTaskNotFound     = errors.New("task not found")
	errTaskListNotFound = errors.New("no tasklists found")
)

// Task contains meta on whether a given task has been
// completed or not yet.
type Task struct {
	// Name of the task as shown on Google Calender.
	Name string
	// Due time of the task relative to midnight. Since Google's API
	// does not return the time of the task, this needs to be passed
	// in manually.
	Due time.Duration
	// Client instance to Google Task service.
	Client *tasks.Service
	// Pointer to the task that matches the Name.
	task *tasks.Task
}

// Done checks whether the given task has been completed.
func (t *Task) Done() (bool, error) {
	if t.task == nil {
		return false, errTaskNotFound
	}

	if strings.Contains(t.task.Status, "completed") {
		return true, nil
	}

	return false, nil
}

// Late checks if the current time is past that of the due date
// of the task.
func (t *Task) Late() (bool, error) {
	if t.task == nil {
		return false, errTaskNotFound
	}

	dd, _ := time.Parse(time.RFC3339, t.task.Due)
	dd = dd.Add(t.Due).Local()
	dd = dd.Add(time.Duration(-1*dd.Hour()) * time.Hour)
	if dd.Add(t.Due).Before(time.Now()) {
		return true, nil
	}

	return false, nil
}

// Find searches for the task (by the title name) from the
// list of returned tasks.
func (t *Task) Find() error {
	tl, err := t.Client.Tasklists.List().Do()
	if err != nil {
		return errTaskListNotFound
	}
	id := tl.Items[0].Id
	tasks, err := t.Client.Tasks.List(id).ShowCompleted(true).ShowHidden(true).ShowDeleted(true).Do()
	if err != nil {
		return err
	}

	for _, task := range tasks.Items {
		if strings.Contains(task.Title, t.Name) {
			t.task = task
			return nil
		}
	}

	return errTaskNotFound
}
