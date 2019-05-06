package libra

/*
Reporter reports task's status and its output.
*/
type Reporter interface {
	ReportStart(task Job)
	Report(task Task, status Status)
	ReportEnd()
}

type nothingReporter struct{}

func (nothingReporter) ReportStart(task Job) {

}
func (nothingReporter) Report(task Task, status Status) {

}
func (nothingReporter) ReportEnd() {

}
