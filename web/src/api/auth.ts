import request from '@/utils/request'
import axios from 'axios'
import type { LoginData, LoginResponse, RegisterData } from './types'

export const apiRegister = (data: RegisterData) => {
  return request.post('/api/register', data)
}

export const apiLogin = (data: LoginData) => {
  return request.post<LoginResponse>('/api/login', data)
}

export const apiLogout = () => {
  return request.post('/api/logout')
}

export const apiPing = () => {
  return request.get('/api/ping')
}
