package handlers

import (
	models "GO-WEB/internal/Models"
	"GO-WEB/internal/config"
	"GO-WEB/internal/forms"
	"GO-WEB/internal/helpers"
	renders "GO-WEB/internal/render"
	"encoding/json"
	"log"
	"net/http"
)

// Repository Type
type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// Creates New Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repo to handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	err := renders.RenderTemp(w, "home.page.tmpl", &models.TemplateData{}, r)
	if err != nil {
		log.Println(err)
	}
}

// About is the About page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	err := renders.RenderTemp(w, "about.page.tmpl", &models.TemplateData{}, r)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) Kratos(w http.ResponseWriter, r *http.Request) {
	err := renders.RenderTemp(w, "kratos.page.tmpl", &models.TemplateData{}, r)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) Batman(w http.ResponseWriter, r *http.Request) {
	err := renders.RenderTemp(w, "batman.page.tmpl", &models.TemplateData{}, r)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	err := renders.RenderTemp(w, "searchAvailability.page.tmpl", &models.TemplateData{}, r)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	startDate := r.Form.Get("start_date")
	endDate := r.Form.Get("end_date")
	m.App.Session.Put(r.Context(), "startDate", startDate)
	m.App.Session.Put(r.Context(), "endDate", endDate)
	// Redirect after a success post using http redirect
	http.Redirect(w, r, "/makeReservation", http.StatusSeeOther)
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// handles request and sends JSON Response
func (m *Repository) RoomAvailabilityJson(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{OK: true, Message: "Available"}
	out, err := json.MarshalIndent(response, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	err := renders.RenderTemp(w, "contact.page.tmpl", &models.TemplateData{}, r)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	// Should have same name of the value of the data in PostMakeReservation
	data["reservation"] = emptyReservation
	// passing the form so that we can re-render the form based on the field validation and don't lose anything entered
	err := renders.RenderTemp(w, "makeReservation.page.tmpl",
		&models.TemplateData{
			Form: forms.NewForm(nil),
			Data: data,
		}, r)
	if err != nil {
		log.Println(err)
	}
}

func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone_number"),
	}
	form := forms.NewForm(r.PostForm)
	form.RequiredFields("first_name", "last_name", "email", "phone_number")
	form.MinLength("first_name", 2)
	form.MinLength("last_name", 2)
	form.ValidEmail("email")
	// form.Has("first_name", r)
	// if Form has error
	if !form.IsValid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		err := renders.RenderTemp(w, "makeReservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		}, r)
		if err != nil {
			log.Println(err)
		}
		return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	// Redirect after a success post using http redirect
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// GET Reservation Summary
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("cannot get reservation from session")
		m.App.Session.Put(r.Context(), "error", "Can't Get Reservation Details for this Session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data := make(map[string]interface{})
	data["summary"] = reservation

	startDate, ok := m.App.Session.Get(r.Context(), "startDate").(string)
	if !ok {
		m.App.ErrorLog.Println("cannot get startDate from session")
		return
	}
	endDate, ok := m.App.Session.Get(r.Context(), "endDate").(string)
	if !ok {
		m.App.ErrorLog.Println("cannot get endDate from session")
		return
	}
	//Remove data from session after getting the needed data
	m.App.Session.Remove(r.Context(), "reservation")
	m.App.Session.Remove(r.Context(), "startDate")
	m.App.Session.Remove(r.Context(), "endDate")
	// passing the form so that we can re-render the form based on the field validation and don't lose anything entered
	err := renders.RenderTemp(w, "reservationSummary.page.tmpl",
		&models.TemplateData{
			Data:  data,
			Start: startDate,
			End:   endDate,
		}, r)
	if err != nil {
		log.Println(err)
	}
}
