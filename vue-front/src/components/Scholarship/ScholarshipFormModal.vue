<script setup lang="ts">
import { ref, watch, computed } from 'vue';
import type { ScholarshipReq, ScholarshipResp } from '@/types/scholarship';
import { scholarshipService } from '@/services/scholarshipService';
import type { StudentResp } from '@/types/student';
import { useAuthStore } from '@/stores/auth';
import router from '@/router';
import { studentService } from '@/services/studentService';
import type { AxiosError } from 'axios';
import type { BudgetResp } from '@/types/budget';
import { budgetService } from '@/services/budgetService';
import { semesterService } from '@/services/semesterService';
import type { SemesterResp } from '@/types/semester';

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
const students = ref<StudentResp[]>([]);
const semesters = ref<SemesterResp[]>([]);
const budget = ref<BudgetResp[]>([]);
const authStore = useAuthStore();

const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

const getSemesters = async () => {
  if (!checkAuth()) return;
  if (semesters.value.length > 0) return;
  try {
    const response = await semesterService.getSemesters(1, 1000);
    semesters.value = response.data;
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      console.error('Failed to fetch semesters', err);
    }
  }
};

const getBudget = async () => {
  if (!checkAuth()) return;
  if (budget.value.length > 0) return;
  try {
    const response = await budgetService.getBudgets(1, 1000);
    budget.value = response.data;
  } catch (err) {
    const axiosError = err as AxiosError;
    if (axiosError.response?.status === 401) {
      authStore.logout();
      router.push('/auth');
    } else {
      console.error('Failed to fetch budget', err);
    }
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

const filteredBudgets = computed(() => {
  if (!formData.value.name_semester) return [];
  return budget.value.filter(bud => {
    return bud.name_semester === formData.value.name_semester;
  });
});

watch(() => props.scholarship, (newScholarship) => {
  if (newScholarship && props.mode === 'edit') {
    formData.value = {
      id_scholarship: newScholarship.id_scholarship,
      id_num_student: newScholarship.id_num_student,
      name_semester: newScholarship.name_semester,
      size_scholarshp: newScholarship.size_scholarshp,
      id_budget: newScholarship.id_budget
    };
    // Загружаем данные для выпадающих списков
    getSemesters();
    getBudget();
    onclick();
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

const handleSubmit = async () => {
  if (validateForm()) {
    try {
      if (props.mode === 'create') {
        await scholarshipService.createScholarship(formData.value);
      } else {
        await scholarshipService.updateScholarship(formData.value);
      }
      emit('submit', formData.value);
      handleClose();
    } catch (error) {
      console.error('Error saving scholarship:', error);
      // Можно добавить отображение ошибки пользователю
    }
  }
};

const handleClose = () => {
  emit('close');
};
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
          <label for="id_num_student">Cтудент:</label>
          <select v-model.number="formData.id_num_student" id="id_num_student" @click="onclick">
            <option v-for="student in students" :key="student.id_num_student" :value="student.id_num_student">{{ student.second_name_student+" "+student.first_name_student+" "+student.surname_student+" "+student.name_group+" "+student.id_num_student }}</option>
          </select>
          <span class="error-message" v-if="errors.id_num_student">{{ errors.id_num_student }}</span>
        </div>

        <div class="form-group">
          <label for="name_semester">Семестр:</label>
          <select
            id="name_semester"
            v-model="formData.name_semester"
            required
            @click="getSemesters"
            @change="formData.id_budget = 0"
          >
            <option value="">Выберите семестр</option>
            <option v-for="semester in semesters"
                    :key="semester.name_semester"
                    :value="semester.name_semester">
              {{ semester.name_semester }}
            </option>
          </select>
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
          <label for="id_budget">Тип стипендии</label>
          <select
            v-model.number="formData.id_budget"
            id="id_budget"
            @click="getBudget"
            :disabled="!formData.name_semester"
          >
            <option value="">Выберите тип стипендии</option>
            <option v-for="bud in filteredBudgets"
                    :key="bud.id_budget"
                    :value="bud.id_budget">
              {{ bud.type_scholarship_budget }}
            </option>
          </select>
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
