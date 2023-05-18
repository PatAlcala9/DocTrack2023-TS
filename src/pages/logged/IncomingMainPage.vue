<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Incoming Commmunications
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

  div.full-width.column.items-center.content-center
    doc-label(text="Entry Code Number" )
    doc-input(v-model:value="entryCode" @keypress.enter = "searchIncoming")

  section(v-if="dataExist")
    div(v-if="sourceName").full-width.column.justify-start.information
      section.column
        doc-label(text="Source Name:").information__label
        span.information__data {{ sourceName }}

    div(v-if="dateReceived").full-width.column.justify-start.information
      section.column
        doc-label(text="Date Received:").information__label
        span.information__data {{ dateReceived }}

    div(v-if="subject").full-width.column.justify-start.information
      section.column
        doc-label(text="Subject:").information__label
        span.information__data {{ subject }}

    div(v-if="details").full-width.column.justify-start.information
      section.column
        doc-label(text="Details:").information__label
        span.information__data {{ details }}


    div(v-if="attachments").full-width.column.justify-start.information
      section.column
        doc-label(text="Attachments:").information__label
        span.information__data {{ attachments }}

    div(v-else).full-width.column.justify-start.information
      section.column
        doc-button(text="Add Attachment").information__label


    div(v-if="notes").full-width.column.justify-start.information
      section.row
        doc-label(text="Notes:").information__label
        span.information__data {{ notes }}

    div(v-else).full-width.column.justify-start.information
      section.column
        doc-button(text="Add Notes").information__label

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
import docLabel from 'components/docLabel.vue'

let entryCode = ref('')
let dataExist = ref(false)

let sourceName = ref('')
let dateReceived = ref('')
let subject = ref('')
let details = ref('')
let attachments = ref('')
let notes = ref('')
let docLogsList = []
let actionMadeList = []

const router = useRouter()
let _currentpage = useCurrentPage()

const searchIncoming = async () => {
  try {
    const response = await api.get('/api/SearchIncoming/' + entryCode.value)
    const data = response.data

    if (data.result.length > 0) {
      const dateReceivied = data.result

      dataExist.value = true
      dateReceived.value = date.formatDate(dateReceivied, 'MMMM D, YYYY')
      sourceName.value = data.result2
      subject.value = data.result3
      details.value = data.result4
      attachments.value = data.result5
      notes.value = data.result6
    } else dataExist.value = false
  } catch {
    dataExist.value = false
  }
}

const searchIncomingAction = async () => {
  try {
    const response = await api.get('/api/SearchIncomingAction/' + entryCode.value)
    const data = response.data

    if (data.result.length > 0) {
      // docLogsList
    } else dataExist.value = false
  } catch {}
}

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

const showData = () => {
  dataExist.value = true
}

;(async () => {
  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/incoming')
})()
</script>

<style lang="sass" scoped>

.calendar
  background-color: rgba(8,25,35, 0.75)
  border: 1px solid rgba(255, 255, 255, 0.525)
  font-family: 'OpenSans'
  font-size: 1.6rem
  border-radius: 12px
  :deep .q-date__calendar-item
    padding-top: 0.6rem
    // padding-left: 0.5rem
    // padding-right: 0.5rem
  :deep .block
    font-size: 1.25rem


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

.information
  padding: 1rem 0 0 1rem

.information__label
  margin-right: 0.6rem

.information__data
  font-family: 'OpenSans'
  font-size: 1.2rem
  border: 1px solid white
  padding: 0.3rem 0.5rem
  // margin-top: -0.5rem

.spaced
</style>
