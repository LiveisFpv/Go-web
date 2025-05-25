package domain

type Scholarship struct {
	Id_scholarship  int64   `json:"id_scholarship"`
	Id_num_student  int64   `json:"id_num_student"`
	Name_semester   string  `json:"name_semester"`
	Size_scholarshp float64 `json:"size_scholarshp"`
	Id_budget       int64   `json:"id_budget"`
	// Additional fields for joined data
	Student_surname         string `json:"surname_student"`
	Student_name            string `json:"first_name_student"`
	Student_second_name     string `json:"second_name_student"`
	Student_group           string `json:"name_group"`
	Type_scholarship_budget string `json:"type_scholarship_budget"`
}
