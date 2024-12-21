package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"

	"backend/internal/app"
	"backend/internal/mytype"
	"backend/internal/ports/httpgin"
	"backend/internal/repository"

	pgxLogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	log "github.com/sirupsen/logrus"
)

type StudentData struct {
	Id_num_student      int64  `json:"id_num_student"`
	Name_group          string `json:"name_group"`
	Email_student       string `json:"email_student"`
	Second_name_student string `json:"second_name_student"`
	First_name_student  string `json:"first_name_student"`
	Surname_student     string `json:"surname_student"`
}
type studentResponse struct {
	Data StudentData `json:"data"`
}
type StudentRequest struct {
	Id_num_student      uint64 `json:"id_num_student"`
	Name_group          string `json:"name_group"`
	Email_student       string `json:"email_student"`
	Second_name_student string `json:"second_name_student"`
	First_name_student  string `json:"first_name_student"`
	Surname_student     string `json:"surname_student"`
}
type StudentDeleteRequest struct {
	Ids_num_student []string `json:"ids"`
}
type GroupDeleteRequest struct {
	Group_names []string `json:"ids"`
}

type groupResponse struct {
	Name_group              string          `json:"name_group"`
	Studies_direction_group string          `json:"studies_direction_group"`
	Studies_profile_group   string          `json:"studies_profile_group"`
	Start_date_group        mytype.JsonDate `json:"start_date_group"`
	Studies_period_group    uint8           `json:"studies_period_group"`
}

type GroupRequest struct {
	Name_group              string          `json:"name_group"`
	Studies_direction_group string          `json:"studies_direction_group"`
	Studies_profile_group   string          `json:"studies_profile_group"`
	Start_date_group        mytype.JsonDate `json:"start_date_group"`
	Studies_period_group    uint8           `json:"studies_period_group"`
}
type markResponce struct {
	Id_mark          int64  `json:"id_mark"`
	Id_num_student   int64  `json:"id_num_student"`
	Name_semester    string `json:"name_semester"`
	Lesson_name_mark string `json:"lesson_name_mark"`
	Score_mark       int8   `json:"score_mark"`
	Type_mark        string `json:"type_mark"`
}

type MarkRequest struct {
	Id_mark          int64  `json:"id_mark"`
	Id_num_student   int64  `json:"id_num_student"`
	Name_semester    string `json:"name_semester"`
	Lesson_name_mark string `json:"lesson_name_mark"`
	Score_mark       int8   `json:"score_mark"`
	Type_mark        string `json:"type_mark"`
}

type MarksDeleteRequest struct {
	Ids_mark []string `json:"ids"`
}

type testClient struct {
	client  *http.Client
	baseURL string
}

func getTestClient() *testClient {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "0000"
	dbName := "University_DB_test"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.WithError(err).Fatalf("can't parse pgxpool config")
	}
	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxLogrus.NewLogger(logger),
		LogLevel: tracelog.LogLevelDebug,
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.WithError(err).Fatalf("can't create new pool")
	}
	repo := repository.NewRepository(pool, logger)
	usecase := app.NewApp(repo)
	server := httpgin.NewHTTPServer(":18080", usecase)
	testServer := httptest.NewServer(server.Handler())

	return &testClient{
		client:  testServer.Client(),
		baseURL: testServer.URL,
	}
}

var (
	ErrBadRequest = fmt.Errorf("bad request")
	ErrForbidden  = fmt.Errorf("forbidden")
)

func (tc *testClient) getResponse(req *http.Request, out any) error {
	resp, err := tc.client.Do(req)
	if err != nil {
		return fmt.Errorf("unexpected error: %w", err)
	}

	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) {
		if resp.StatusCode == http.StatusBadRequest {
			return ErrBadRequest
		}
		if resp.StatusCode == http.StatusForbidden {
			return ErrForbidden
		}
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}
func (tc *testClient) getStudent(id_num_student int) (studentResponse, error) {
	req, err := http.NewRequest(http.MethodGet, tc.baseURL+"/api/v1/student/"+strconv.Itoa(id_num_student), &bytes.Reader{})
	if err != nil {
		return studentResponse{}, fmt.Errorf("unable to create request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	var response studentResponse
	err = tc.getResponse(req, &response)
	if err != nil {
		return studentResponse{}, err
	}
	return response, nil
}
func (tc *testClient) createStudent(id_num_student uint64, name_group, email_student,
	second_name_student, first_name_student, surname_student string) (studentResponse, error) {
	body := StudentRequest{
		Id_num_student:      id_num_student,
		Name_group:          name_group,
		Email_student:       email_student,
		Second_name_student: second_name_student,
		First_name_student:  first_name_student,
		Surname_student:     surname_student,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return studentResponse{}, fmt.Errorf("unable to marshal: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, tc.baseURL+"/api/v1/student/", bytes.NewReader(data))
	if err != nil {
		return studentResponse{}, fmt.Errorf("unable to create request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	var response studentResponse
	err = tc.getResponse(req, &response)
	if err != nil {
		return studentResponse{}, err
	}
	return response, nil
}
