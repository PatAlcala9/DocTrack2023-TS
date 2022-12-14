import { defineStore } from 'pinia'

export const useIsIncoming = defineStore('isincoming', {
  state: () => ({
    isincoming: 0,
  }),
})
