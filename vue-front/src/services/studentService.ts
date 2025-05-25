import type { StudentReq, StudentResp, StudentDeleteReq, StudentsResp } from "@/types/student";
import { api } from "@/services/apiService";

export const studentService = {
  async getStudents(
    page: number = 1,
    limit: number = 10,
    sort?: string,
    order?: 'asc'|'desc',
    filters?: Record<string, string>
  ): Promise<StudentsResp> {
    try {
      const response = await api.get(`/api/v1/student/`, {
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

  async createStudent(student: StudentReq): Promise<StudentResp> {
    try {
      const response = await api.post(`/api/v1/student/`, student);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async updateStudent(student: StudentReq): Promise<StudentResp> {
    try {
      const response = await api.put(`/api/v1/student/`, student);
      return response.data;
    } catch (error) {
      throw error;
    }
  },

  async deleteStudents(ids: string[]): Promise<void> {
    try {
      await api.delete(`/api/v1/student/ids`, {
        data: { ids }
      });
    } catch (error) {
      throw error;
    }
  }
};
