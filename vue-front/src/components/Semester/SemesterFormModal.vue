<script setup lang="ts">
import { ref, watch } from 'vue';
import type { SemesterReq, SemesterResp } from '@/types/semester';

const props = defineProps<{
  show: boolean;
  semester?: SemesterResp;
  mode: 'create' | 'edit';
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'submit', semester: SemesterReq): void;
}>();

const formData = ref<SemesterReq>({
  name_semester: '',
  date_start_semester: '',
  date_end_semester: ''
});

const errors = ref<Record<string, string>>({});

watch(() => props.semester, (newSemester) => {
  if (newSemester && props.mode === 'edit') {
    formData.value = { ...newSemester };
  } else {
    formData.value = {
      name_semester: '',
      date_start_semester: '',
      date_end_semester: ''
    };
  }
}, { immediate: true });

const validateForm = (): boolean => {
  errors.value = {};
  let isValid = true;

  if (!formData.value.name_semester) {
    errors.value.name_semester = 'Название обязательно';
    isValid = false;
  }
  if (!formData.value.date_start_semester) {
    errors.value.date_start_semester = 'Дата начала обязательна';
    isValid = false;
  }
  if (!formData.value.date_end_semester) {
    errors.value.date_end_semester = 'Дата окончания обязательна';
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
        <h2>{{ mode === 'create' ? 'Создать семестр' : 'Редактировать семестр' }}</h2>
        <button class="close-button" @click="handleClose">&times;</button>
      </div>
      <form @submit.prevent="handleSubmit" class="modal-form">
        <div class="form-group">
          <label for="name_semester">Название:</label>
          <input
            type="text"
            id="name_semester"
            v-model="formData.name_semester"
            :disabled="mode === 'edit'"
          />
          <span class="error" v-if="errors.name_semester">{{ errors.name_semester }}</span>
        </div>

        <div class="form-group">
          <label for="date_start_semester">Дата начала:</label>
          <input
            type="date"
            id="date_start_semester"
            v-model="formData.date_start_semester"
          />
          <span class="error" v-if="errors.date_start_semester">{{ errors.date_start_semester }}</span>
        </div>

        <div class="form-group">
          <label for="date_end_semester">Дата окончания:</label>
          <input
            type="date"
            id="date_end_semester"
            v-model="formData.date_end_semester"
          />
          <span class="error" v-if="errors.date_end_semester">{{ errors.date_end_semester }}</span>
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
