package main

import (
	"go_todo/list"
)

func main() {
	joblist := new(list.TodoList)
	joblist.CreateJob("Go To StarBucks")
	joblist.CreateJob("Make Money")
	joblist.CreateJob("Sleep")
	joblist.TodoListPrint()
	joblist.DeleteJob(0)
	joblist.TodoListPrint()
}
