
export interface UserReq{
  user_id: number
  user_login: string
  user_email: string
  user_student_id: number | null
  user_role: string | null
  user_password: string
}

export interface UserResp{
  user_id: number
  user_login: string
  user_email: string
  user_student_id: number | null
  user_role: string | null
}
