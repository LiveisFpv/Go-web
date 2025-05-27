import type { Pagination } from "./meta";

export interface MarkReq {
  id_mark: number;
  id_num_student: number;
  name_semester: string;
  lesson_name_mark: string;
  score_mark: number;
  type_mark?: string;
  type_exam: string;
}
export interface MarkResp {
  id_mark: number;
  id_num_student: number;
  name_semester: string;
  lesson_name_mark: string;
  score_mark: number;
  type_mark: string;
  type_exam: string;
  name_group: string;
  second_name_student: string
  first_name_student: string
  surname_student: string
}

export interface MarksResp {
  data: MarkResp[];
  pagination: Pagination;
  error: Error;
}

export interface MarkDeleteReq {
  ids: number[];
}
