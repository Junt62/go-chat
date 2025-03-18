export interface RegisterData {
  username: string
  password: string
  email: string
}

export interface LoginData {
  username: string
  password: string
}

export interface LoginResponse {
  token: string
}
