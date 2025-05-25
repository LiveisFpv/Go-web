import type { BudgetReq, BudgetResp, BudgetsResp } from "@/types/budget";
import { api } from "./apiService";

export const budgetService = {
  async getBudgets(
    page: number = 1,
    limit: number = 10,
    sort?: string,
    order?: 'asc'|'desc',
    filters?: Record<string, string>
  ): Promise<BudgetsResp> {
    try {
      const response = await api.get(`/api/v1/budget/`, {
        params: {
          page,
          limit,
          sort,
          order,
          ...filters
        }
      });
      return response.data;
    } catch (error) {
      throw error;
    }
  },
  async createBudget(budget: BudgetReq): Promise<BudgetResp> {
    try {
      const response = await api.post(`/api/v1/budget/`, budget);
      return response.data;
    } catch (error) {
      throw error;
    }
  },
  async updateBudget(budget: BudgetReq): Promise<BudgetResp> {
    try {
      const response = await api.put(`/api/v1/budget/`, budget);
      return response.data;
    } catch (error) {
      throw error;
    }
  },
  async deleteBudgets(ids: number[]): Promise<void> {
    try {
      await api.delete(`/api/v1/budget/ids`, {
        data: { ids }
      });
    } catch (error) {
      throw error;
    }
  }

};
