import type { Message } from '@/types/chat'

let socket: WebSocket | null = null

export const connectWebSocket = (userId: string, onMessage: (msg: Message) => void): void => {
  socket = new WebSocket(`ws://localhost:8080/chat?userId=${userId}`)

  socket.onmessage = (event) => {
    const message: Message = JSON.parse(event.data)
    onMessage(message)
  }
}

export const sendMessage = (message: string): void => {
  if (socket) {
    socket.send(JSON.stringify({ content: message }))
  }
}
