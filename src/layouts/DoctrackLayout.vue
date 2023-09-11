<template lang="pug">

q-layout(view="hHh lpR lFr")
  q-page-container.page
    img(v-show="_pagewithtable.pagewithtable === false" src="../assets/davao2.webp" alt="Davao Logo").davaologo
    router-view

</template>

<script setup lang="ts">
import { usePageWithTable } from 'stores/pagewithtable'

const _pagewithtable = usePageWithTable()
</script>

<style lang="sass" scoped>
.page
  // background: #000c13
  // background: linear-gradient(83deg,#000c13 0%, #305b74 80%)
  background-image: url('../assets/background4.webp')
  background-size: cover
  font-family: 'Inter'
  color: #ffffff
  background-attachment: fixed
  overflow-y: hidden

// .cabinet
//   width: 20rem
//   height: auto
//   opacity: 0.05
//   position: absolute
//   margin-top: 2rem
//   margin-right: 2rem
//   right: 0

.davaologo
  position: fixed
  top: calc(100vh - 5rem)
  left: calc(100% - 20rem)
  width: 18rem
  height: auto
  opacity: 0.5

@media screen and (max-width: 500px)
  .davaologo
    left: 50%
    bottom: 0
    transform: translateX(-50%)
</style>
