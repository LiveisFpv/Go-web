import type { Pagination } from "./meta";

export interface AchivementReq {
  id_achivment: number;
  id_num_student: number;
  id_category: number;
  name_achivement: string;
  date_achivement: string;
}

export interface AchivementResp {
  id_achivment: number;
  id_num_student: number;
  id_category: number;
  name_achivement: string;
  date_achivement: string;
  second_name_student: string;
  first_name_student: string;
  surname_student: string;
  achivments_type_category: string;
  name_group: string;
}

export interface AchivementsResp {
  data: AchivementResp[];
  pagination: Pagination;
  error: Error;
}

export interface AchivementDeleteReq {
  ids: string[];
}
