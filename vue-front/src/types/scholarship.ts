import type { Pagination } from "./meta";

export interface ScholarshipReq {
  id_scholarship: number;
  id_num_student: number;
  name_semester: string;
  size_scholarshp: number;
  id_budget: number;
}

export interface ScholarshipResp {
  id_scholarship: number;
  id_num_student: number;
  name_semester: string;
  size_scholarshp: number;
  id_budget: number;
  surname_student: string;
  first_name_student: string;
  second_name_student: string;
  name_group: string;
  type_scholarship_budget: string;
}

export interface AssignScholarshipReq{
  current_semester: string;
  budget_type: string;
}

export interface ScholarshipsResp {
  data: ScholarshipResp[];
  pagination: Pagination;
  error: Error;
}

export interface ScholarshipDeleteReq {
  ids: number[];
}
