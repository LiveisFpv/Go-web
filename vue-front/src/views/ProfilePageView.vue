<script setup lang="ts">
import router from '@/router';
import { userService } from '@/services/userService';
import { studentService } from '@/services/studentService';
import { useAuthStore } from '@/stores/auth';
import type { UserReq, UserResp } from '@/types/user';
import type { StudentResp, StudentReq } from '@/types/student';
import type { AxiosError } from 'axios';
import { onMounted, ref } from 'vue';

const authStore = useAuthStore()
const user = ref<UserResp>();
const student = ref<StudentResp>();
const isEditing = ref(false);
const editedUser = ref<UserReq>({
  user_id: 0,
  user_login: '',
  user_email: '',
  user_student_id: null,
  user_role: null,
  user_password: ''
});

const editedStudent = ref<StudentReq>({
  id_num_student: 0,
  name_group: '',
  email_student: '',
  second_name_student: '',
  first_name_student: '',
  surname_student: ''
});

const roles = ['USER', 'STUDENT', 'ACCOUNTANT', 'DEAN', 'ADMIN'];

const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

const fetchStudentData = async (studentId: number) => {
  try {
    const response = await studentService.getStudentById(studentId.toString());
    student.value = response;
  } catch (error) {
    console.error('Failed to fetch student data:', error);
  }
};

const fetchProfile = async () => {
  if (!checkAuth()) return;

  try {
    let email = authStore.email;
    if (email !== null && email !== undefined) {
      const response = await userService.getUserbyEmail(email);
      user.value = response;
      editedUser.value = {
        ...response,
        user_password: ''
      };

      // Если есть student ID, загружаем данные студента
      if (response.user_student_id) {
        await fetchStudentData(response.user_student_id);
      }
    } else {
      console.warn('Email is not set in authStore');
    }
  } catch (error) {
    const axiosError = error as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      console.error('Failed to fetch user ', error);
    }
  }
};

const startEditing = () => {
  isEditing.value = true;
  if (student.value) {
    editedStudent.value = { ...student.value };
  }
};

const cancelEditing = () => {
  isEditing.value = false;
  if (user.value) {
    editedUser.value = {
      ...user.value,
      user_password: ''
    };
  }
  if (student.value) {
    editedStudent.value = { ...student.value };
  }
};

const saveChanges = async () => {
  try {
    if (!editedUser.value.user_password) {
      alert('Пожалуйста, введите пароль для подтверждения изменений');
      return;
    }

    // Обновляем данные пользователя
    const updatedUser = await userService.updateUser(editedUser.value);
    user.value = updatedUser;

    // Если есть данные студента и они изменились, обновляем их
    if (student.value && editedStudent.value) {
      await studentService.updateStudent(editedStudent.value);
      // Обновляем локальные данные студента
      const updatedStudent = await studentService.getStudentById(editedStudent.value.id_num_student.toString());
      student.value = updatedStudent;
    }

    isEditing.value = false;
    alert('Профиль успешно обновлен');
  } catch (error) {
    console.error('Failed to update profile:', error);
    alert('Не удалось обновить профиль');
  }
};

const deleteProfile = async () => {
  if (!confirm('Are you sure you want to delete your profile? This action cannot be undone.')) {
    return;
  }

  try {
    if (!editedUser.value.user_password) {
      alert('Please enter your password to confirm deletion');
      return;
    }

    await userService.deleteUser(editedUser.value);
    authStore.logout();
    router.push('/auth');
  } catch (error) {
    console.error('Failed to delete profile:', error);
    alert('Failed to delete profile');
  }
};

onMounted(() => {
  authStore.initialize();
  fetchProfile();
});
</script>

<template>
  <div class="profile-container">
    <div class="profile-header">
      <div class="profile-avatar">
        <div class="avatar-placeholder">
          {{ user?.user_email?.charAt(0).toUpperCase() }}
        </div>
      </div>
      <h1>{{ user?.user_email }}</h1>
      <div class="profile-role">{{ user?.user_role }}</div>
    </div>

    <div class="profile-content">
      <div class="profile-section">
        <div class="section-header">
          <h2>Информация пользователя</h2>
          <div class="action-buttons">
            <button v-if="!isEditing" @click="startEditing" class="edit-button">
              Редактировать
            </button>
            <button v-if="!isEditing" @click="deleteProfile" class="delete-button">
              Удалить профиль
            </button>
          </div>
        </div>

        <div class="info-grid">
          <div class="info-item">
            <span class="info-label">Email</span>
            <span v-if="!isEditing" class="info-value">{{ user?.user_email }}</span>
            <input v-else v-model="editedUser.user_email" type="email" class="edit-input" />
          </div>

          <div class="info-item">
            <span class="info-label">Login</span>
            <span v-if="!isEditing" class="info-value">{{ user?.user_login }}</span>
            <input v-else v-model="editedUser.user_login" type="text" class="edit-input" />
          </div>

          <div class="info-item">
            <span class="info-label">Роль</span>
            <span v-if="!isEditing" class="info-value">{{ user?.user_role }}</span>
            <select v-else v-model="editedUser.user_role" class="edit-input">
              <option v-for="role in roles" :key="role" :value="role">{{ role }}</option>
            </select>
          </div>

          <div class="info-item">
            <span class="info-label">Студент ID</span>
            <span v-if="!isEditing" class="info-value">{{ user?.user_student_id || 'Не указан' }}</span>
            <input v-else v-model="editedUser.user_student_id" type="number" class="edit-input" />
          </div>

          <div v-if="isEditing" class="info-item">
            <span class="info-label">Пароль (обязателен для изменения)</span>
            <input v-model="editedUser.user_password" type="password" class="edit-input" />
          </div>
        </div>

        <!-- Секция с информацией о студенте -->
        <div v-if="student" class="profile-section student-section">
          <h2>Карточка студента</h2>
          <div class="info-grid">
            <div class="info-item">
              <span class="info-label">Фамилия</span>
              <span v-if="!isEditing" class="info-value">{{ student.second_name_student }}</span>
              <input v-else v-model="editedStudent.second_name_student" type="text" class="edit-input" />
            </div>
            <div class="info-item">
              <span class="info-label">Имя</span>
              <span v-if="!isEditing" class="info-value">{{ student.first_name_student }}</span>
              <input v-else v-model="editedStudent.first_name_student" type="text" class="edit-input" />
            </div>
            <div class="info-item">
              <span class="info-label">Отчество</span>
              <span v-if="!isEditing" class="info-value">{{ student.surname_student }}</span>
              <input v-else v-model="editedStudent.surname_student" type="text" class="edit-input" />
            </div>
            <div class="info-item">
              <span class="info-label">Группа</span>
              <span v-if="!isEditing" class="info-value">{{ student.name_group }}</span>
              <input v-else v-model="editedStudent.name_group" type="text" class="edit-input" disabled/>
            </div>
          </div>
        </div>

        <div v-if="isEditing" class="edit-actions">
          <button @click="saveChanges" class="save-button">Сохраить изменения</button>
          <button @click="cancelEditing" class="cancel-button">Отмена</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-container {
  max-width: 800px;
  margin: 40px auto;
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  animation: fadeIn 0.8s ease-in-out;
}

.profile-header {
  background: linear-gradient(135deg, var(--color-accent) 0%, var(--color-accent-hover) 100%);
  padding: 40px 30px;
  text-align: center;
  color: white;
}

.profile-avatar {
  margin-bottom: 20px;
}

.avatar-placeholder {
  width: 100px;
  height: 100px;
  background: rgba(255, 255, 255, 0.2);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto;
  font-size: 40px;
  font-weight: 500;
  color: white;
  border: 3px solid rgba(255, 255, 255, 0.3);
}

.profile-header h1 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
}

.profile-role {
  margin-top: 8px;
  font-size: 16px;
  opacity: 0.9;
}

.profile-content {
  padding: 30px;
}

.profile-section {
  margin-bottom: 30px;
}

.profile-section h2 {
  color: #333;
  font-size: 20px;
  margin: 0 0 20px 0;
  font-weight: 600;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
}

.info-item {
  background: #f8f9fa;
  padding: 15px;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
}

.info-label {
  color: #666;
  font-size: 14px;
  margin-bottom: 5px;
}

.info-value {
  color: #333;
  font-size: 16px;
  font-weight: 500;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.action-buttons {
  display: flex;
  gap: 10px;
}

.edit-button, .delete-button, .save-button, .cancel-button {
  padding: 8px 16px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease;
}

.edit-button {
  background-color: var(--color-accent);
  color: white;
}

.edit-button:hover {
  background-color: var(--color-accent-hover);
}

.delete-button {
  background-color: #dc3545;
  color: white;
}

.delete-button:hover {
  background-color: #c82333;
}

.save-button {
  background-color: #28a745;
  color: white;
}

.save-button:hover {
  background-color: #218838;
}

.cancel-button {
  background-color: #6c757d;
  color: white;
}

.cancel-button:hover {
  background-color: #5a6268;
}

.edit-input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  margin-top: 4px;
}

.edit-input:focus {
  outline: none;
  border-color: var(--color-accent);
}

.edit-actions {
  display: flex;
  gap: 10px;
  margin-top: 20px;
  justify-content: flex-end;
}

.student-section {
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #eee;
}

.student-section h2 {
  color: var(--color-accent);
  margin-bottom: 20px;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (max-width: 600px) {
  .profile-container {
    margin: 20px;
  }

  .profile-header {
    padding: 30px 20px;
  }

  .profile-content {
    padding: 20px;
  }
}
</style>
