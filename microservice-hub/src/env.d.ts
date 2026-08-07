/// <reference types="vite/client" />

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