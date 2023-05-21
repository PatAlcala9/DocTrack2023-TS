<template lang="pug">

div
  component(:is="vueQr" :text="text" :size="size")

</template>

<script setup lang="ts">
import vueQr from 'vue-qr/src/packages/vue-qr.vue'

export interface Props {
  text: string,
  size: number
}

const props = defineProps<Props>()
</script>
