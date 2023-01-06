import { defineStore } from 'pinia'

export const useCurrentPage = defineStore('currentpage', {
  state: () => ({
    currentpage: '',
  }),
})
