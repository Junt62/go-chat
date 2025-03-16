<script setup lang="ts">
import { ref, defineEmits } from 'vue'

const message = ref<string>('')
const emit = defineEmits<{ (e: 'send', message: string): void }>()

if (message.value.trim()) {
  emit('send', message.value)
  message.value = ''
}
</script>
