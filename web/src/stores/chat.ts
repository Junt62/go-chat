import { defineStore } from 'pinia'
import type { Message } from '@/types/chat'

export const useChatStore = defineStore('chat', {
  state: () => ({
    messages: [] as Message[],
  }),
  actions: {
    addMessage(message: Message) {
      this.messages.push(message)
    },
  },
})
