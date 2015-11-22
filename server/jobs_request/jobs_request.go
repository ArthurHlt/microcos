package jobs_request
import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/bamzi/jobrunner"
	"net/http"
	"fmt"
)
type ReminderEmails struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email
	fmt.Printf("Every 5 sec send reminder emails \n")
}
func CreateJobsRoutes(r martini.Router) {
	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Schedule("@every 5s", ReminderEmails{})
	jobrunner.Schedule("@every 10s", ReminderEmails{})
	entries := jobrunner.Entries()
	fmt.Println(entries[len(entries) - 1].ID)
	r.Get("/status", requestJobs)
}
func requestJobs(r render.Render, req *http.Request) {

	if req.URL.Query().Get("json") != "" {
		r.JSON(200, jobrunner.StatusJson())
	}else {
		r.HTML(200, "jobs/Status", jobrunner.StatusPage())
	}
}