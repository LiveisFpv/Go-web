<script setup lang="ts">
import { ref, watch } from 'vue';
import type { StudentReq, StudentResp } from '@/types/student';

const props = defineProps<{
  show: boolean;
  student?: StudentResp;
  mode: 'create' | 'edit';
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'submit', student: StudentReq): void;
}>();

const formData = ref<StudentReq>({
  id_num_student: 0,
  name_group: '',
  email_student: '',
  second_name_student: '',
  first_name_student: '',
  surname_student: ''
});

const errors = ref<Record<string, string>>({});

watch(() => props.student, (newStudent) => {
  if (newStudent && props.mode === 'edit') {
    formData.value = { ...newStudent };
  } else {
    formData.value = {
      id_num_student: 0,
      name_group: '',
      email_student: '',
      second_name_student: '',
      first_name_student: '',
      surname_student: ''
    };
  }
}, { immediate: true });

const validateForm = (): boolean => {
  errors.value = {};
  let isValid = true;

  if (!formData.value.id_num_student) {
    errors.value.id_num_student = 'Номер билета обязателен';
    isValid = false;
  }
  if (!formData.value.name_group) {
    errors.value.name_group = 'Группа обязательна';
    isValid = false;
  }
  if (!formData.value.email_student) {
    errors.value.email_student = 'Email обязателен';
    isValid = false;
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.value.email_student)) {
    errors.value.email_student = 'Некорректный email';
    isValid = false;
  }
  if (!formData.value.second_name_student) {
    errors.value.second_name_student = 'Фамилия обязательна';
    isValid = false;
  }
  if (!formData.value.first_name_student) {
    errors.value.first_name_student = 'Имя обязательно';
    isValid = false;
  }
  if (!formData.value.surname_student) {
    errors.value.surname_student = 'Отчество обязательно';
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
        <h2>{{ mode === 'create' ? 'Создать студента' : 'Редактировать студента' }}</h2>
        <button class="close-button" @click="handleClose">&times;</button>
      </div>
      <form @submit.prevent="handleSubmit" class="modal-form">
        <div class="form-group">
          <label for="id_num_student">Номер билета:</label>
          <input
            type="number"
            id="id_num_student"
            v-model.number="formData.id_num_student"
            :disabled="mode === 'edit'"
          />
          <span class="error" v-if="errors.id_num_student">{{ errors.id_num_student }}</span>
        </div>

        <div class="form-group">
          <label for="name_group">Группа:</label>
          <input
            type="text"
            id="name_group"
            v-model="formData.name_group"
          />
          <span class="error" v-if="errors.name_group">{{ errors.name_group }}</span>
        </div>

        <div class="form-group">
          <label for="email_student">Email:</label>
          <input
            type="email"
            id="email_student"
            v-model="formData.email_student"
          />
          <span class="error" v-if="errors.email_student">{{ errors.email_student }}</span>
        </div>

        <div class="form-group">
          <label for="second_name_student">Фамилия:</label>
          <input
            type="text"
            id="second_name_student"
            v-model="formData.second_name_student"
          />
          <span class="error" v-if="errors.second_name_student">{{ errors.second_name_student }}</span>
        </div>

        <div class="form-group">
          <label for="first_name_student">Имя:</label>
          <input
            type="text"
            id="first_name_student"
            v-model="formData.first_name_student"
          />
          <span class="error" v-if="errors.first_name_student">{{ errors.first_name_student }}</span>
        </div>

        <div class="form-group">
          <label for="surname_student">Отчество:</label>
          <input
            type="text"
            id="surname_student"
            v-model="formData.surname_student"
          />
          <span class="error" v-if="errors.surname_student">{{ errors.surname_student }}</span>
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
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.modal-header h2 {
  margin: 0;
  color: var(--color-accent);
}

.close-button {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #666;
}

.modal-form {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.form-group label {
  font-weight: 500;
  color: #333;
}

.form-group input {
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

.form-group input:focus {
  outline: none;
  border-color: var(--color-accent);
}

.error {
  color: #dc3545;
  font-size: 12px;
}

.form-actions {
  display: flex;
  gap: 10px;
  margin-top: 20px;
}

.submit-button, .cancel-button {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: 500;
}

.submit-button {
  background-color: var(--color-accent);
  color: white;
}

.submit-button:hover {
  background-color: var(--color-accent-hover);
}

.cancel-button {
  background-color: #f1f1f1;
  color: #333;
}

.cancel-button:hover {
  background-color: #e1e1e1;
}
</style> 