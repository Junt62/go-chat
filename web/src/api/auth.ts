import request from '@/utils/request'
import axios from 'axios'

interface RegisterData {
  username: string
  password: string
  email: string
}

interface LoginData {
  username: string
  password: string
}

interface LoginResponse {
  token: string
}

export const register = (data: RegisterData) => {
  return request.post('/register', data)
}

export const login = (data: LoginData) => {
  return request.post<LoginResponse>('/login', data)
}

export const logout = () => {
  return request.post('/logout')
}

export const ping = () => {
  return request.get('/ping')
}
