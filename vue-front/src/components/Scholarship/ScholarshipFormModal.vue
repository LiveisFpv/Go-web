<script setup lang="ts">
import { ref, watch } from 'vue';
import type { ScholarshipReq, ScholarshipResp } from '@/types/scholarship';
import { scholarshipService } from '@/services/scholarshipService';
import type { StudentResp } from '@/types/student';
import { useAuthStore } from '@/stores/auth';
import router from '@/router';
import { studentService } from '@/services/studentService';
import type { AxiosError } from 'axios';

const props = defineProps<{
  show: boolean;
  scholarship?: ScholarshipResp;
  mode: 'create' | 'edit';
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'submit', scholarship: ScholarshipReq): void;
}>();

const formData = ref<ScholarshipReq>({
  id_scholarship: 0,
  id_num_student: 0,
  name_semester: '',
  size_scholarshp: 0,
  id_budget: 0
});

const errors = ref<Record<string, string>>({});

watch(() => props.scholarship, (newScholarship) => {
  if (newScholarship && props.mode === 'edit') {
    formData.value = { ...newScholarship };
  } else {
    formData.value = {
      id_scholarship: 0,
      id_num_student: 0,
      name_semester: '',
      size_scholarshp: 0,
      id_budget: 0
    };
  }
}, { immediate: true });

const validateForm = (): boolean => {
  errors.value = {};
  let isValid = true;

  if (!formData.value.id_num_student || formData.value.id_num_student <= 0) {
    errors.value.id_num_student = 'Введите корректный ID студента';
    isValid = false;
  }
  if (!formData.value.name_semester || formData.value.name_semester.trim() === '') {
    errors.value.name_semester = 'Введите семестр';
    isValid = false;
  }
  if (!formData.value.size_scholarshp || formData.value.size_scholarshp <= 0) {
    errors.value.size_scholarshp = 'Введите корректный размер стипендии';
    isValid = false;
  }
  if (!formData.value.id_budget || formData.value.id_budget <= 0) {
    errors.value.id_budget = 'Введите корректный ID бюджета';
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

const students = ref<StudentResp[]>([]);
const authStore = useAuthStore();
const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

const onclick = async () => {
  if (!checkAuth()) return;
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
  <div v-if="show" class="modal-overlay" @click="handleClose">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>{{ mode === 'create' ? 'Создать стипендию' : 'Редактировать стипендию' }}</h2>
        <button class="close-button" @click="handleClose">&times;</button>
      </div>

      <form @submit.prevent="handleSubmit" class="modal-form">
        <div class="hidden form-group">
          <label for="id_scholarship">ID стипендии:</label>
          <input
            type="number"
            id="id_scholarship"
            v-model.number="formData.id_scholarship"
          />
          <span class="error-message" v-if="errors.id_scholarship">
            {{ errors.id_scholarship }}
          </span>
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
          <input
            type="text"
            id="name_semester"
            v-model="formData.name_semester"
            required
          />
          <span class="error-message" v-if="errors.name_semester">
            {{ errors.name_semester }}
          </span>
        </div>

        <div class="form-group">
          <label for="size_scholarshp">Размер:</label>
          <input
            type="number"
            id="size_scholarshp"
            v-model.number="formData.size_scholarshp"
            step="0.01"
            required
          />
          <span class="error-message" v-if="errors.size_scholarshp">
            {{ errors.size_scholarshp }}
          </span>
        </div>

        <div class="form-group">
          <label for="id_budget">ID бюджета:</label>
          <input
            type="number"
            id="id_budget"
            v-model.number="formData.id_budget"
            required
          />
          <span class="error-message" v-if="errors.id_budget">
            {{ errors.id_budget }}
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
