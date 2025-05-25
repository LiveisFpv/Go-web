import type { MarkReq, MarkResp, MarksResp } from "@/types/mark";
import { api } from "@/services/apiService";

export const markService = {
  async getMarks(
    page: number = 1,
    limit: number = 10,
    sort?: string,
    order?: 'asc'|'desc',
    filters?: Record<string, string>
  ): Promise<MarksResp>{
    try {
      const response = await api.get(`/api/v1/mark/`, {
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

  async createMark(mark: MarkReq): Promise<MarkResp>{
    try {
      const response = await api.post(`/api/v1/mark/`, mark);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async updateMark(mark: MarkReq): Promise<MarkResp>{
    try {
      const response = await api.put(`/api/v1/mark/`, mark);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async deleteMarks(ids: string[]): Promise<void>{
    try{
      await api.delete(`/api/v1/mark/ids`,{
        data:{ids}
      });
    } catch (error){
      throw error;
    }
  }

};
