import { defineStore } from 'pinia'

export const useIsLogged = defineStore('islogged', {
  state: () => ({
    islogged: false,
  }),
})
