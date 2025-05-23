import type { Pagination } from "./meta"

export interface StudentReq{
  id_num_student: number
  name_group: string
  email_student: string
  second_name_student: string
  first_name_student: string
  surname_student: string
}

export interface StudentDeleteReq{
  ids: string[]
}

export interface StudentResp{
  id_num_student: number
  name_group: string
  email_student: string
  second_name_student: string
  first_name_student: string
  surname_student: string
}

export interface StudentsResp{
  data: StudentResp[]
  pagination: Pagination
  error: Error
}
