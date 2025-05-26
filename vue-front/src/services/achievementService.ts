import type { AchivementReq, AchivementResp, AchivementsResp } from "@/types/achievement";
import { api } from "./apiService";

export const achievementService = {
  async getAchievements(
    page: number = 1,
    limit: number = 10,
    sort?: string,
    order?: 'asc' | 'desc',
    filters?: Record<string, string>
  ): Promise<AchivementsResp> {
    try {
      const response = await api.get(`/api/v1/achievement/`, {
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

  async createAchievement(achievement: AchivementReq): Promise<AchivementResp> {
    try {
      const response = await api.post(`/api/v1/achievement/`, achievement);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async updateAchievement(achievement: AchivementReq): Promise<AchivementResp> {
    try {
      const response = await api.put(`/api/v1/achievement/`, achievement);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async deleteAchievements(ids: string[]): Promise<void> {
    try {
      await api.delete(`/api/v1/achievement/ids`, {
        data: { ids }
      });
    } catch (error) {
      throw error;
    }
  }
};
