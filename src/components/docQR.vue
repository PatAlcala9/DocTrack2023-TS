<template lang="pug">

div
  component(:is="vueQr" text="text" :size="200")

</template>

<script setup lang="ts">
import vueQr from 'vue-qr/src/packages/vue-qr.vue'

export interface Props {
  text: string
}

const props = defineProps<Props>()
</script>
