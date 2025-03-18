/// <reference types="vite/client" />
interface ImportMetaEnv {
  VITE_SERVER_ADDRESS: string
}
interface ImportMeta {
  readonly env: ImportMetaEnv
}
