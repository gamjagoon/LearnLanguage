package list

import (
	"fmt"
	"go_todo/lib/todo"
	"strings"
)

type TodoList struct {
	List []todo.Myjob
}

func (list *TodoList) addTodo(myjob *todo.Myjob) {
	list.List = append(list.List, *myjob)
}

func (list *TodoList) CreateJob(name string) {
	// fmt.Printf("Create Job %s \n", name)
	list.addTodo(todo.NewMyjob(name))
}

func (list *TodoList) UpdateJob(num int, set bool) {
	if len(list.List) > num {
		// fmt.Printf("Update Job %d ", num)
		if set {
			list.List[num].StateSet()
		} else {
			list.List[num].StateReset()
		}
	}
}

// Delete Job in the TodoList
// num : list index
// .  number < len(list.List)
func (list *TodoList) DeleteJob(num int) {
	if len(list.List) > num {
		//fmt.Printf("Delete Job %d", num)
		list.List = append(list.List[:num], list.List[num+1:]...)
	}
}

func (list *TodoList) TodoListPrint() {
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("*%11sMyTodo%11s*\n", " ", " ")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for i, v := range list.List {
		fmt.Printf(" %2d | ", i)
		v.Print()
	}
}
