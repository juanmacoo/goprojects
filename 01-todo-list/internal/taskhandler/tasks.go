package taskhandler

import (
	"fmt"
	"os"
	"path"
	"text/tabwriter"
	"time"

	"tasks/internal/timehelper"
)

type Tasks []task

type task struct {
	description string
	createdAt   time.Time
	isCompleted bool
}

func (t *Tasks) Add(description string) {
	*t = append(*t, task{
		description: description,
		createdAt:   time.Now().Truncate(time.Second),
		isCompleted: false,
	})

	saveTasksToDisk(path.Join(TaskFileLocation, TasksFile), t)
}

func (t *Tasks) List(listAll bool) {
	table := tabwriter.NewWriter(os.Stdout, 3, 5, 3, ' ', 0)
	defer table.Flush()

	if listAll {
		outputFormat := "%v\t%v\t%v"
		fmt.Fprintln(table, fmt.Sprintf(outputFormat, "ID", "Description", "Created At"))
		for index, task := range *t {
			if task.isCompleted {
				verboseCreatedAt := helper.TimeDiffCalculator(task.createdAt)
				fmt.Fprintln(table, fmt.Sprintf(outputFormat, index+1, task.description, verboseCreatedAt))
			}
		}
	} else {
		outputFormat := "%v\t%v\t%v\t%v"
		fmt.Fprintln(table, fmt.Sprintf(outputFormat, "ID", "Description", "Created At", "Done"))
		for index, task := range *t {
			verboseCreatedAt := helper.TimeDiffCalculator(task.createdAt)
			fmt.Fprintln(table, fmt.Sprintf(outputFormat, index+1, task.description, verboseCreatedAt, task.isCompleted))
		}
	}
}

func (t *Tasks) Delete(id int) {
	*t = append((*t)[:id], (*t)[id+1:]...)
	saveTasksToDisk(path.Join(TaskFileLocation, TasksFile), t)
}

func (t *Tasks) Complete(id int) {
	(*t)[id-1].isCompleted = true
	saveTasksToDisk(path.Join(TaskFileLocation, TasksFile), t)
}
