package taskhandler

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

const TaskFileLocation = "bin/"
const TasksFile = "tasks.csv"
const timeLayout = "2006-01-02 15:04:05"

func InitializeFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating:", err)
			os.Exit(1)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		writer.Write([]string{"id", "description", "createdAt", "isCompleted"})
		writer.Flush()
	}
}

func ReadTasksFromDisk(filename string) (tasksFromDisk *Tasks) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening:", err)
		os.Exit(1)
	}
	lockFile(file)
	defer unlockFile(file)

	reader := csv.NewReader(file)
	csvTasks, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tasksFromDisk = &Tasks{}

	for index, csvTask := range csvTasks {
		if index == 0 {
			continue
		} else {
			timeFromFile, _ := time.Parse(timeLayout, csvTask[2])
			isCompletedFromFile, _ := strconv.ParseBool(csvTask[3])

			*tasksFromDisk = append(*tasksFromDisk, task{
				description: csvTask[1],
				createdAt:   timeFromFile,
				isCompleted: isCompletedFromFile,
			})
		}
	}

	return
}

func saveTasksToDisk(filename string, t *Tasks) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating:", err)
		os.Exit(1)
	}
	lockFile(file)
	defer unlockFile(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"id", "description", "createdAt", "isCompleted"}); err != nil {
		fmt.Println("Error writing:", err)
		os.Exit(1)
	}
	for index, task := range *t {
		if err := writer.Write([]string{strconv.FormatInt((int64(index+1)), 10), task.description, task.createdAt.Format(timeLayout), fmt.Sprint(task.isCompleted)}); err != nil {
			fmt.Println("Error writing:", err)
			os.Exit(1)
		}
	}
}


func lockFile(f *os.File){
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		fmt.Println("Error when locking file")
		os.Exit(1)
	}
}

func unlockFile(f *os.File){
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	f.Close()
}
