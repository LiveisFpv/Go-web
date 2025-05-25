import type { Pagination } from "./meta";

export interface BudgetReq {
  id_budget: number;
  size_budget: number;
  type_scholarship_budget: string;
  name_semester: string;
}

export interface BudgetResp {
  id_budget: number;
  size_budget: number;
  type_scholarship_budget: string;
  name_semester: string;
}

export interface BudgetsResp{
  data: BudgetResp[];
  pagination: Pagination;
  error: Error;
}

export interface BudgetDeleteReq {
  ids: number[];
}
