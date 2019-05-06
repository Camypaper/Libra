package libra

import (
	"path/filepath"
	"strings"
)

/*
Job !
*/
type Job interface {
	Initializer() Task
	Subtasks() []Task
	Name() string
}

type abstractJob struct {
	src Src
}

func (job abstractJob) Name() string {
	base := filepath.Base(job.src.Name)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func (job abstractJob) Initializer() Task {
	return job.src.compiler()
}
