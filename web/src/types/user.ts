export interface User {
  id: string
  name: string
  avatar: string
  email: string
}

export interface ShowUser {
  avatar: string
  name: string
}

export interface NewUser {
  username: string
  password: string
  email: string
}
