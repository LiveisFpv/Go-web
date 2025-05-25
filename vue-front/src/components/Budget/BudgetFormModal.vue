<script setup lang="ts">
import { ref, watch } from 'vue';
import type { BudgetReq, BudgetResp } from '@/types/budget';

const props = defineProps<{
  show: boolean;
  budget?: BudgetResp;
  mode: 'create' | 'edit';
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'submit', budget: BudgetReq): void;
}>();

const formData = ref<BudgetReq>({
  id_budget: 0,
  size_budget: 0,
  type_scholarship_budget: '',
  name_semester: ''
});

const errors = ref<Record<string, string>>({});

watch(() => props.budget, (newBudget) => {
  if (newBudget && props.mode === 'edit') {
    formData.value = { ...newBudget };
  } else {
    formData.value = {
      id_budget: 0,
      size_budget: 0,
      type_scholarship_budget: '',
      name_semester: ''
    };
  }
}, { immediate: true });

const validateForm = (): boolean => {
  errors.value = {};
  let isValid = true;

  if (formData.value.size_budget <= 0) {
    errors.value.size_budget = 'Размер бюджета должен быть больше 0';
    isValid = false;
  }
  if (!formData.value.type_scholarship_budget) {
    errors.value.type_scholarship_budget = 'Тип стипендии обязателен';
    isValid = false;
  }
  if (!formData.value.name_semester) {
    errors.value.name_semester = 'Семестр обязателен';
    isValid = false;
  }

  return isValid;
};

const handleSubmit = () => {
  if (validateForm()) {
    emit('submit', { ...formData.value });
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
        <h2>{{ mode === 'create' ? 'Создать бюджет' : 'Редактировать бюджет' }}</h2>
        <button class="close-button" @click="handleClose">&times;</button>
      </div>
      <form @submit.prevent="handleSubmit" class="modal-form">
        <div class="form-group">
          <label for="size_budget">Размер бюджета:</label>
          <input
            type="number"
            id="size_budget"
            v-model.number="formData.size_budget"
            min="0"
            step="0.01"
          />
          <span class="error" v-if="errors.size_budget">{{ errors.size_budget }}</span>
        </div>

        <div class="form-group">
          <label for="type_scholarship_budget">Тип стипендии:</label>
          <input
            type="text"
            id="type_scholarship_budget"
            v-model="formData.type_scholarship_budget"
          />
          <span class="error" v-if="errors.type_scholarship_budget">{{ errors.type_scholarship_budget }}</span>
        </div>

        <div class="form-group">
          <label for="name_semester">Семестр:</label>
          <input
            type="text"
            id="name_semester"
            v-model="formData.name_semester"
          />
          <span class="error" v-if="errors.name_semester">{{ errors.name_semester }}</span>
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
