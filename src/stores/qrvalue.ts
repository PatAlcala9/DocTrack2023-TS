import { defineStore } from 'pinia'

export const useQrValue = defineStore('qrvalue', {
  state: () => ({
    qrvalue: '',
  }),
})
