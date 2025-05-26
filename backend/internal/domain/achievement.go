package domain

type Achievement struct {
	Id_achivment    int64  `json:"id_achivment"`
	Id_num_student  int64  `json:"id_num_student"`
	Id_category     int64  `json:"id_category"`
	Name_achivement string `json:"name_achivement"`
	Date_achivment  string `json:"date_achivment"`
	// Student information
	Student_surname     string `json:"surname_student"`
	Student_name        string `json:"first_name_student"`
	Student_second_name string `json:"second_name_student"`
	Name_group          string `json:"name_group"`
	// Category information
	Category_type string `json:"achivments_type_category"`
}
