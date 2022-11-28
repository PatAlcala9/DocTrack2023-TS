import { defineStore } from 'pinia';

export const useUserID = defineStore('userid', {
  state: () => ({
    userid: 0
  }),
});
