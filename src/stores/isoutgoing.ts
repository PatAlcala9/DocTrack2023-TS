import { defineStore } from 'pinia';

export const useIsOutgoing = defineStore('isoutgoing', {
  state: () => ({
    isoutgoing: 0
  }),
});
