package domain

type Mark struct {
	Id_mark          int64
	Id_num_student   int64
	Name_semester    string
	Lesson_name_mark string
	Score_mark       int8
	Type_mark        string
	// Student information
	Student_surname     string
	Student_name        string
	Student_second_name string
	Student_group       string
}
