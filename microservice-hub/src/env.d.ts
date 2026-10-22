/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_USER_SERVICE_URL: string
  readonly VITE_ORDER_SERVICE_URL: string
  readonly VITE_INVENTORY_SERVICE_URL: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/concrete-component-props
  const component: DefineComponent<Record<string,unknown>, Record<string,unknown>, unknown>
  export default component
}
declare module 'vue-cal' {
  import type { DefineComponent } from 'vue'

  const VueCal: DefineComponent<{}, {}, any>
  export default VueCal
}