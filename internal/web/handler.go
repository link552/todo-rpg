package web

import (
	"sort"
	"time"
	"strconv"
	"strings"
	"net/http"
	"html/template"
	"todorpg/internal/task"
	"todorpg/internal/storage"
)

type option struct {
	Index int
	Value string
	Name template.HTML
	Text template.HTML
}

type createTaskForm struct {
	Title string
	Short string
	Long string
	Energy string
	Errors []string
	ShortOptions []option
	LongOptions []option
	EnergyOptions []option
}

type currentTaskForm struct {
	Id string
	Title string
	Short string
	Long string
	Energy string
	Errors []string
	ShortOptions []option
	LongOptions []option
	EnergyOptions []option
}

type currentTaskRow struct {
	Priority string
	Title string
	Short template.HTML
	Long template.HTML
	Energy template.HTML
	Form currentTaskForm
}

type completedTaskRow struct {
	CompletedOn string
	Title string
	Short template.HTML
	Long template.HTML
	Energy template.HTML
}

type indexPage struct {
	CreateTaskForm createTaskForm
	CurrentTasksTable []currentTaskRow
	CompletedTasksTable []completedTaskRow
}

func getShortOptions() []option {
	return []option{
		{Index: 0, Value: "", Name: "", Text: ""},
		{Index: 5, Value: "5", Name: "(5) Immediate", Text: "(5) Immediate - Must be done right now."},
		{Index: 4, Value: "4", Name: "(4) Urgent", Text: "(4) Urgent - Needs to be done ASAP."},
		{Index: 3, Value: "3", Name: "(3) High", Text: "(3) High - Should be done today."},
		{Index: 2, Value: "2", Name: "(2) Medium", Text: "(2) Medium - Nice to have done today."},
		{Index: 1, Value: "1", Name: "(1) Low", Text: "(1) Low - Can be done whenever."},
	}
}

func getShortByValue(value string) template.HTML {
	shortOptions := getShortOptions()

	for i := range shortOptions {
		if shortOptions[i].Value == value {
			return shortOptions[i].Name
		}
	}

	return ""
}

func getLongOptions() []option {
	return []option{
		{Index: 0, Value: "", Name: "", Text: ""},
		{Index: 0, Value: "5", Name: "(5) Life Goal", Text: "(5) Life Goal - Must achieve in this life time."},
		{Index: 0, Value: "4", Name: "(4) Significant", Text: "(4) Significant - Life changing achievement."},
		{Index: 0, Value: "3", Name: "(3) High", Text: "(3) High - Really important to achieve."},
		{Index: 0, Value: "2", Name: "(2) Medium", Text: "(2) Medium - Nice to achieve."},
		{Index: 0, Value: "1", Name: "(1) Low", Text: "(1) Low - Don't need to achieve."},
	}
}

func getLongByValue(value string) template.HTML {
	longOptions := getLongOptions()

	for i := range longOptions {
		if longOptions[i].Value == value {
			return longOptions[i].Name
		}
	}

	return ""
}

func getEnergyOptions() []option {
	return []option{
		{Index: 0, Value: "", Name: "", Text: ""},
		{Index: 0, Value: "3", Name: "&star;&star;&star; Extreme", Text: template.HTML("&star;&star;&star; Extreme - Requires tons of energy.")},
		{Index: 0, Value: "2", Name: "&star;&star; Moderate", Text: template.HTML("&star;&star; Moderate - Requires decent energy.")},
		{Index: 0, Value: "1", Name: "&star; Mild", Text: template.HTML("&star; Mild - Requires very little energy.")},
	}
}

func getEnergyByValue(value string) template.HTML {
	energyOptions := getEnergyOptions()

	for i := range energyOptions {
		if energyOptions[i].Value == value {
			return energyOptions[i].Name
		}
	}

	return ""
}

func GetIndex(writer http.ResponseWriter, request *http.Request) {
	page := indexPage{
		CreateTaskForm: createTaskForm{
			ShortOptions: getShortOptions(),
			LongOptions: getLongOptions(),
			EnergyOptions: getEnergyOptions(),
		},
	}

	currentTasks := storage.SelectCurrentTasks()

	// Sort current tasks by priority.
	// Priority = short-term priority value + long-term priority value.
	sort.Slice(currentTasks, func(i, j int) bool {
		return currentTasks[i].Short + currentTasks[i].Long > currentTasks[j].Short + currentTasks[j].Long
	})

	for i := range currentTasks {
		// Assign numerical values to the priorities.
		currentTasks[i].Priority = i + 1

		// Create row and add to the table.
		t := currentTasks[i]

		short := strconv.Itoa(t.Short)
		long := strconv.Itoa(t.Long)
		energy := strconv.Itoa(t.Energy)

		row := currentTaskRow{
			Priority: strconv.Itoa(t.Priority),
			Title: t.Title,
			Short: getShortByValue(short),
			Long: getLongByValue(long),
			Energy: getEnergyByValue(energy),
			Form: currentTaskForm{
				Id: strconv.Itoa(t.Id),
				Title: t.Title,
				Short: short,
				Long: long,
				Energy: energy,
				ShortOptions: getShortOptions(),
				LongOptions: getLongOptions(),
				EnergyOptions: getEnergyOptions(),
			},
		}

		page.CurrentTasksTable = append(page.CurrentTasksTable, row)
	}

	completedTasks := storage.SelectCompletedTasks()

	// Sort completed tasks by completed date.
	sort.Slice(completedTasks, func(i, j int) bool {
		timeA := completedTasks[i].CompletedOn
		timeB := completedTasks[j].CompletedOn
		return timeA.After(timeB)
	})

	for i := range completedTasks {
		// Create row and add to the table.
		t := completedTasks[i]

		row := completedTaskRow{
			CompletedOn: t.CompletedOn.Format(time.DateTime),
			Title: t.Title,
			Short: getShortByValue(strconv.Itoa(t.Short)),
			Long: getLongByValue(strconv.Itoa(t.Long)),
			Energy: getEnergyByValue(strconv.Itoa(t.Energy)),
		}

		page.CompletedTasksTable = append(page.CompletedTasksTable, row)
	}

	tmpl, err := template.ParseFiles("web/page/index.html", "web/partial/create-task-form.html", "web/partial/edit-task-form.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(writer, "index-page", page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostTask(writer http.ResponseWriter, request *http.Request) {
	form := createTaskForm{
		ShortOptions: getShortOptions(),
		LongOptions: getLongOptions(),
		EnergyOptions: getEnergyOptions(),
	}

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Failed to parse form.", http.StatusBadRequest)
		return
	}

	form.Title = request.FormValue("title")
	title := strings.TrimSpace(form.Title)
	if title == "" {
		form.Errors = append(form.Errors, "Title cannot be empty.")
	}

	form.Short = request.FormValue("short")
	short, err := strconv.Atoi(form.Short)
	if err != nil || short < 1 || short > 5 {
		form.Errors = append(form.Errors, "Invalid short-term priority.")
	}

	form.Long = request.FormValue("long")
	long, err := strconv.Atoi(form.Long)
	if err != nil || long < 1 || long > 5 {
		form.Errors = append(form.Errors, "Invalid long-term priority.")
	}

	form.Energy = request.FormValue("energy")
	energy, err := strconv.Atoi(form.Energy)
	if err != nil || energy < 1 || energy > 3 {
		form.Errors = append(form.Errors, "Invalid energy requirement.")
	}

	if (len(form.Errors) > 0) {
		tmpl, err := template.ParseFiles("web/partial/create-task-form.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(writer, "create-task-form", form)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	t := task.Task{
		Title: title,
		Short: short,
		Long: long,
		Energy: energy,
	}

	storage.InsertTask(t)

	writer.Header().Set("HX-Redirect", "/")
}

func PutTask(writer http.ResponseWriter, request *http.Request) {
	form := currentTaskForm{
		ShortOptions: getShortOptions(),
		LongOptions: getLongOptions(),
		EnergyOptions: getEnergyOptions(),
	}

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Failed to parse form.", http.StatusBadRequest)
		return
	}

	form.Id = request.FormValue("id")
	id, err := strconv.Atoi(form.Id)
	if err != nil {
		form.Errors = append(form.Errors, "Invalid id.")
	}

	form.Title = request.FormValue("title")
	title := strings.TrimSpace(form.Title)
	if title == "" {
		form.Errors = append(form.Errors, "Title cannot be empty.")
	}

	form.Short = request.FormValue("short")
	short, err := strconv.Atoi(form.Short)
	if err != nil || short < 1 || short > 5 {
		form.Errors = append(form.Errors, "Invalid short-term priority.")
	}

	form.Long = request.FormValue("long")
	long, err := strconv.Atoi(form.Long)
	if err != nil || long < 1 || long > 5 {
		form.Errors = append(form.Errors, "Invalid long-term priority.")
	}

	form.Energy = request.FormValue("energy")
	energy, err := strconv.Atoi(form.Energy)
	if err != nil || energy < 1 || energy > 3 {
		form.Errors = append(form.Errors, "Invalid energy requirement.")
	}

	if (len(form.Errors) > 0) {
		tmpl, err := template.ParseFiles("web/partial/edit-task-form.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(writer, "edit-task-form", form)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	t := task.Task{
		Id: id,
		Title: title,
		Short: short,
		Long: long,
		Energy: energy,
	}

	storage.UpdateTask(t)

	writer.Header().Set("HX-Redirect", "/")
}

func PostCompleteTask(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Failed to parse form.", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, "Invalid id.", http.StatusBadRequest)
		return
	}

	storage.CompleteTask(id)

	writer.Header().Set("HX-Redirect", "/")
}

func DeleteTask(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, "Failed to parse form.", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, "Invalid id.", http.StatusBadRequest)
		return
	}

	storage.DeleteTask(id)

	writer.Header().Set("HX-Redirect", "/")
}
