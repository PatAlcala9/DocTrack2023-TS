<template lang="pug">

q-page(padding).full-width.column.items-start.content-center
  img(src="../assets/ocbologo.svg" alt="OCBO Logo").logo

  div.fit.column.text-center
    transition(appear @before-enter="beforeEnterTitle" @enter="enterTitle")
      span.name {{_employeename.employeename}}
    transition(appear @before-enter="beforeEnterMenu" @enter="enterMenu")
      span.menu Menu
    transition(appear @before-enter="beforeEnterLogout" @enter="enterLogout")
      q-btn(flat size="md" label="Logout" @click="logout" icon="logout" ).logout-button

  div.button-area.fit.row.justify-evenly.content-start
    transition(appear @before-enter="beforeEnterButton" @enter="enterButton")
      component(:is="docButton" text="Incoming" @click="gotoPage('incoming')")
    component(:is="docButton" text="Outgoing")
    component(:is="docButton" text="Releasing")
    component(:is="docButton" text="Inventory")
    component(:is="docButton" text="Other Docs")

  </template>

<script setup lang="ts">
import { ref } from 'vue'
import { useQuasar } from 'quasar'
import { gsap } from 'gsap'
import { useRouter } from 'vue-router'

import docButton from 'components/docButton.vue'

import { useEmployeeName } from 'stores/employeename'

const router = useRouter()
const quasar = useQuasar()
let _employeename = useEmployeeName()

const setName = async () => {
  if (_employeename.employeename !== '') quasar.sessionStorage.set('EmployeeName', _employeename.employeename)
}

const getName = async () => {
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const session: string = quasar.sessionStorage.getItem('EmployeeName')!
  if (session !== '') _employeename.employeename = session
}

const beforeEnterTitle = (el: any) => {
  el.style.transform = 'translateY(-50px)'
  el.style.opacity = 0
}
const enterTitle = (el: any) => {
  gsap.to(el, { duration: 0.8, y: 0, opacity: 1 })
}

const beforeEnterMenu = (el: any) => {
  el.style.transform = 'translateY(-50px)'
  el.style.opacity = 0
}
const enterMenu = (el: any) => {
  gsap.to(el, { delay: 0.3, duration: 0.8, y: 0, opacity: 1 })
}

const beforeEnterLogout = (el: any) => {
  el.style.transform = 'translateX(30px)'
  el.style.opacity = 0
}
const enterLogout = (el: any) => {
  gsap.to(el, { delay: 0.1, duration: 0.8, x: 0, opacity: 1 })
}

const beforeEnterButton = (el: any) => {
  el.style.transform = 'translateX(30px)'
  el.style.opacity = 0
}
const enterButton = (el: any) => {
  gsap.to(el, { delay: 0.1, duration: 0.8, x: 0, opacity: 1 })
}

const logout = async () => {
  quasar.sessionStorage.remove('EmployeeName')
  gotoPage('/')
}

const gotoPage = (page: string) => {
  router.push(page)
}

;(async () => {
  await setName()
  await getName()
})()
</script>

<style lang="sass" scoped>
.logo
  width: 16rem
  height: auto
  position: fixed
  opacity: 0.05

.name
  font-size: 3.4rem
  font-family: 'RalewayBold'
  color: #ffffff

.menu
  margin-top: 1rem
  font-size: 1.8rem
  font-family: 'RalewayBold'
  color: #ffffff

.button-area
  margin-top: 4.2rem

.logout-button
  font-family: 'Montserrat'
  font-weight: bold
  padding: 1.6rem
  position: absolute
  top: 0
  right: 0
</style>
