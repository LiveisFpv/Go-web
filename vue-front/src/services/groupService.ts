import type { GroupReq,GroupResp,GroupsResp,GroupDeleteReq } from "@/types/group";
import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:15432';

const api = axios.create({
  baseURL: API_URL,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  }
});

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});


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
