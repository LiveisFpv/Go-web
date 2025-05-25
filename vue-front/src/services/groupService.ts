import type { GroupReq,GroupResp,GroupsResp,GroupDeleteReq } from "@/types/group";
import { api } from "@/services/apiService";

export const groupService = {
  async getGroups(
    page: number = 1,
    limit: number = 10,
    sort?: string,
    order?: 'asc'|'desc',
    filters?: Record<string, string>
  ): Promise<GroupsResp>{
    try{
      const response = await api.get(`/api/v1/group/`,{
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

  async createGroup(group: GroupReq): Promise<GroupResp>{
    try {
      const response = await api.post(`/api/v1/group/`, group);
      return response.data;
    } catch (error){
      throw error;
    }
  },

  async updateGroup(group: GroupReq): Promise<GroupResp>{
    try {
      const response = await api.put(`/api/v1/group/`, group);
      return response.data;
    } catch (error){
      throw error;
    }
  },

  async deleteGroups(ids: string[]): Promise<void>{
    try {
      await api.delete(`/api/v1/group/ids`, {
        data: {ids}
      });
    } catch (error){
      throw error;
    }
  }
};
