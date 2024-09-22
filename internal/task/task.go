package task

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     string
	Priority    string
	Completed   bool
}

var tasks = make(map[int]Task)
var nextID int = 0

func NewTask(title, description, dueDate, priority string) Task {
	nextID++
	newTask := Task{
		ID:          nextID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Priority:    priority,
		Completed:   false,
	}
	tasks[newTask.ID] = newTask
	return newTask
}

func GetTasks() []Task {
	allTasks := make([]Task, 0)
	for _, task := range tasks {
		allTasks = append(allTasks, task)
	}
	return allTasks
}

func MarkTaskComplete(id int) {
	if task, ok := tasks[id]; ok {
		task.Completed = true
	}
}

func DeleteTask(id int) {
	if _, ok := tasks[id]; ok {
		delete(tasks, id)
	}
}

func EditTask(id int, title, description, dueDate, priority string, completed bool) {
	if task, ok := tasks[id]; ok {
		task.Title = title
		task.Description = description
		task.DueDate = dueDate
		task.Priority = priority
		task.Completed = completed
		tasks[id] = task
	}
}
