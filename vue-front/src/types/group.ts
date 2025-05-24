import type { Pagination } from "./meta";

export interface GroupReq{
  name_group: string
  studies_direction_group: string
  studies_profile_group: string
  start_date_group: Date
  studies_period_group: number
}

export interface GroupResp{
  name_group: string
  studies_direction_group: string
  studies_profile_group: string
  start_date_group: Date
  studies_period_group: number
}

export interface GroupsResp{
  data: GroupResp[]
  pagination: Pagination
  error: Error
}

export interface GroupDeleteReq{
  ids: string[]
}
