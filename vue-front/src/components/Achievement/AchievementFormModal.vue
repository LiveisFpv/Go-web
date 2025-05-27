<script setup lang="ts">
import { ref, watch, computed, onMounted } from 'vue';
import type { AchivementReq, AchivementResp } from '@/types/achievement';
import { achievementService } from '@/services/achievementService';
import type { StudentResp } from '@/types/student';
import { useAuthStore } from '@/stores/auth';
import router from '@/router';
import { studentService } from '@/services/studentService';
import type { AxiosError } from 'axios';
import type { CategoryResp } from '@/types/category';
import { categoryService } from '@/services/categoryService';
import { userService } from '@/services/userService';

const props = defineProps<{
  show: boolean;
  achievement?: AchivementResp;
  mode: 'create' | 'edit';
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'submit', achievement: AchivementReq): void;
}>();

const formData = ref<AchivementReq>({
  id_achivment: 0,
  id_num_student: 0,
  id_category: 0,
  name_achivement: '',
  date_achivement: ''
});

const errors = ref<Record<string, string>>({});
const students = ref<StudentResp[]>([]);
const categories = ref<CategoryResp[]>([]);
const authStore = useAuthStore();

const isStudent = computed(() => {
  return authStore.user_role === 'STUDENT';
});

const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

const getCategories = async () => {
  if (!checkAuth()) return;
  if (categories.value.length > 0) return;
  try {
    const response = await categoryService.getCategories(1, 1000);
    categories.value = response.data;
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      console.error('Failed to fetch categories', err);
    }
  }
};

const getStudentId = async () => {
  if (!authStore.email) return;
  try {
    const userData = await userService.getUserbyEmail(authStore.email);
    if (userData.user_student_id) {
      formData.value.id_num_student = userData.user_student_id;
    }
  } catch (err) {
    console.error('Failed to fetch user data:', err);
  }
};

const onclick = async () => {
  if (!checkAuth()) return;
  if (students.value.length > 0) return;
  try {
    const response = await studentService.getStudents(1, 1000);
    students.value = response.data;
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      console.error('Failed to fetch students', err);
    }
  }
};

onMounted(async () => {
  if (isStudent.value) {
    await getStudentId();
  }
  await getCategories();
});

watch(() => props.achievement, (newAchievement) => {
  if (newAchievement && props.mode === 'edit') {
    formData.value = {
      id_achivment: newAchievement.id_achivment,
      id_num_student: newAchievement.id_num_student,
      id_category: newAchievement.id_category,
      name_achivement: newAchievement.name_achivement,
      date_achivement: newAchievement.date_achivement
    };
  } else {
    formData.value = {
      id_achivment: 0,
      id_num_student: 0,
      id_category: 0,
      name_achivement: '',
      date_achivement: ''
    };
    if (isStudent.value) {
      getStudentId();
    }
  }
}, { immediate: true });

const validateForm = (): boolean => {
  errors.value = {};
  let isValid = true;

  if (!formData.value.id_num_student || formData.value.id_num_student <= 0) {
    errors.value.id_num_student = 'Введите корректный ID студента';
    isValid = false;
  }
  if (!formData.value.id_category || formData.value.id_category <= 0) {
    errors.value.id_category = 'Введите корректную категорию';
    isValid = false;
  }
  if (!formData.value.name_achivement || formData.value.name_achivement.trim() === '') {
    errors.value.name_achivement = 'Введите название достижения';
    isValid = false;
  }
  if (!formData.value.date_achivement || formData.value.date_achivement.trim() === '') {
    errors.value.date_achivement = 'Введите дату достижения';
    isValid = false;
  }

  return isValid;
};

const handleSubmit = () =>{
  if (validateForm()) {
    emit('submit', formData.value);
  }
}

const handleClose = () => {
  emit('close');
};
</script>

<template>
  <div v-if="show" class="modal-overlay" @click="handleClose">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>{{ mode === 'create' ? 'Создать достижение' : 'Редактировать достижение' }}</h2>
        <button class="close-button" @click="handleClose">&times;</button>
      </div>

      <form @submit.prevent="handleSubmit" class="modal-form">
        <div class="hidden form-group">
          <label for="id_achivment">ID достижения:</label>
          <input
            type="number"
            id="id_achivment"
            v-model.number="formData.id_achivment"
          />
          <span class="error-message" v-if="errors.id_achivment">
            {{ errors.id_achivment }}
          </span>
        </div>

        <div class="form-group" v-if="!isStudent">
          <label for="id_num_student">Студент:</label>
          <select v-model.number="formData.id_num_student" id="id_num_student" @click="onclick">
            <option v-for="student in students" :key="student.id_num_student" :value="student.id_num_student">
              {{ student.second_name_student+" "+student.first_name_student+" "+student.surname_student+" "+student.name_group+" "+student.id_num_student }}
            </option>
          </select>
          <span class="error-message" v-if="errors.id_num_student">{{ errors.id_num_student }}</span>
        </div>

        <div class="form-group">
          <label for="id_category">Категория:</label>
          <select
            v-model.number="formData.id_category"
            id="id_category"
            @click="getCategories"
          >
            <option value="">Выберите категорию</option>
            <option v-for="category in categories"
                    :key="category.id_category"
                    :value="category.id_category">
              {{ category.achivments_type_category }}
            </option>
          </select>
          <span class="error-message" v-if="errors.id_category">
            {{ errors.id_category }}
          </span>
        </div>

        <div class="form-group">
          <label for="name_achivement">Название достижения:</label>
          <input
            type="text"
            id="name_achivement"
            v-model="formData.name_achivement"
            required
          />
          <span class="error-message" v-if="errors.name_achivement">
            {{ errors.name_achivement }}
          </span>
        </div>

        <div class="form-group">
          <label for="date_achivement">Дата достижения:</label>
          <input
            type="date"
            id="date_achivement"
            v-model="formData.date_achivement"
            required
          />
          <span class="error-message" v-if="errors.date_achivement">
            {{ errors.date_achivement }}
          </span>
        </div>

        <div class="form-actions">
          <button type="submit" class="submit-button">
            {{ mode === 'create' ? 'Создать' : 'Сохранить' }}
          </button>
          <button type="button" class="cancel-button" @click="handleClose">Отмена</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
@import '../../assets/modal.css';
</style>
