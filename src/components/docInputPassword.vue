<template lang="pug">

div
  input(:value="props.value" type="password" @keypress.enter="$emit('keypress.enter')" @input="updateValue($event.target.value)").input

</template>

<script setup lang="ts">
export interface Props {
  value: string,
}

const props = defineProps<Props>()
const emit = defineEmits(['update:value'])

const updateValue = (value: any) => {
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
  border: 1px solid $button
  color: #000000
  width: 16rem

.input:focus
  outline: none
  border: 1px solid #ffffff
  background-color: #274c62
  color: #ffffff
</style>
