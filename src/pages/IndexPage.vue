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
