<template lang="pug">

div
  q-btn(outline padding="0.2rem 2rem" rounded color="white" :label="text" size="lg" @click="$emit('click')").button

</template>

<script setup lang="ts">
export interface Props {
  text: string
}

const props = withDefaults(defineProps<Props>(), {
  text: 'button'
})
</script>

<style lang="sass" scoped>
.button
  font-family: 'Raleway',
</style>
