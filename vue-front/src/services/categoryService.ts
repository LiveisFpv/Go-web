import type { CategoriesResp, CategoryResp } from "@/types/category";
import { api } from "./apiService";


export const categoryService = {
  async getCategories(
    page: number = 1,
    limit: number = 10,
    sort?: string,
    order?: 'asc'|'desc',
    filters?: Record<string, string>
  ): Promise<CategoriesResp> {
    try{
      const response = await api.get(`/api/v1/achievement-category/`, {
        params: {
          page,
          limit,
          sort,
          order,
          ...filters
        }
      });
      return response.data;
    } catch (error){
      throw error;
    }
  },

  async createCategory(category: CategoryResp): Promise<CategoryResp> {
    try {
      const response = await api.post(`/api/v1/achievement-category/`, category);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async updateCategory(category: CategoryResp): Promise<CategoryResp> {
    try {
      const response = await api.put(`/api/v1/achievement-category/`, category);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async deleteCategories(ids: number[]): Promise<void> {
    try {
      await api.delete(`/api/v1/achievement-category/ids`, {
        data: { ids }
      });
    } catch (error) {
      throw error;
    }
  }
}
