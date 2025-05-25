import type { ScholarshipReq, ScholarshipResp, ScholarshipsResp } from "@/types/scholarship";
import { api } from "@/services/apiService";

export const scholarshipService = {
  async getScholarships(
    page: number = 1,
    limit: number = 10,
    sort?: string,
    order?: 'asc'|'desc',
    filters?: Record<string, string>
  ): Promise<ScholarshipsResp>{
    try {
      const response = await api.get(`/api/v1/scholarship/`, {
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

  async createScholarship(scholarship: ScholarshipReq): Promise<ScholarshipResp>{
    try {
      const response = await api.post(`/api/v1/scholarship/`, scholarship);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async updateScholarship(scholarship: ScholarshipReq): Promise<ScholarshipResp>{
    try {
      const response = await api.put(`/api/v1/scholarship/`, scholarship);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async deleteScholarships(ids: string[]): Promise<void>{
    try{
      await api.delete(`/api/v1/scholarship/ids`,{
        data:{ids}
      });
    } catch (error){
      throw error;
    }
  }
}; 