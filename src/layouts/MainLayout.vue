<template lang="pug">

q-layout(view="hHh lpR lFr").page
  img(src="../assets/cabinet.svg" alt="Cabinet").cabinet
  q-page-container
    router-view

  </template>

<script setup lang="ts"></script>

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
</style>
