<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Inquiry for Received Documents
    q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button

  div.fit.row.wrap.justify-start.items-start.content-start.search-area
    section.column
      span Search
      q-input(dense filled v-model="searchValue" bg-color="white" :input-style="{ fontSize: '1.2rem', fontFamily: 'OpenSans' }" @keydown.enter="getIncomingUsingValue")
    section.column
      span Search By
      //- input
      q-btn-dropdown(unelevated color="grey-10" :label="searchByValue" size="1rem" :content-style="{ fontSize: '1.2rem', fontFamily: 'Raleway' }")
        q-list
          q-item(clickable v-close-popup @click="searchByCode" @keypress="formatApplication($event)")
            q-item-section
              q-item-label ENTRY CODE

          q-item(clickable v-close-popup @click="searchByValue = 'DATE'")
            q-item-section
              q-item-label DATE

          q-item(clickable v-close-popup @click="searchByValue = 'NAME'")
            q-item-section
              q-item-label NAME

          q-item(clickable v-close-popup @click="searchByValue = 'SUBJECT'")
            q-item-section
              q-item-label SUBJECT

  div.flex.flex-center
    section(v-if="incomingList.result.length > 0").dialog-content-table
      table.table
        thead
          tr
            th Entry Code
            th Received Date
            th Name
            th Subject
            th Details
        tbody
          tr(v-for="(item, index) in incomingList.result" :key="item").table-content-group
            td {{item}}
            td {{incomingList.result2[index]}}
            td {{incomingList.result3[index]}}
            td {{incomingList.result4[index]}}
            td
              q-btn(color="button" icon="visibility" :ripple="false" @click="openDetails(item, incomingList.result2[index], incomingList.result3[index], incomingList.result4[index])").button-view
              //- q-btn(v-if="showText === false" v-else color="button" label="View" :ripple="false" @mouseleave="mouseLeave").button-view

    section(v-else).table-loading.column.items-center
      span Loading Contents
      q-spinner-orbit(color="white" size="4em" style="margin-top: 2rem")

q-dialog(v-model="details" maximized)
  q-card.detail-dialog
    q-card-section.full-width.column.items-center
      section.full-width.row.justify-between
        span.detail-dialog__info--large {{entryCodeDetail}}
        span.detail-dialog__info--large {{dateReceivedDetail}}
      span.detail-dialog__info {{nameDetail}}
      section.full-width.row
        span.detail-dialog__info Subject: {{subjectDetail}}
      section.full-width.row
        span.detail-dialog__info Details: {{detailsDetail}}

    q-card-section.full-width.column.items-center
      section.full-width.column.justify-around
        span(v-if="attachmentDetail !== null").detail-dialog__info--subinfo Attachments: {{attachmentDetail}}
        span(v-if="noteDetail !== null").detail-dialog__info--subinfo Notes: {{noteDetail}}

    q-card-section.full-width.column.items-center
      span.detail-dialog__info Document Logs
      table.table
        thead
          tr
            th Date
            th Forwarded To
        tbody
          tr(v-for="(item, index) in doclogDetail.result" :key="item").table-content-group
            td {{item}}
            td {{doclogDetail.result2[index]}}

    q-card-section.full-width.column.items-center
      span.detail-dialog__info Actions Logs
      table.table
        thead
          tr
            th Date
            th Actions Made
        tbody
          tr(v-for="(item, index) in actionlogDetail.result" :key="item").table-content-group
            td {{item}}
            td {{actionlogDetail.result2[index]}}

    q-card-actions(align="right")
      q-btn(flat label="OK" color="primary" v-close-popup)
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { api } from 'boot/axios'
import { useRouter } from 'vue-router'

const router = useRouter()

let searchValue = ref('')
let searchByValue = ref('')
let incomingList: any = ref({})
let showText = ref(false)

const getIncomingUsingValue = async () => {
  if (searchValue.value.length > 0) {
    switch (searchByValue.value) {
      case 'ENTRY CODE':
        getIncomingByEntryCode()
        break
      case 'NAME':
        getIncomingBySourceName()
        break
      case 'SUBJECT':
        getIncomingBySubject()
        break
    }
  } else getIncoming()
}

const getIncoming = async () => {
  const response = await api.get('/api/GetIncoming')
  const data = response.data

  if (data !== undefined) {
    incomingList.value = data
  }
}

const getIncomingByEntryCode = async () => {
  const response = await api.get('/api/GetIncomingByEntryCode/' + searchValue.value)
  const data = response.data

  if (data !== undefined) {
    incomingList.value = data
  }
}

const getIncomingBySourceName = async () => {
  const response = await api.get('/api/GetIncomingBySourceName/' + searchValue.value)
  const data = response.data

  if (data !== undefined) {
    incomingList.value = data
  }
}

const getIncomingBySubject = async () => {
  const response = await api.get('/api/GetIncomingBySubject/' + searchValue.value)
  const data = response.data

  if (data !== undefined) {
    incomingList.value = data
  }
}

const gotoHome = async () => {
  await router.push('/')
  location.reload()
}

const mouseHover = () => {
  showText.value = true
}

const mouseLeave = () => {
  showText.value = false
}

const searchByCode = async () => {
  searchByValue.value = 'ENTRY CODE'
}

const formatApplication = (evt: any) => {
  evt = evt ? evt : window.event
  var charCode = evt.which ? evt.which : evt.keyCode

  // applicationNo.value;

  if (searchValue.value !== null && searchValue.value.length === 2) {
    searchValue.value = searchValue.value + '-'
  }

  if (charCode > 31 && (charCode < 45 || charCode > 57) && charCode !== 46) {
    evt.preventDefault()
  } else {
    return true
  }
}

let details = ref(false)
let entryCodeDetail = ref('')
let dateReceivedDetail = ref('')
let nameDetail = ref('')
let subjectDetail = ref('')
let detailsDetail = ref('')
let attachmentDetail = ref('')
let noteDetail = ref('')

let doclogDetail: any = ref([])
let actionlogDetail: any = ref([])

const openDetails = async (entryCode: string, date: string, sourcename: string, subjectdetail: string) => {
  details.value = true
  entryCodeDetail.value = entryCode
  dateReceivedDetail.value = date
  nameDetail.value = sourcename
  subjectDetail.value = subjectdetail

  await getIncomingDetails()
  await getIncomingDocLog()
  await getIncomingActionLog()
}

const getIncomingDetails = async () => {
  try {
    const response = await api.get('/api/GetIncomingDetails/' + entryCodeDetail.value)
    const data = response.data
    if (data !== undefined) {
      detailsDetail.value = data.result
      attachmentDetail.value = data.result2.length > 1 ? data.result2 : null
      noteDetail.value = data.result3.length > 1 ? data.result3 : null
    }
  } catch {
    detailsDetail.value = ''
    attachmentDetail.value = ''
    noteDetail.value = ''
  }
}

const getIncomingDocLog = async () => {
  try {
    const response = await api.get('/api/GetIncomingDocLog/' + entryCodeDetail.value)
    const data = response.data
    if (data !== undefined) doclogDetail.value = data
  } catch {
    doclogDetail.value = []
  }
}

const getIncomingActionLog = async () => {
  try {
    const response = await api.get('/api/GetIncomingActionLog/' + entryCodeDetail.value)
    const data = response.data
    if (data !== undefined) actionlogDetail.value = data
  } catch {
    actionlogDetail.value = []
  }
}

;(async () => {
  await getIncoming()
  await searchByCode()
})()
</script>

<style lang="sass" scoped>


.table-loading
  font-size: 2rem
  margin-top: 3rem

.table
  font-family: 'OpenSans'
  font-size: 0.9rem
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
  padding: 1rem 4rem
  width: 20rem

.table tbody tr
  border-bottom: 1px solid #dddddd

.table tbody tr:nth-of-type(even)
  background-color: #1C4157

.table tbody tr:last-of-type
  border-bottom: 2px solid #000406

.button-view
  width: 4rem
  background-color: $button

.search-area
  padding: 2rem 0 0 0

.search-area section
  padding-right: 1.4rem
  font-size: 1.1rem

.search-area section input
  border-radius: 0.6rem
  text-align: center
  border: 0
  height: 2rem
  box-shadow: none

.detail-dialog
  font-family: 'Raleway'

.detail-dialog__info
  font-size: 2rem

.detail-dialog__info--large
  font-family: 'OpenSans'
  font-size: 3rem

.detail-dialog__info--subinfo
  font-size: 1.2rem
</style>
