import { defineStore } from 'pinia';

export const useDemoMode = defineStore('demomode', {
  state: () => ({
    demomode: false
  })
});
