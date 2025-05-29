<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import type { AssignScholarshipReq } from '@/types/scholarship';
import { scholarshipService } from '@/services/scholarshipService';
import { semesterService } from '@/services/semesterService';
import { budgetService } from '@/services/budgetService';
import type { SemesterResp } from '@/types/semester';
import type { BudgetResp } from '@/types/budget';
import { useAuthStore } from '@/stores/auth';
import router from '@/router';

const props = defineProps<{
  show: boolean;
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'refresh'): void;
}>();

const formData = ref<AssignScholarshipReq>({
  current_semester: '',
  budget_type: ''
});

const errors = ref<Record<string, string>>({});
const semesters = ref<SemesterResp[]>([]);
const budgets = ref<BudgetResp[]>([]);
const authStore = useAuthStore();

// Add computed property for filtered budgets
const filteredBudgets = computed(() => {
  if (!formData.value.current_semester) return [];
  return budgets.value.filter(budget => 
    budget.name_semester === formData.value.current_semester
  );
});

// Reset budget_type when semester changes
watch(() => formData.value.current_semester, () => {
  formData.value.budget_type = '';
});

const checkAuth = () => {
  if (!authStore.isAuthenticated) {
    router.push('/auth');
    return false;
  }
  return true;
};

const getSemesters = async () => {
  if (!checkAuth()) return;
  try {
    const response = await semesterService.getSemesters(1, 1000);
    semesters.value = response.data;
  } catch (err) {
    console.error('Failed to fetch semesters:', err);
  }
};

const getBudgets = async () => {
  if (!checkAuth()) return;
  try {
    const response = await budgetService.getBudgets(1, 1000);
    budgets.value = response.data;
  } catch (err) {
    console.error('Failed to fetch budgets:', err);
  }
};

const validateForm = (): boolean => {
  errors.value = {};
  let isValid = true;

  if (!formData.value.current_semester) {
    errors.value.current_semester = 'Выберите семестр';
    isValid = false;
  }
  if (!formData.value.budget_type) {
    errors.value.budget_type = 'Выберите тип стипендии';
    isValid = false;
  }

  return isValid;
};

const handleSubmit = async () => {
  if (!checkAuth()) return;
  if (validateForm()) {
    try {
      await scholarshipService.assignScholarships(formData.value);
      emit('refresh');
      handleClose();
    } catch (error) {
      console.error('Error generating scholarships:', error);
    }
  }
};

const handleClose = () => {
  emit('close');
};

// Load data when modal opens
const loadData = async () => {
  await getSemesters();
  await getBudgets();
};

if (props.show) {
  loadData();
}
</script>

<template>
  <div v-if="show" class="modal-overlay" @click="handleClose">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <h2>Генерация стипендий</h2>
        <button class="close-button" @click="handleClose">&times;</button>
      </div>

      <form @submit.prevent="handleSubmit" class="modal-form">
        <div class="form-group">
          <label for="current_semester">Семестр:</label>
          <select
            id="current_semester"
            v-model="formData.current_semester"
            required
            @click="getSemesters"
          >
            <option value="">Выберите семестр</option>
            <option v-for="semester in semesters"
                    :key="semester.name_semester"
                    :value="semester.name_semester">
              {{ semester.name_semester }}
            </option>
          </select>
          <span class="error-message" v-if="errors.current_semester">
            {{ errors.current_semester }}
          </span>
        </div>

        <div class="form-group">
          <label for="budget_type">Тип стипендии:</label>
          <select
            id="budget_type"
            v-model="formData.budget_type"
            required
            @click="getBudgets"
            :disabled="!formData.current_semester"
          >
            <option value="">Выберите тип стипендии</option>
            <option v-for="budget in filteredBudgets"
                    :key="budget.type_scholarship_budget"
                    :value="budget.type_scholarship_budget">
              {{ budget.type_scholarship_budget }}
            </option>
          </select>
          <span class="error-message" v-if="errors.budget_type">
            {{ errors.budget_type }}
          </span>
        </div>

        <div class="form-actions">
          <button type="submit" class="submit-button">Сгенерировать</button>
          <button type="button" class="cancel-button" @click="handleClose">Отмена</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
@import '../../assets/modal.css';
</style> 