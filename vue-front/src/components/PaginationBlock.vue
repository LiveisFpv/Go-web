<script setup lang="ts">
const props = defineProps<{
  page: number;
  limit: number;
  total: number;
}>();

const emit = defineEmits<{
  (e: 'update:page', value: number): void;
}>();
</script>

<template>
  <div class="pagination">
    <button @click="emit('update:page', props.page - 1)" :disabled="props.page <= 1"><</button>
    <span>Страница {{ props.page }} из {{ Math.ceil(props.total / props.limit) }}</span>
    <button @click="emit('update:page', props.page + 1)" :disabled="props.page >= Math.ceil(props.total / props.limit)">></button>
  </div>
</template>
<style scoped>
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 20px 0;
}
.pagination button {
  background-color: var(--color-accent);
  border-radius: 10px;
  color: white;
  border: none;
  padding: 10px 20px;
  margin: 0 5px;
  cursor: pointer;
  font-weight: 600;
  font-size: 1.2em;
  transition: transform 0.3s ease;
}
.pagination button:hover:enabled {
  transform: scale(1.1);
}
.pagination button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
.pagination span {
  margin: 0 10px;
  font-size: 1.2em;
  font-weight: 600;
  color: var(--color-text);
  text-align: center;
}
</style>
