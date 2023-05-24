<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Incomings - New Entry
    q-btn(flat size="md" label="Back" @click="gotoIncomingDashboard" icon="arrow_back").close-button

  section.form-area.fit.row.items-center.justify-center
    div.column
      span.inputs__label Received Date:
        span.requiredWarning {{ missingDate }}
      q-date(flat v-model="receivedDate" minimal color="$button" @click="sample").calendar
      span(v-if="formattedReceivedDate.length > 0").inputs__label--date {{ formattedReceivedDate }}
      span(v-else).inputs__label--date No Date Selected

    div.inputs.column
      span.inputs__label Source:
        span.requiredWarning {{ missingSource }}
      component(:is="docInput" width="40" alignment="left" transform="initial" v-model:value="inSource")

      span.inputs__label Subject:
        span.requiredWarning {{ missingSubject }}
      component(:is="docInput" width="40" alignment="left" transform="initial" v-model:value="inSubject")

      span.inputs__label Details:
        span.requiredWarning {{ missingDetails }}
      component(:is="docTextArea" width="40" v-model:value="inDetails")

      span.inputs__label Attachments:
        span.requiredWarning {{ missingAttachments }}
      component(:is="docTextArea" width="40" v-model:value="inAttachments")

  section.button-area.fit.row.items-center.justify-center
    doc-button(text="Save" type="submit" @click="saveNewIncoming")

  //- section.list-area.column.text-center
  //-   span.list-area__span(v-if="!showList" @click="showListTrigger").inquiry-text Show List of Incomings
  //-   span.list-area__span(v-else @click="showListTrigger").inquiry-text Hide List of Incomings

  //- section(v-if="showList === true" style="margin-top: -3rem")
  //-   table.table
  //-     thead
  //-       tr
  //-         th Entry Code
  //-         th Received Date
  //-         th Source
  //-         th Subject
  //-         th Details
  //-     tbody
  //-       tr(v-for="(item, index) in incomingList.result" :key="item").table-content-group
  //-         td {{item}}
  //-         td {{incomingList.result2[index]}}
  //-         td {{incomingList.result3[index]}}
  //-         td {{incomingList.result4[index]}}
  //-         td

q-dialog(v-model="dialog" transition-show="flip-right" transition-hide="flip-left")
  q-card.dialog-card.text-white.flex.flex-center
    q-card-section.dialog-card__section
      div.dialog-title-area.column.justify-center.items-center
        span.dialog-card__title {{dialogMessage}}
        span.dialog-card__info {{dialogSubMessage}}
        component(:is="docButton" text="OK" v-close-popup)
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { date } from 'quasar'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'

import { useCurrentPage } from 'stores/currentpage'

import docButton from 'components/docButton.vue'
import docTextArea from 'components/docTextArea.vue'
import docInput from 'components/docInput.vue'

const router = useRouter()
let _currentpage = useCurrentPage()

let receivedDate = ref('')
let formattedReceivedDate = ref('')

let inSource = ref('')
let inSubject = ref('')
let inDetails = ref('')
let inAttachments = ref('')

let missingDate = ref('')
let missingSource = ref('')
let missingSubject = ref('')
let missingDetails = ref('')
let missingAttachments = ref('')

let dialog = ref(false)
let dialogMessage = ref('')
let dialogSubMessage = ref('')

let showList = ref(false)
type Incoming = {
  result: string
  result2: string
  result3: string
  result4: string
}
let incomingList = ref({} as Incoming)

const showListTrigger = async () => {
  showList.value = !showList.value

  if (showList.value === true) {
    await getIncomingDesc()
  }
}

const sample = () => {
  formattedReceivedDate.value = date.formatDate(Date.parse(receivedDate.value), 'MMMM D, YYYY')
}

const getIncomingDesc = async () => {
  const response = await api.get('/api/GetIncomingDesc')
  const data = response.data

  if (data !== undefined) {
    incomingList.value = data
  }
}

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

const gotoIncomingDashboard = () => {
  _currentpage.currentpage = 'incoming'
  router.push('/incoming')
}

let entryLast = ''
const getMaxEntryCode = async () => {
  const today = Date.now()
  const yy = date.formatDate(today, 'YY')

  try {
    const response = await api.get(`/api/GetMaxEntryCode/${yy}`)
    const data = response.data
    entryLast = data.result
    console.log(entryLast)
  } catch {
    entryLast = ''
  }
}

let newEntryCode = ''
const generateNewEntryCode = async () => {
  const today = Date.now()
  const yy = date.formatDate(today, 'YY')

  const series = parseInt(entryLast.substring(3))
  newEntryCode = yy + '-' + (series + 1).toString()
}

const saveIncoming = async () => {
  const incomingDate = date.formatDate(receivedDate.value, 'YYYY-MM-DD HH:mm:ss')
  let data: string
  try {
    const post = await api.post(`/api/SaveIncoming/${newEntryCode}/${incomingDate}/${inSource.value}/${inSubject.value}/${inDetails.value}/${inAttachments.value}`)
    data = post.data
  } catch {
    data = ''
  }

  if (data.includes('Success')) {
    dialog.value = true
    dialogMessage.value = 'Successfully Saved'
    dialogSubMessage.value = `Entry Code: ${newEntryCode}`
  } else {
    dialog.value = true
    dialogMessage.value = 'Failed on Saving Incoming'
    dialogSubMessage.value = ''
  }
}

let missingItems: any = ref([])
const checkData = async () => {
  if (receivedDate.value === '') {
    missingItems.value.push('ReceivedDate')
    missingDate.value = ' (**required**)'
  } else {
    for (let i in missingItems.value) {
      if (missingItems.value[i] === 'ReceivedDate') missingItems.value.splice(i, 1)
    }
    missingDate.value = ''
  }

  if (inSource.value === '') {
    missingItems.value.push('Source')
    missingSource.value = ' (**required**)'
  } else {
    for (let i in missingItems.value) {
      if (missingItems.value[i] === 'Source') missingItems.value.splice(i, 1)
    }
    missingSource.value = ''
  }

  if (inSubject.value === '') {
    missingItems.value.push('Subject')
    missingSubject.value = ' (**required**)'
  } else {
    for (let i in missingItems.value) {
      if (missingItems.value[i] === 'Subject') missingItems.value.splice(i, 1)
    }
    missingSubject.value = ''
  }

  if (inDetails.value === '') {
    missingItems.value.push('Details')
    missingDetails.value = ' (**required**)'
  } else {
    for (let i in missingItems.value) {
      if (missingItems.value[i] === 'Details') missingItems.value.splice(i, 1)
    }
    missingDetails.value = ''
  }

  if (inAttachments.value === '') {
    missingItems.value.push('Attachments')
    missingAttachments.value = ' (**required**)'
  } else {
    for (let i in missingItems.value) {
      if (missingItems.value[i] === 'Attachments') missingItems.value.splice(i, 1)
    }
    missingAttachments.value = ''
  }
}

const notifyMissingData = async () => {
  if (missingItems.value.length > 0) {
    dialog.value = true
    dialogMessage.value = 'Missing Data'
    dialogSubMessage.value = 'Please Input Missing Data'
  }
}

let connectionOK = false
const checkConnection = async () => {
  const response = await api.get('/api/CheckConnection')
  const data = response.data
  connectionOK = data.result === '1' ? true : false
}

const saveNewIncoming = async () => {
  await checkConnection()

  if (connectionOK === true) {
    await getMaxEntryCode()
    await checkData()
    await notifyMissingData()

    // if (missingItems.value.length === 0) {
    //   await generateNewEntryCode()
    //   await saveIncoming()
    // }
  }
}

;(async () => {
  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/incoming')
})()
</script>

<style lang="sass" scoped>

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

.username-input
  font-family: 'Raleway'
  font-size: 1.3rem
  border-radius: 0.6rem
  text-align: center
  text-transform: uppercase

.password-area
  margin: 1rem

.password-label
  @extend .username-label

.password-input
  @extend .username-input

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

.dialog-subtitle
  font-family: 'Raleway'
  font-size: 1.2rem

.inputs
  margin-left: 2rem

.inputs__label
  font-size: 1.2rem

.inputs__input
  font-family: 'OpenSans'
  font-size: 1.3rem
  border-radius: 0.6rem
  // text-align: center
  margin-bottom: 1rem
  width: 50vw

.inputs__label--date
  @extend .inputs__label
  font-family: 'OpenSans'
  margin-top: 0.5rem
  font-size: 1.4rem

.button-area
  margin-top: 2rem

.list-area
  padding-top: 2rem
  font-size: 0.5rem

.list-area__span
  cursor: pointer
  text-decoration: underline
  font-size: 1.2rem

.form-area
  margin-top: 2rem

.requiredWarning
  color: red

.table
  font-family: 'Montserrat'
  font-size: 0.8rem
  text-transform: uppercase
  border-collapse: collapse
  margin: 2rem 0
  min-width: 25rem
  border-radius: 1rem 1rem 0 0
  overflow: hidden
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2)

.table thead tr
  background-color: #000406
  color: #ffffff
  text-align: left
  font-weight: bold

.table th, .table td
  padding: 1rem 1rem
  width: 20rem

.table tbody tr
  border-bottom: 1px solid #dddddd

.table tbody tr:nth-of-type(even)
  background-color: #1C4157

.table tbody tr:last-of-type
  border-bottom: 2px solid #000406
</style>