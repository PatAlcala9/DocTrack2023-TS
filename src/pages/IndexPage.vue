<template lang="pug">

q-page(padding).full-width.column.items-start.content-center
  img(src="../assets/ocbologo.svg" alt="OCBO Logo").logo

  div.fit.column.text-center
    span.name {{_employeename.employeename}}
    span.menu Menu

  div.button-area.fit.row.justify-evenly.content-start
    q-btn(outline text-color="white" color="white" to="incoming").button-access Incoming
    q-btn(outline text-color="white" color="white" ).button-access Outgoing
    q-btn(outline text-color="white" color="white" ).button-access Releasing
    q-btn(outline text-color="white" color="white" ).button-access Inventory
    q-btn(outline text-color="white" color="white" ).button-access Other Docs

  </template>

<script setup lang="ts">
import { ref } from 'vue'
import { SessionStorage } from 'quasar'

import { useEmployeeName } from 'stores/employeename'

let _employeename = useEmployeeName()
let employeenameLocal = ref('')

const setName = async () => {
  if (_employeename.employeename !== undefined) {
    SessionStorage.set('EmployeeName', _employeename.employeename)
  }
  // employeenameLocal.value = SessionStorage.getItem('EmployeeName')

  employeenameLocal.value = _employeename.employeename
}

;(async () => {
  await setName()
})()
</script>

<style lang="sass" scoped>
.logo
  width: 16rem
  height: auto
  position: fixed
  opacity: 0.05

.name
  font-size: 4rem
  font-family: 'RalewayBold'
  color: #ffffff

.menu
  margin-top: 1rem
  font-size: 2.2rem
  font-family: 'RalewayBold'
  color: #ffffff

.button-area
  margin-top: 4.2rem

.button-access
  font-family: 'RalewayBold'
  font-size: 1.4rem
  border-radius: 8px
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.18)
  border: none
  padding: 0.80rem 1.85rem
  min-width: 10ch
  min-height: 44px
  text-align: center
  line-height: 1.1
  cursor: pointer
  text-transform: uppercase

.button-access:hover
  background-color: #021926
  color: #ffffff
  // border: 1px solid #ffffff
</style>
