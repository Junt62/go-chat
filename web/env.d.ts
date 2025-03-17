/// <reference types="vite/client" />
interface ImportMetaEnv {
  VITE_SERVER: string
}
interface ImportMeta {
  readonly env: ImportMetaEnv
}
