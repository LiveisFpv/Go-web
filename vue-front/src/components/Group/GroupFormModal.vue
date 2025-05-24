<script setup lang="ts">
import { ref, watch } from 'vue';
import type { GroupReq, GroupResp } from '@/types/group';

const props = defineProps<{
  show: boolean;
  group?: GroupResp;
  mode: 'create' | 'edit';
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'submit', group: GroupReq): void;
}>();

const formData = ref<GroupReq>({
  name_group: '',
  studies_direction_group: '',
  studies_profile_group: '',
  start_date_group: new Date(),
  studies_period_group: 4,
});

const errors = ref<Record<string, string>>({});

watch(() => props.group, (newGroup) => {
  if (newGroup && props.mode === 'edit') {
    formData.value = { ...newGroup };
  } else {
    formData.value = {
      name_group: '',
      studies_direction_group: '',
      studies_profile_group: '',
      start_date_group: new Date(),
      studies_period_group: 4,
    };
  }
}, { immediate: true });

const validateForm = (): boolean => {
  errors.value = {};
  let isValid = true;

  if (!formData.value.name_group) {
    errors.value.name_group = 'Название группы обязательно';
    isValid = false;
  }
  if (!formData.value.studies_direction_group) {
    errors.value.studies_direction_group = 'Направление обучения обязательно';
    isValid = false;
  }
  if (!formData.value.studies_profile_group) {
    errors.value.studies_profile_group = 'Профиль обучения обязателен';
    isValid = false;
  }
  if (!formData.value.start_date_group) {
    errors.value.start_date_group = 'Дата начала обучения обязательна';
    isValid = false;
  }
  if (!formData.value.studies_period_group || formData.value.studies_period_group <= 0) {
    errors.value.studies_period_group = 'Период обучения должен быть положительным числом';
    isValid = false;
  }

  return isValid;
};

const handleSubmit = () => {
  if (validateForm()) {
    emit('submit', {...formData.value});
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
        <h2>{{ mode === 'create' ? 'Создать группу' : 'Редактировать группу' }}</h2>
        <button class="close-button" @click="handleClose">&times;</button>
      </div>
      <form @submit.prevent="handleSubmit" class="modal-form">
        <div class="form-group">
          <label for="name_group">Название группы</label>
          <input id="name_group"
          type="text"
          v-model="formData.name_group"
          :disabled="mode === 'edit'"
          />
          <span class="error" v-if="errors.name_group">{{ errors.name_group }}</span>
        </div>

        <div class="form-group">
          <label for="studies_direction_group">Направление обучения</label>
          <input id="studies_direction_group"
          type="text"
          v-model="formData.studies_direction_group"
          />
          <span class="error" v-if="errors.studies_direction_group">{{ errors.studies_direction_group }}</span>
        </div>

        <div class="form-group">
          <label for="studies_profile_group">Профиль обучения</label>
          <input id="studies_profile_group"
          type="text"
          v-model="formData.studies_profile_group"
          />
          <span class="error" v-if="errors.studies_profile_group">{{ errors.studies_profile_group }}</span>

        </div>

        <div class="form-group">
          <label for="start_date_group">Дата начала обучения</label>
          <input id="start_date_group"
          type="date"
          v-model="formData.start_date_group"
          />
          <span class="error" v-if="errors.start_date_group">{{ errors.start_date_group }}</span>
        </div>

        <div class="form-group">
          <label for="studies_period_group">Срок обучения (в годах)</label>
          <input id="studies_period_group"
          type="number"
          v-model.number="formData.studies_period_group"
          />
          <span class="error" v-if="errors.studies_period_group">{{ errors.studies_period_group }}</span>
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
