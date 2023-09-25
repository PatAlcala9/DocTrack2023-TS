import { defineStore } from 'pinia';

export const useEmployeeID = defineStore('employeeid', {
  state: () => ({
    employeeid: 0
  })
});
