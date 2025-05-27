import type { UserReq, UserResp } from "@/types/user";
import { api } from "./apiService";

export const userService = {
  async getUserbyEmail(email: string): Promise<UserResp> {
    try{
      const response = await api.get(`/api/v1/user/email/${email}`);
      return response.data.data;
    } catch (error) {
      throw error;
    }
  },

  async getUserbyId(id: number): Promise<UserResp> {
    try {
      const response = await api.get(`/api/v1/user/${id}`);
      return response.data.data;
    } catch (error) {
      throw error;
    }
  },

  async updateUser(user: UserReq): Promise<UserResp> {
    try {
      const response = await api.put(`/api/v1/user/`, user);
      return response.data.data;
    } catch (error) {
      throw error;
    }
  },

  async deleteUser(user: UserReq): Promise<void> {
    try {
      const response = await api.delete(`/api/v1/user/`, {
        data: user
      });
    } catch (error) {
      throw error;
    }
  }
}
