<template lang="pug">

div
  input(:value="props.value" @input="updateValue($event.target.value)").input

</template>

<script setup lang="ts">
export interface Props {
  value: string
}

const props = defineProps<Props>()
const emit = defineEmits(['update:value'])

const updateValue = (value: string) => {
  emit('update:value', value)
}
</script>

<style lang="sass" scoped>
.input
  font-family: 'Raleway'
  font-size: 1.3rem
  border-radius: 0.6rem
  text-align: center
  text-transform: uppercase
</style>
