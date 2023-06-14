<template lang="pug">

div
  component(:is="vueQr" :text="props.text" :size="props.size" :id="props.id")

</template>

<script setup lang="ts">
import vueQr from 'vue-qr/src/packages/vue-qr.vue'

export interface Props {
  text: string,
  size: number
  id: string
}

const props = defineProps<Props>()
</script>
