<template lang="pug">

div
  q-btn(outline padding="0.4rem 3rem" rounded color="white" :label="text" size="lg" @click="$emit('click')")

</template>

<script setup lang="ts">
export interface Props {
  text: string
}

const props = withDefaults(defineProps<Props>(), {
  text: 'button'
})
</script>

<style lang="sass" scoped></style>
