import { defineStore } from 'pinia';

export const useIsReleasing = defineStore('isreleasing', {
  state: () => ({
    isreleasing: 0
  }),
});
