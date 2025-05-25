import type { Pagination } from "./meta";

export interface SemesterReq{
  name_semester: string;
  date_start_semester: string;
  date_end_semester: string;
}

export interface SemesterResp{
  name_semester: string;
  date_start_semester: string;
  date_end_semester: string;
}

export interface SemestersResp{
  data: SemesterResp[];
  pagination: Pagination;
  error: Error;
}

export interface SemesterDeleteReq{
  names: string[];
}
