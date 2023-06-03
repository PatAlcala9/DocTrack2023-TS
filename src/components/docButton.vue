<template lang="pug">

div(v-if="$q.screen.width > 500")
  q-btn(outline padding="0.2rem 2rem" rounded color="white" :label="text" size="md" @click="('click')" ).button

div(v-else)
  q-btn(outline padding="0.2rem 2rem" rounded color="white" :label="text" size="lg" @click="('click')" ).button

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
  font-family: 'Inter'
  font-weight: 100 200
</style>
