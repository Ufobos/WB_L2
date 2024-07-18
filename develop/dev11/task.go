package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Event struct {
	Date time.Time `json:"date"`
}

type Result struct {
	Result []Event `json:"result"`
}

func main() {
	eventsHandler := &EventsHandler{}
	routes := CreateRoutes(eventsHandler)
	handler := RequestLog(routes)
	server := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("close server error: %s", err.Error())
	}
}

// Роутер запросов.
func CreateRoutes(eventsHandler *EventsHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/events_for_month", eventsHandler.GetEventByMonth)
	mux.HandleFunc("/events_for_week", eventsHandler.GetEventByWeek)
	mux.HandleFunc("/events_for_day", eventsHandler.GetEventByDay)
	mux.HandleFunc("/create_event", eventsHandler.CreateEvent)
	mux.HandleFunc("/update_event", eventsHandler.UpdateEvent)
	mux.HandleFunc("/delete_event", eventsHandler.DeleteEvent)
	return mux
}

// Промежуточное ПО для логирования запросов.
func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(w, req)
		log.Printf("Request: %s Method: %s\n", req.RequestURI, req.Method)
	})
}

// Интерфейс сервиса бизнес-логики.
type EventsProcessorInterface interface {
	CreateEvent(userID, date string) (Result, error)
	UpdateEvent(userID, date string) (Result, error)
	DeleteEvent(userID, date string) error
	GetEventByDay(userID, date string) (Result, error)
	GetEventByWeek(userID, date string) (Result, error)
	GetEventByMonth(userID, date string) (Result, error)
}

// Контроллер, обработчик запросов.
type EventsHandler struct {
	processor EventsProcessorInterface
}

func NewEventsHandler(processor EventsProcessorInterface) *EventsHandler {
	return &EventsHandler{processor}
}

func (e *EventsHandler) GetEventByMonth(w http.ResponseWriter, r *http.Request) {
	userID, date, err := ValidateGetRequest(r)
	if err != nil {
		switch {
		case errors.As(err, new(BadMethodError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(BadRequestError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
	}
	result, err := e.processor.GetEventByMonth(userID, date)
	if err != nil {
		switch {
		case errors.As(err, new(NotFoundError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(ServiceUnavailableError)):
			WrapErrorWithStatus(w, err, http.StatusServiceUnavailable)
			return
		default:
			WrapErrorWithStatus(w, err, http.StatusInternalServerError)
			return
		}
	}
	WrapOk(w, result)
}

func (e *EventsHandler) GetEventByWeek(w http.ResponseWriter, r *http.Request) {
	userID, date, err := ValidateGetRequest(r)
	if err != nil {
		switch {
		case errors.As(err, new(BadMethodError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(BadRequestError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
	}
	result, err := e.processor.GetEventByWeek(userID, date)
	if err != nil {
		switch {
		case errors.As(err, new(NotFoundError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(ServiceUnavailableError)):
			WrapErrorWithStatus(w, err, http.StatusServiceUnavailable)
			return
		default:
			WrapErrorWithStatus(w, err, http.StatusInternalServerError)
			return
		}
	}
	WrapOk(w, result)
}

func (e *EventsHandler) GetEventByDay(w http.ResponseWriter, r *http.Request) {
	userID, date, err := ValidateGetRequest(r)
	if err != nil {
		switch {
		case errors.As(err, new(BadMethodError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(BadRequestError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
	}
	result, err := e.processor.GetEventByDay(userID, date)
	if err != nil {
		switch {
		case errors.As(err, new(NotFoundError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(ServiceUnavailableError)):
			WrapErrorWithStatus(w, err, http.StatusServiceUnavailable)
			return
		default:
			WrapErrorWithStatus(w, err, http.StatusInternalServerError)
			return
		}
	}
	WrapOk(w, result)
}

func (e *EventsHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	userID, date, err := ValidatePostRequest(r)
	if err != nil {
		switch {
		case errors.As(err, new(BadMethodError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(BadRequestError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
	}
	result, err := e.processor.CreateEvent(userID, date)
	if err != nil {
		switch {
		case errors.As(err, new(NotFoundError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(ServiceUnavailableError)):
			WrapErrorWithStatus(w, err, http.StatusServiceUnavailable)
			return
		default:
			WrapErrorWithStatus(w, err, http.StatusInternalServerError)
			return
		}
	}
	WrapOk(w, result)
}

func (e *EventsHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	userID, date, err := ValidatePostRequest(r)
	if err != nil {
		switch {
		case errors.As(err, new(BadMethodError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(BadRequestError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
	}
	result, err := e.processor.UpdateEvent(userID, date)
	if err != nil {
		switch {
		case errors.As(err, new(NotFoundError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(ServiceUnavailableError)):
			WrapErrorWithStatus(w, err, http.StatusServiceUnavailable)
			return
		default:
			WrapErrorWithStatus(w, err, http.StatusInternalServerError)
			return
		}
	}
	WrapOk(w, result)
}

func (e *EventsHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	userID, date, err := ValidatePostRequest(r)
	if err != nil {
		switch {
		case errors.As(err, new(BadMethodError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(BadRequestError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		}
	}
	err = e.processor.DeleteEvent(userID, date)
	if err != nil {
		switch {
		case errors.As(err, new(NotFoundError)):
			WrapErrorWithStatus(w, err, http.StatusBadRequest)
			return
		case errors.As(err, new(ServiceUnavailableError)):
			WrapErrorWithStatus(w, err, http.StatusServiceUnavailable)
			return
		default:
			WrapErrorWithStatus(w, err, http.StatusInternalServerError)
			return
		}
	}
	WrapOkDelete(w)
}

// Валидация GET запросов и получение параметров запроса.
func ValidateGetRequest(r *http.Request) (userID, date string, err error) {
	if r.Method != http.MethodGet {
		err = BadMethodError{r.Method, http.MethodGet}
		return "", "", err
	}
	q := r.URL.Query()
	_, ok := q["user_id"]
	if !ok {
		err = BadRequestError{"в запросе отсутствует user_id"}
		return "", "", err
	}
	_, ok = q["date"]
	if !ok {
		err = BadRequestError{"в запросе отсутствует date"}
		return "", "", err
	}
	userID = q.Get("user_id")
	date = q.Get("date")
	return userID, date, nil
}

// Валидация POST запросов и получение патаметров запроса.
func ValidatePostRequest(r *http.Request) (userID, date string, err error) {
	if r.Method != http.MethodPost {
		err = BadMethodError{r.Method, http.MethodPost}
		return userID, date, err
	}
	userID = r.FormValue("user_id")
	if len(userID) == 0 {
		err = BadRequestError{"в теле запроса отсутствует user_id"}
		return userID, date, err
	}
	date = r.FormValue("date")
	if len(date) == 0 {
		err = BadRequestError{"в теле запроса отсутствует date"}
		return userID, date, err
	}
	return userID, date, nil
}

// Сериализация результата и отправка ответа.
func WrapOk(w http.ResponseWriter, result Result) {
	res, err := json.Marshal(result)
	if err != nil {
		err = InternalServerError{err.Error()}
		WrapErrorWithStatus(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Println(err)
	}
}

// Отправка ответа об успешном удалении.
func WrapOkDelete(w http.ResponseWriter) {
	var m = map[string]string{
		"result": "event deleted",
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		log.Println(err)
	}
}

// Оборачивает ошибки в json.
func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string{
		"error": err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpStatus)
	_, err = w.Write(res)
	if err != nil {
		log.Println(err)
	}
}

type BadRequestError struct {
	msg string
}

func (b BadRequestError) Error() string {
	return fmt.Sprintf("Bad Request: %s", b.msg)
}

type BadMethodError struct {
	wrongMethod string
	rightMethod string
}

func (b BadMethodError) Error() string {
	return fmt.Sprintf("bad request: bad method %s, method must be %s", b.wrongMethod, b.rightMethod)
}

type NotFoundError struct {
	resource string
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", n.resource)
}

type ServiceUnavailableError struct{}

func (s ServiceUnavailableError) Error() string {
	return "service unavailable"
}

type InternalServerError struct {
	err string
}

func (i InternalServerError) Error() string {
	return fmt.Sprintf("Internal Server Error: %s", i.err)
}
