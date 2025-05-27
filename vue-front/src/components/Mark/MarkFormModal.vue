<script setup lang="ts">
import router from '@/router';
import { studentService } from '@/services/studentService';
import { useAuthStore } from '@/stores/auth';
import type { MarkReq, MarkResp } from '@/types/mark';
import type { StudentResp } from '@/types/student';
import type { AxiosError } from 'axios';
import { ref, watch, computed } from 'vue';

const authStore = useAuthStore();

const isStudent = computed(() => {
  return authStore.user_role === 'STUDENT';
});

const checkAuth = () => {
  if (!authStore.isAuthenticated || isStudent.value) {
    router.push('/auth');
    return false;
  }
  return true;
};

const props = defineProps<{
  show: boolean;
  mark?: MarkResp;
  mode: 'create' | 'edit';
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'submit', mark: MarkReq): void;
}>();

const formData = ref<MarkReq>({
  id_mark: 0,
  id_num_student: 0,
  name_semester: '',
  lesson_name_mark: '',
  score_mark: 0,
  type_mark: '',
  type_exam: ''
});

const errors = ref<Record<string, string>>({});
watch(() => props.mark, (newMark) => {
  if (newMark && props.mode === 'edit') {
    formData.value = { ...newMark };
  } else {
    formData.value = {
      id_mark: 0,
      id_num_student: 0,
      name_semester: '',
      lesson_name_mark: '',
      score_mark: 0,
      type_mark: '',
      type_exam: ''
    };
  }
}, { immediate: true });

const validateForm = (): boolean => {
  errors.value = {};
  let isValid = true;

  if (!formData.value.id_num_student) {
    errors.value.id_num_student = 'ID студента обязателен';
    isValid = false;
  }
  if (!formData.value.name_semester) {
    errors.value.name_semester = 'Семестр обязателен';
    isValid = false;
  }
  if (!formData.value.lesson_name_mark) {
    errors.value.lesson_name_mark = 'Название предмета обязательно';
    isValid = false;
  }
  if (formData.value.score_mark < 0 || formData.value.score_mark > 100) {
    errors.value.score_mark = 'Оценка должна быть от 0 до 100';
    isValid = false;
  }
  if (!formData.value.type_exam) {
    errors.value.type_exam = 'Тип экзамена обязателен';
    isValid = false;
  }

  return isValid;
};

const handleSubmit = () => {
  if (!checkAuth()) return;
  if (validateForm()) {
    emit('submit', formData.value);
  }
};

const handleClose = () => {
  emit('close');
};

const students = ref<StudentResp[]>([]);
const onclick = async () => {
  if (students.value.length > 0) return;
  try{
    const response = await studentService.getStudents(1, 1000);
    students.value = response.data
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      console.error('Failed to fetch students', err);
    }
  }
}
</script>

<template>
  <div v-if="show && !isStudent" class="modal-overlay" @click="handleClose">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>{{ mode === 'create' ? 'Создать оценку' : 'Редактировать оценку' }}</h2>
        <button class="close-button" @click="handleClose">&times;</button>
      </div>
      <form @submit.prevent="handleSubmit" class="modal-form">
        <div class="hidden form-group">
          <label for="id_mark">ID оценки:</label>
          <input v-model.number="formData.id_mark" type="number" id="id_mark" placeholder="Введите ID оценки" />
          <span class="error-message" v-if="errors.id_mark">{{ errors.id_mark }}</span>
        </div>
        <div class="form-group">
          <label for="id_num_student">ID студента:</label>
          <select v-model.number="formData.id_num_student" id="id_num_student" @click="onclick">
            <option v-for="student in students" :key="student.id_num_student" :value="student.id_num_student">{{ student.second_name_student+" "+student.first_name_student+" "+student.surname_student+" "+student.name_group+" "+student.id_num_student }}</option>
          </select>
          <span class="error-message" v-if="errors.id_num_student">{{ errors.id_num_student }}</span>
        </div>
        <div class="form-group">
          <label for="name_semester">Семестр:</label>
          <input v-model="formData.name_semester" type="text" id="name_semester" placeholder="Введите семестр" />
          <span class="error-message" v-if="errors.name_semester">{{ errors.name_semester }}</span>
        </div>
        <div class="form-group">
          <label for="lesson_name_mark">Название предмета:</label>
          <input v-model="formData.lesson_name_mark" type="text" id="lesson_name_mark" placeholder="Введите название предмета" />
          <span class="error-message" v-if="errors.lesson_name_mark">{{ errors.lesson_name_mark }}</span>
        </div>
        <div class="form-group">
          <label for="score_mark">Оценка:</label>
          <input v-model.number="formData.score_mark" type="number" id="score_mark" placeholder="Введите оценку" />
          <span class="error-message" v-if="errors.score_mark">{{ errors.score_mark }}</span>
        </div>
        <div class="form-group">
          <label for="type_exam">Тип экзамена:</label>
          <input v-model="formData.type_exam" type="text" id="type_exam" placeholder="Введите тип экзамена" />
          <span class="error-message" v-if="errors.type_exam">{{ errors.type_exam }}</span>
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
