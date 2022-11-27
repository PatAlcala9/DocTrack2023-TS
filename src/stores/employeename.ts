import { defineStore } from 'pinia'

export const useEmployeeName = defineStore('employeename', {
  state: () => ({
    employeename: '',
  }),
})
