<script setup lang="ts">
import type { CategoryReq, CategoryResp } from '@/types/category';
import { ref, watch } from 'vue';


const props = defineProps<{
  show: boolean;
  category?: CategoryResp;
  mode: 'create' | 'edit';
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'submit', category: CategoryReq): void;
}>();

const formData = ref<CategoryReq>({
  id_category: 0,
  achivments_type_category: '',
  score_category: 0,
});

const errors = ref<Record<string, string>>({});

watch(() => props.category, (newCategory)=>{
  if (newCategory && props.mode ==='edit'){
    formData.value = {...newCategory};
  } else {
    formData.value = {
      id_category:0,
      achivments_type_category:'',
      score_category:0,
    };
  }
}, {immediate: true});

const validateForm = (): boolean => {
  errors.value = {};
  let isValid = true;

  if (!formData.value.achivments_type_category) {
    errors.value.achivments_type_category = 'Тип достижения обязателен';
    isValid = false;
  }
  if (formData.value.score_category <= 0) {
    errors.value.score_category = 'Баллы должны быть больше нуля и меньше 10';
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
        <h2>{{ mode === 'create' ? 'Создать категорию' : 'Редактировать категорию' }}</h2>
        <button class="close-button" @click="handleClose">&times;</button>
      </div>
      <form @submit.prevent="handleSubmit" class="modal-form">
        <div class="hidden form-group">
          <label for="id_category">ID категории:</label>
          <input
            type="number"
            id="id_category"
            v-model.number="formData.id_category"
            :class="{ 'is-invalid': errors.id_category }"
            readonly
          />
          <span class="error-message" v-if="errors.id_category">{{ errors.id_category }}</span>
        </div>
        <div class="form-group">
          <label for="achivments_type_category">Тип достижения:</label>
          <input
            type="text"
            id="achivments_type_category"
            v-model="formData.achivments_type_category"
            :class="{ 'is-invalid': errors.achivments_type_category }"
          />
          <span class="error-message" v-if="errors.achivments_type_category">{{ errors.achivments_type_category }}</span>
        </div>
        <div class="form-group">
          <label for="score_category">Баллы:</label>
          <input
            type="number"
            id="score_category"
            v-model.number="formData.score_category"
            :class="{ 'is-invalid': errors.score_category }"
          />
          <span class="error-message" v-if="errors.score_category">{{ errors.score_category }}</span>
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
