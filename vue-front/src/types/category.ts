import type { Pagination } from "@/types/meta";

export interface CategoryReq {
  id_category: number;
  achivments_type_category: string;
  score_category: number;
};

export interface CategoryResp {
  id_category: number;
  achivments_type_category: string;
  score_category: number;
};

export interface CategoriesResp{
  data: CategoryResp[];
  pagination: Pagination;
  error: Error;
}

export interface CategoryDeleteReq{
  ids: number[];
}
