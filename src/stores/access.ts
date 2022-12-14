import { defineStore } from 'pinia'

interface IUserState {
  access: string[]
}

export const useAccess = defineStore('access', {
  state: (): IUserState => ({
    access: [],
  }),
})
