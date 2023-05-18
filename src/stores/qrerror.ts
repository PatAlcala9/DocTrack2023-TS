import { defineStore } from 'pinia'

export const useQrError = defineStore('qrerror', {
  state: () => ({
    qrerror: '',
  }),
})
