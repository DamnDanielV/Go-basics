package main

// estructura de lista de tareas
type taskList struct {
	tasks []*task
}

// AddTask: añade una nueva tarea al slice
// newTask: la tarea a añadir
func (actualTasks *taskList) AddTask(newTask *task) {
	actualTasks.tasks = append(actualTasks.tasks, newTask)
}

// deleteTask: elimina una tarea del slice
// indexTask: indice de la tarea a eliminar
func (actualTasks *taskList) deleteTask(indexTask int) {
	actualTasks.tasks = append(actualTasks.tasks[:indexTask], actualTasks.tasks[indexTask+1:]...)
}

// printTasks: imprime un slice de tareas
func (actualTasks *taskList) printTasks() {
	for _, v := range actualTasks.tasks {
		println("nombre", v.name)
		println("decripvion", v.description)
	}
}

// printTasksCompleted: imprime un slice de tareas
func (actualTasks *taskList) printTasksCompleted() {
	for _, v := range actualTasks.tasks {
		if v.completed {
			println("nombre", v.name)
			println("decripvion", v.description)

		}
	}
}

// estructura de tarea
type task struct {
	name        string
	description string
	completed   bool
}

// markComplete: actualiza una tarea a completada
func (t *task) markComplete() {
	t.completed = true
}

// updateDescription: actualiza la descripcion de una tarea
func (t *task) updateDescription(newDescription string) {
	t.description = newDescription
}

// updateName: actualiza el nombre
func (t *task) updateName(newName string) {
	t.name = newName
}
func main() {
	t1 := &task{
		name:        "Go basics",
		description: "Comprender las base de go",
	}
	t2 := &task{
		name:        "Go RestServer",
		description: "Crear un servidor con Go",
	}
	t3 := &task{
		name:        "Go go go",
		description: "more go go go go",
	}
	lista := &taskList{
		[]*task{
			t1,
			t2,
		},
	}
	t2.markComplete()
	lista.AddTask(t3)
	lista.printTasksCompleted()
}
