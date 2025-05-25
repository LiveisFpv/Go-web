import type { SemesterReq, SemesterResp, SemesterDeleteReq, SemestersResp } from "@/types/semester";
import { api } from "@/services/apiService";

export const semesterService = {
  async getSemesters(
    page: number = 1,
    limit: number = 10,
    sort?: string,
    order?: 'asc'|'desc',
    filters?: Record<string, string>
  ): Promise<SemestersResp> {
    try {
      const response = await api.get(`/api/v1/semester/`, {
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

  async createSemester(semester: SemesterReq): Promise<SemesterResp> {
    try {
      const response = await api.post(`/api/v1/semester/`, semester);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async updateSemester(semester: SemesterReq): Promise<SemesterResp> {
    try {
      const response = await api.put(`/api/v1/semester/`, semester);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async deleteSemesters(names: string[]): Promise<void> {
    try {
      await api.delete(`/api/v1/semester/ids`, {
        data: { names }
      });
    } catch (error) {
      throw error;
    }
  }
};
