<template lang="pug">

div(v-if="mode===1")
  section.online-main
    q-icon(name="circle" size="1rem" :style="{ color: color }")
    span(:style="{ color: color }").text {{ props.state }}

div(v-if="mode===2")
  section.online
    q-icon(name="circle" size="1rem" :style="{ color: color }")
    span(:style="{ color: color }").text {{ props.state }}

</template>

<script setup lang="ts">
import { ref } from 'vue'

export interface Props {
  state: string
  mode: number
}

const props = withDefaults(defineProps<Props>(), {
  state: 'OFFLINE',
  mode: 1,
})

let color = ref('')

if (props.state === 'OFFLINE') color.value = '#E21D38'
else color.value = '#279D21'
</script>

<style lang="sass" scoped>
.online-main
  position: fixed
  top: 0
  left: 0
  padding: 1.6rem

.online
  position: fixed
  top: 0
  left: 90%
  padding-top: 1.5rem

.text
  font-family: 'Inter'
  font-weight: 800
  font-size: 0.9rem
  padding-left: 0.3rem
</style>
