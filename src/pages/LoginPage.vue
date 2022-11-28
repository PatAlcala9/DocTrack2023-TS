<template lang="pug">

q-page(padding)
  div.grid
    section.title.full-width.row.justify-between
      div.column
        img(src="../assets/ocbologo.svg" alt="OCBO Logo").logo
        section.name.fit.column.wrap.justify-start.items-start.content-start
          span.ocbo-title OCBO
          span.ocbo-text Doctrack System 2023

    div(v-if="inquiry === false").login
      section.username-area.column.wrap.justify-center.items-center.content-center
        component(:is="docLabel" text="Username")
        component(:is="docInput" v-model:value="usernameEntry")

      section.password-area.column.wrap.justify-center.items-center.content-center
        component(:is="docLabel" text="Password")
        component(:is="docInputPassword" v-model:value="passwordEntry" @keypress.enter="login")

      section.button-area.column.wrap.justify-center.items-center.content-center
        component(:is="docButton" text="login" @click="login")

      section.register.column.wrap.justify-center.items-center.content-center
        span(@click="gotoRegister") Create New Account

    div(v-else).login
      section.username-area.column.wrap.justify-center.items-center.content-center
        span.inquiry-label Inquiry
      section.username-area.column.wrap.justify-center.items-center.content-center
        component(:is="docButton" text="Received" @click="inquireReceivedTrigger")
      section.password-area.column.wrap.justify-center.items-center.content-center
        component(:is="docButton" text="Released" @click="inquireReleasedTrigger")

    div.inquiry
      span(v-if="inquiry" @click="inquiry = !inquiry").inquiry-text Login
      span(v-else @click="inquiry = !inquiry").inquiry-text Inquire without logging In

    div.davao
      img(src="../assets/davao.svg" alt="Davao Logo").davaologo

  q-dialog(v-model="inquireReceived" persistent full-width full-height transition-show="scale" transition-hide="scale")
    q-card.dialog-card.text-white
      q-card-section
        div.dialog-title-area.row.justify-between
          span.dialog-title Inquiry for Received Documents
          q-btn(flat size="md" label="close" v-close-popup).dialog-close

      q-card-section
        div.fit.row.wrap.justify-start.items-start.content-start
          section.column
            span.dialog-content-label Search
            input.dialog-content-input
          section.column
            span.dialog-content-label Date
            input.dialog-content-input

          section(v-if="incomingList.length > 0").dialog-content-table
            table.table
              thead
                tr.table-header-group
                  th.table-header Entry Code
                  th.table-header Received Date
                  th.table-header Source
                  th.table-header Subject
                  th.table-header Details
              tbody
                tr(v-for="data in incomingList" :key="data").table-content-group
                  td.table-content {{data.entryCodeNo}}
                  td.table-content {{data.receivedDate}}
                  td.table-content {{data.sourceName}}
                  td.table-content {{data.subjectInfo}}
                  td.table-content View
          section(v-else).dialog-content-table
            span Loading
            //- q-table(:rows="incomingList" :columns="incomingHeaderList" row-key="name" :table-header-style="{ backgroundColor: '#021926', color: '#ffffff', fontFamily: 'Raleway', fontSize: '12px' }" :table-style="{ backgroundColor: 'red' }")

</template>

<script setup lang="ts">
import { ref } from 'vue'
import { api } from 'boot/axios'
import { useRouter } from 'vue-router'
import { encrypt, comparePassword } from 'src/js/OCBO'

import { useEmployeeName } from 'stores/employeename'
import { useUserID } from 'stores/userid'

import docButton from 'components/docButton.vue'
import docInput from 'components/docInput.vue'
import docInputPassword from 'components/docInputPassword.vue'
import docLabel from 'components/docLabel.vue'

let _employeename = useEmployeeName()
let _userid = useUserID()

const router = useRouter()

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

const showLogin = () => {
  inquiry.value = !inquiry.value
}

const inquireReceivedTrigger = async () => {
  router.push('/received')
}

const inquireReleasedTrigger = async () => {
  router.push('/released')
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
    router.push('/dashboard')
  }
}

const gotoRegister = async () => {
  router.push('/register')
}

let usernameAccepted = false
const checkUsername = async () => {
  try {
    const response = await api.get('/api/CheckUsername/' + usernameEntry.value.toUpperCase())
    const data = response.data
    if (data !== undefined) usernameAccepted = data.result === '1' ? true : false
  } catch {
    usernameAccepted = false
  }
}

let passwordAccepted = false
const checkPassword = async () => {
  try {
    const response = await api.get('/api/CheckPassword/' + usernameEntry.value.toUpperCase())
    const data = response.data
    if (data !== undefined && comparePassword(data.result, passwordEntry.value.toUpperCase())) passwordAccepted = true
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
      detailsAllowed = true
    }
  } catch {
    detailsAllowed = false
  }
}

const login = async () => {
  await checkUsername()
  if (usernameAccepted === false) {
    console.log('INVALUD USERNAME')
    return
  }

  await checkPassword()
  if (passwordAccepted === false) {
    console.log('INVALID PASSWORD')
    return
  }

  await getUserDetails()
  if (detailsAllowed === false) {
    console.log('NO DETAILS')
    return
  }

  router.push('/dashboard')
}
</script>

<style lang="sass" scoped>
.cabinet
  width: 25rem
  height: auto
  opacity: 0.05

.ocbo-title
  font-size: 5rem
  font-family: 'RalewayBold'
  color: #ffffff

.ocbo-text
  font-size: 3.5rem
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
  height: 100vh

.title
  grid-area: title
  justify-self: start
  align-self: stretch
  padding: 0 0 0 1rem
  margin: 0 0 -2rem 0

.logo
  width: 12rem
  height: auto
  opacity: 0.9
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
  margin-top: -10rem
  // backdrop-filter: blur(16px) saturate(180%)

.inquiry
  grid-area: inquiry
  justify-self: start
  align-self: center

.inquiry-text
  font-family: 'Raleway'
  font-size: 1.2rem
  text-decoration: underline
  padding: 2rem
  cursor: pointer
  color: #ffffff

.username-area
  margin: 1rem

.username-label
  font-family: 'Raleway'
  font-size: 1.4rem
  color: #ffffff

.password-area
  margin: 1rem

.password-label
  @extend .username-label


.button-area
  padding-top: 1rem

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
  // justify-items: end

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
</style>
