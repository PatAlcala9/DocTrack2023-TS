<template lang="pug">

q-page(padding)
  div.grid
    section.title.full-width.row.justify-between
      div.column
        transition(appear @before-enter="beforeEnterLogo" @enter="enterLogo")
          img(src="../assets/ocbologo.avif" alt="OCBO Logo" @click="test").logo
        transition(appear @before-enter="beforeEnterTitle" @enter="enterTitle")
          section.name.fit.column.wrap.justify-start.items-start.content-start
            span.ocbo-title OCBO
            span.ocbo-text Doctrack System 2023

    //- transition( appear @before-enter="beforeEnterForm" @enter="enterForm")
    div(v-if="inquiry === false").login#login
      div
        section.username-area.column.wrap.justify-center.items-center.content-center
          doc-label(text="Username")
          component(v-if="$q.screen.width <=500" :is="docInput" v-model:value="usernameEntry" width="10")
          doc-input(v-else v-model:value="usernameEntry")

        section.password-area.column.wrap.justify-center.items-center.content-center
          doc-label(text="Password")
          component(v-if="$q.screen.width <=500" :is="docInputPassword" v-model:value="passwordEntry" @keypress.enter="login" width="10")
          doc-input-password(v-else v-model:value="passwordEntry" @keypress.enter="login")

        section.button-area.column.wrap.justify-center.items-center.content-center
          doc-button(text="login" @click="login")

        section.register.column.wrap.justify-center.items-center.content-center
          span(@click="gotoRegister") Create New Account

    div(v-else).login#inquiry
      div
        section.username-area.column.wrap.justify-center.items-center.content-center
          span.inquiry-label Inquiry
        section.button-area--inquiry.column.wrap.justify-center.items-center.content-center
          doc-button(text="Received" @click="inquireReceivedTrigger")
        section.button-area--inquiry.column.wrap.justify-center.items-center.content-center
          doc-button(text="Released" @click="inquireReleasedTrigger")
        section.button-area--inquiry.column.wrap.justify-center.items-center.content-center
          doc-button(text="Scan using QR" @click="qrScannerTrigger")

    transition(appear @before-enter="beforeEnterSwitch" @enter="enterSwitch")
      div.inquiry
        span(v-if="inquiry" @click="showInquiry").inquiry-text Login
        span(v-else @click="showLogin").inquiry-text Inquire without logging In

  //- q-dialog(v-model="inquireReceived" persistent full-width full-height transition-show="scale" transition-hide="scale").dialog#dialog
  //-   q-card.dialog-card.text-white
  //-     q-card-section
  //-       div.dialog-title-area.row.justify-between
  //-         span.dialog-title Inquiry for Received Documents
  //-         q-btn(flat size="md" label="close" v-close-popup).dialog-close

  //-     q-card-section
  //-       div.fit.row.wrap.justify-start.items-start.content-start
  //-         section.column
  //-           span.dialog-content-label Search
  //-           input.dialog-content-input
  //-         section.column
  //-           span.dialog-content-label Date
  //-           input.dialog-content-input

  //-         section(v-if="incomingList.length > 0").dialog-content-table
  //-           table.table
  //-             thead
  //-               tr.table-header-group
  //-                 th.table-header Entry Code
  //-                 th.table-header Received Date
  //-                 th.table-header Source
  //-                 th.table-header Subject
  //-                 th.table-header Details
  //-             tbody
  //-               tr(v-for="data in incomingList" :key="data").table-content-group
  //-                 td.table-content {{data.entryCodeNo}}
  //-                 td.table-content {{data.receivedDate}}
  //-                 td.table-content {{data.sourceName}}
  //-                 td.table-content {{data.subjectInfo}}
  //-                 td.table-content View
  //-         section(v-else).dialog-content-table
  //-           span Loading
  //-           //- q-table(:rows="incomingList" :columns="incomingHeaderList" row-key="name" :table-header-style="{ backgroundColor: '#021926', color: '#ffffff', fontFamily: 'Raleway', fontSize: '12px' }" :table-style="{ backgroundColor: 'red' }")

q-dialog(v-model="error" transition-show="flip-right" transition-hide="flip-left" @keypress.enter="error=false").dialog#dialog
  q-card.dialog-card.text-white.flex.flex-center#dialog
    q-card-section.dialog-card__section
      div.dialog-title-area.column.justify-center.items-center
        span.dialog-card__title {{errorMessage}}
        span.dialog-card__info {{errorInformation}}
        doc-button(text="OK" @click="error=false")
</template>

<!-- <script lang="ts">
export default {
  preFetch({ redirect }) {
    if (_currentpage.currentpage !== 'register') {
      redirect({ path: '/login' })
    }
  },
}
</script> -->

<script setup lang="ts">
import { ref } from 'vue'
import { api } from 'boot/axios'
import { useRouter } from 'vue-router'
import { comparePassword } from 'src/js/OCBO'
import { gsap } from 'gsap'
import { SessionStorage, Platform } from 'quasar'

import { useEmployeeName } from 'stores/employeename'
import { useUserID } from 'stores/userid'
import { usePageWithTable } from 'stores/pagewithtable'
import { useAccess } from 'stores/access'
import { useCurrentPage } from 'stores/currentpage'

import docButton from 'components/docButton.vue'
import docInput from 'components/docInput.vue'
import docInputPassword from 'components/docInputPassword.vue'
import docLabel from 'components/docLabel.vue'

let _employeename = useEmployeeName()
let _userid = useUserID()
let _pagewithtable = usePageWithTable()
let _access = useAccess()
let _currentpage = useCurrentPage()

let error = ref(false)
let errorMessage = ref('')
let errorInformation = ref('')

const router = useRouter()

const run = async () => {
  const wasmModule = await WebAssembly.instantiateStreaming(fetch('src/js/ocbo.wasm'))
  const encrypt = wasmModule.instance.exports.encrypt as CallableFunction
  const encryptedMessage = await encrypt('hello', 'hello', 1, 128)
  return encryptedMessage
}

const test = async () => {
  const result = await run()
  console.log(result)
}

let inquiry = ref(false)
let inquireReceived = ref(false)
let inquireReleased = ref(false)

const incomingList: any = ref([])
let incomingHeaderList = [
  { name: 'entryCodeNo', align: 'center', label: 'Entry Code', field: 'entryCodeNo', sortable: true },
  { name: 'receivedDate', label: 'Received Date', field: 'receivedDate', sortable: true },
  { name: 'sourceName', label: 'Source', field: 'sourceName' },
  { name: 'subjectInfo', label: 'Subject', field: 'subjectInfo' },
  { name: 'view', label: 'Details', field: 'view' },
]

let usernameEntry = ref('')
let passwordEntry = ref('')
let userid = 0
let employeeName = null
let loginSuccess = false

const showInquiry = async () => {
  await exitInquiry()
  await enterLogin()
}

const showLogin = async () => {
  await exitLogin()
  await enterInquiry()
}

const inquireReceivedTrigger = async () => {
  _pagewithtable.pagewithtable = true
  SessionStorage.set('CurrentPage', 'received')
  _currentpage.currentpage = 'received'
  router.push('/received')
}

const inquireReleasedTrigger = async () => {
  _pagewithtable.pagewithtable = true
  SessionStorage.set('CurrentPage', 'released')
  _currentpage.currentpage = 'released'
  router.push('/released')
}

const qrScannerTrigger = async () => {
  _pagewithtable.pagewithtable = false
  SessionStorage.set('CurrentPage', 'qrscanner')
  _currentpage.currentpage = 'qrscanner'
  router.push('/qrscanner')
}

const getIncoming = async () => {
  const response = await api.get('/api/GetIncoming')
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    incomingList.value = data
  }
}

const tempLogin = async () => {
  const username = 'test'
  const password = 'doctrack'

  if (usernameEntry.value === username && passwordEntry.value === password) {
    _employeename.employeename = 'JUAN DELA CRUZ'
    SessionStorage.set('CurrentPage', 'dashboard')
    _currentpage.currentpage = 'dashboard'
    router.push('/dashboard')
  }
}

const gotoRegister = async () => {
  _pagewithtable.pagewithtable = false
  SessionStorage.set('CurrentPage', 'register')
  _currentpage.currentpage = 'register'
  router.push('/register')
}

let usernameAccepted = false
const checkUsername = async () => {
  try {
    const response = await api.get('/api/CheckUsername/' + usernameEntry.value.toUpperCase())
    const data = response.data
    const dataNum = parseInt(data.result)
    console.log(data)
    if (data !== undefined) usernameAccepted = dataNum > 0 ? true : false
  } catch {
    usernameAccepted = false
  }
}

let passwordAccepted = false
const checkPassword = async () => {
  try {
    const response = await api.get('/api/CheckPassword/' + usernameEntry.value.toUpperCase())
    const data = response.data
    if (data !== undefined && comparePassword(data.result, passwordEntry.value.toUpperCase(), 'doctrack2023', 3, 128)) passwordAccepted = true
  } catch {
    passwordAccepted = false
  }
}

let detailsAllowed = false
const getUserDetails = async () => {
  try {
    const response = await api.get('/api/GetUserDetails/' + usernameEntry.value.toUpperCase())
    const data = response.data

    if (data !== undefined) {
      _userid.userid = data.result
      _employeename.employeename = data.result2

      if (data.result3 === '1') _access.access.push('incoming')
      if (data.result4 === '1') _access.access.push('outgoing')
      if (data.result5 === '1') _access.access.push('releasing')
      if (data.result6 === '1') _access.access.push('inventory')
      if (data.result7 === '1') _access.access.push('otherdocuments')
      if (data.result8 === '1') _access.access.push('complaint')
      detailsAllowed = true
    }
  } catch {
    detailsAllowed = false
  }
}

const login = async () => {
  await checkUsername()
  if (usernameAccepted === false) {
    error.value = true
    errorMessage.value = 'Invalid Username'

    if (usernameEntry.value.length > 0) errorInformation.value = `${usernameEntry.value.toUpperCase()} does not exist`
    else errorInformation.value = 'Username is Empty'

    return
  }

  await checkPassword()
  if (passwordAccepted === false) {
    error.value = true
    errorMessage.value = 'Invalid Password'
    errorInformation.value = `Password does not match with ${usernameEntry.value.toUpperCase()}`

    if (passwordEntry.value.length > 0) errorInformation.value = `Password does not match with ${usernameEntry.value.toUpperCase()}`
    else errorInformation.value = 'Password is Empty'
    return
  }

  await getUserDetails()
  if (detailsAllowed === false) {
    error.value = true
    errorMessage.value = 'No Details Found'
    return
  }

  _pagewithtable.pagewithtable = false
  SessionStorage.set('CurrentPage', 'dashboard')
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

const exitLogin = async () => {
  await gsap.to('#login', { duration: 0.25, rotationY: 90, opacity: 0 })
  inquiry.value = true
}

const enterInquiry = async () => {
  gsap.from('#inquiry', { duration: 0.25, rotationY: 90, opacity: 0 })
}

const exitInquiry = async () => {
  await gsap.to('#inquiry', { duration: 0.25, rotationY: 90, opacity: 0 })
  inquiry.value = false
}

const enterLogin = async () => {
  gsap.from('#login', { duration: 0.25, rotationY: 90, opacity: 0 })
}

const beforeEnterLogo = (el: any) => {
  el.style.scale = 0
}
const enterLogo = (el: any) => {
  gsap.to(el, { duration: 1, scale: 1, ease: 'elastic.out(1, 0.5)' })
}

const beforeEnterTitle = (el: any) => {
  el.style.transform = 'translateX(100px)'
  el.style.opacity = 0
}
const enterTitle = (el: any) => {
  gsap.to(el, { duration: 0.8, x: 1, opacity: 1 })
}

const beforeEnterForm = (el: any) => {
  el.style.scale = 0
}
const enterForm = (el: any) => {
  gsap.to(el, { duration: 0.6, scale: 1, ease: 'back.out(1.6)' })
}

const beforeEnterSwitch = (el: any) => {
  el.style.transform = 'translateX(-100px)'
  el.style.opacity = 0
}
const enterSwitch = (el: any) => {
  gsap.to(el, { duration: 0.8, x: 1, opacity: 1 })
}

let browserName = ''
let browserVersion = 0
const detectBrowser = () => {
  browserName = Platform.is.name
  browserVersion = Platform.is.version ?? Platform.is.version : 0
}

;(async () => {
  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/')
})()
</script>

<style lang="sass" scoped>
// .cabinet
//   width: 25rem
//   height: auto
//   opacity: 0.05

.ocbo-title
  font-size: 4rem
  font-family: 'RalewayBold'
  color: #ffffff

.ocbo-text
  font-size: 3.3rem
  font-family: 'Montserrat'
  margin-top: -2rem
  color: #ffffff

.grid
  display: grid
  grid-template-columns: 1fr 1fr 1fr
  grid-template-rows: 1fr 1fr 1fr
  gap: 0px 0px
  grid-auto-flow: row
  justify-items: end
  grid-template-areas: "title title title" ". login ." "inquiry . davao"
  height: calc(100vh - 2rem)

.title
  grid-area: title
  justify-self: start
  align-self: stretch
  padding: 0 0 10rem 1rem
  margin: 0 0 2rem 0
  overflow: hidden

.logo
  width: 10rem
  height: auto
  opacity: 0.6
  margin-right: 2rem

.name
  grid-area: name
  justify-self: start
  align-self: center

.login
  grid-area: login
  justify-self: center
  align-self: center
  padding: 1.4rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  border-radius: 2rem
  background-color: #021926
  margin-top: -1rem
  width: 22rem
  height: 22rem

.inquiry
  grid-area: inquiry
  justify-self: left
  align-self: left

.inquiry-text
  font-family: 'Raleway'
  font-size: 1.2rem
  text-decoration: underline
  padding: 2rem
  cursor: pointer
  color: #ffffff

.username-area
  margin: 1rem

.password-area
  margin: 1rem

.button-area
  padding-top: 1rem

.button-area--inquiry
  padding-top: 2rem

.davao
  grid-area: davao
  justify-self: end
  align-self: center

.davaologo
  width: 18rem
  height: auto
  opacity: 0.4

.inquiry-label
  font-family: 'Raleway'
  font-size: 2.1rem
  color: #ffffff

.dialog
  width: 100vw

.dialog-title-area
  padding: 0.1rem 0  1rem 0.1rem

.dialog-title
  font-family: 'Raleway'
  font-size: 1.4rem

.dialog-card
  background-color: #021926

.dialog-content
  display: grid
  grid-template-columns: 1fr 1fr 1fr
  grid-template-rows: 1fr 1fr
  gap: 0px 0px
  grid-template-areas: "search search search" "table table table"

.dialog-content-search
  font-family: 'Raleway'
  grid-area: search
  justify-self: stretch
  align-self: start
  height: 4em

.dialog-content-input
  font-family: 'Raleway'
  font-size: 1.3rem
  border-radius: 0.6rem
  text-align: center

.dialog-content-table
  // grid-area: table
  // justify-self: center
  // align-self: center
  margin-top: 2rem

.table
  width: 93vw

.table-header-group
  font-family: 'Montserrat'
  font-size: 0.9em

.table-header
  padding: 0.2rem
  background-color: #ffffff
  color: #021926
  text-transform: uppercase

.table-content-group
  font-family: 'Montserrat'
  font-size: 0.9rem

.table-content
  padding: 0.2rem
  color: #ffffff

.register
  padding-top: 1.2rem
  font-size: 1.2rem
  cursor: pointer

.register:hover
  text-decoration: underline


@media screen and (max-width: 900px)
  .grid
    grid-template-areas: "title title title" "login login login" "inquiry inquiry ."
  .logo
    width: 9rem

  .ocbo-title
    font-size: 3.2rem

  .ocbo-text
    font-size: 2.5rem
    margin-top: -1rem

  .register
    font-size: 1.1rem

  .inquiry-text
    font-size: 1.1rem

@media screen and (max-width: 500px)
  .grid
    grid-template-areas: "title title title" "login login login" "inquiry inquiry inquiry"
  .logo
    width: 5rem
    opacity: 0.6

  .ocbo-title
    font-size: 2.2rem

  .ocbo-text
    font-size: 1.2rem

  .login
    margin-top: -5rem
    width: 14rem
    height: 22rem

  .register
    font-size: 0.9rem

  .inquiry
    margin-top: 1.2rem
    justify-self: center

  .inquiry-text
    font-size: 1.1rem
</style>
