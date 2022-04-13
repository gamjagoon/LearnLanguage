package todo

import "fmt"

type Myjob struct {
	Name     *string
	Complete bool
}

func NewMyjob(name string) *Myjob {
	ret := Myjob{
		Name:     &name,
		Complete: false,
	}
	return &ret
}

func (job *Myjob) StateSet() {
	job.Complete = true
}

func (job *Myjob) StateReset() {
	job.Complete = false
}

func (job *Myjob) Print() {
	if job.Complete {
		fmt.Printf("[✅] %s\n", *job.Name)
	} else {
		fmt.Printf("[❌] %s\n", *job.Name)
	}
}
