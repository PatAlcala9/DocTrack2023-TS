<template lang="pug">

q-layout(view="hHh lpR lFr")
  q-page-container.page
    img(src="../assets/cabinet.svg" alt="Cabinet").cabinet
    img(v-show="_pagewithtable.pagewithtable === false" src="../assets/davao.avif" alt="Davao Logo").davaologo
    router-view

</template>

<script setup lang="ts">
import { usePageWithTable } from 'stores/pagewithtable'

const _pagewithtable = usePageWithTable()
</script>

<style lang="sass" scoped>
.page
  background: #000c13
  background: linear-gradient(83deg,#000c13 0%, #305b74 80%)
  font-family: 'Raleway'
  color: #ffffff

.cabinet
  width: 20rem
  height: auto
  opacity: 0.05
  position: absolute
  margin-top: 2rem
  margin-right: 2rem
  right: 0

.davaologo
  position: absolute
  top: calc(100vh - 8rem)
  left: calc(100% - 20rem)
  width: 18rem
  height: auto
  opacity: 0.4
</style>
