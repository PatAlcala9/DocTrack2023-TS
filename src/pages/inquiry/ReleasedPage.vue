<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Inquiry for Released Documents
    q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button

  div.fit.row.wrap.justify-start.items-start.content-start.search-area
    section.column
      span Search
      q-input(dense filled v-model="searchValue" bg-color="white" :input-style="{ fontSize: '1.2rem', fontFamily: 'OpenSans' }" @keydown.enter="getOutgoingUsingValue")
    section.column
      span Search By
      //- input
      q-btn-dropdown(unelevated color="grey-10" :label="searchByValue" size="1rem" :content-style="{ fontSize: '1.2rem', fontFamily: 'Raleway' }")
        q-list
          q-item(clickable v-close-popup @click="searchByReference")
            q-item-section
              q-item-label REFERENCE NUMBER

          q-item(clickable v-close-popup @click="searchByValue = 'RESPONDENT'")
            q-item-section
              q-item-label RESPONDENT

          q-item(clickable v-close-popup @click="searchByValue = 'ADDRESS'")
            q-item-section
              q-item-label ADDRESS

          q-item(clickable v-close-popup @click="searchByValue = 'SUBJECT'")
            q-item-section
              q-item-label SUBJECT

          q-item(clickable v-close-popup @click="searchByValue = 'DATE'")
            q-item-section
              q-item-label DATE

  div.flex.flex-center
    //- section(v-if="outgoingList.result !== ''").dialog-content-table
    //-   table.table
    //-     thead
    //-       tr
    //-         th Reference Number
    //-         th Respondent
    //-         th Address
    //-         th Subject
    //-         th Date Released
    //-         th Details
    //-     tbody
    //-       tr(v-for="(item, index) in outgoingList.result" :key="item").table-content-group
    //-         td {{item}}
    //-         td {{outgoingList.result2[index]}}
    //-         td {{outgoingList.result3[index]}}
    //-         td {{outgoingList.result4[index]}}
    //-         td {{outgoingList.result5[index]}}
    //-         td
    //-           q-btn(color="button" icon="visibility" :ripple="false" @click="openDetails(item, outgoingList.result2[index], outgoingList.result3[index], outgoingList.result4[index])").button-view
    //-           //- q-btn(v-if="showText === false" v-else color="button" label="View" :ripple="false" @mouseleave="mouseLeave").button-view

    //- section(v-else).table-loading.column.items-center
    //-   span Loading Contents
    //-   q-spinner-orbit(color="white" size="4em" style="margin-top: 2rem")
    section.column.text-center(style="font-size: 1.2rem")
      span Cannot display table
      span No Connection on Server

q-dialog(v-model="details" maximized)
  q-card.detail-dialog
    q-card-section.full-width.column.items-center
      section.full-width.column.justify-between
        span.detail-dialog__info--large {{referenceDetail}}
        span.detail-dialog__info Respondent Name:
          a.detail-dialog__info--detail {{respondentDetail}}
        span.detail-dialog__info Address:
          a.detail-dialog__info--detail {{addressDetail}}
        br
        span(v-if="subjectDetail !== null").detail-dialog__info--subinfo Subject: {{subjectDetail}}
        span(v-if="detailsDetail !== null").detail-dialog__info--subinfo Details: {{detailsDetail}}
        span(v-if="attachmentDetail !== null").detail-dialog__info--subinfo Attachments: {{attachmentDetail}}

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

    q-card-actions(align="center")
      component(:is="docButton" text="Close" v-close-popup).detail-dialog__button
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { api } from 'boot/axios'
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'

import docButton from 'components/docButton.vue'

const router = useRouter()
const _currentpage = useCurrentPage()

let searchValue = ref('')
let searchByValue = ref('')

type Outgoing = {
  result: string
  result2: string
  result3: string
  result4: string
  result5: string
}
let outgoingList = ref({} as Outgoing)
let showText = ref(false)

const gotoHome = () => {
  _currentpage.currentpage = '/'
  router.push('/')
  // location.reload()
}

const searchByReference = async () => {
  searchByValue.value = 'REFERENCE NUMBER'
}

const getOutgoingUsingValue = async () => {
  if (searchValue.value.length > 0) {
    switch (searchByValue.value) {
      case 'REFERENCE NUMBER':
        getOutgoingByReference()
        break
      case 'RESPONDENT':
        console.log(searchByValue.value)
        getOutgoingByRespondent()
        break
      case 'ADDRESS':
        getOutgoingByAddress()
        break
      case 'SUBJECT':
        getOutgoingBySubject()
        break
      case 'DATE RELEASED':
        // getIncomingBySubject()
        break
    }
  } else getOutgoing()
}

const getOutgoing = async () => {
  const response = await api.get('/api/GetOutgoing')
  const data = response.data

  if (data !== undefined) {
    outgoingList.value = data
  }
}

const getOutgoingByReference = async () => {
  const response = await api.get('/api/GetOutgoingByReference/' + searchValue.value)
  const data = response.data

  if (data !== undefined) {
    outgoingList.value = data
  }
}

const getOutgoingByRespondent = async () => {
  const response = await api.get('/api/GetOutgoingByRespondent/' + searchValue.value)
  const data = response.data

  if (data !== undefined) {
    outgoingList.value = data
  }
}

const getOutgoingByAddress = async () => {
  const response = await api.get('/api/GetOutgoingByAddress/' + searchValue.value)
  const data = response.data

  if (data !== undefined) {
    outgoingList.value = data
  }
}

const getOutgoingBySubject = async () => {
  const response = await api.get('/api/GetOutgoingBySubject/' + searchValue.value)
  const data = response.data

  if (data !== undefined) {
    outgoingList.value = data
  }
}

let details = ref(false)
let referenceDetail = ref('')
let respondentDetail = ref('')
let addressDetail = ref('')
let fromDetail = ref('')
let subjectDetail = ref('')
let attachmentDetail = ref()
let detailsDetail = ref('')

let doclogDetail: any = ref([])
let actionlogDetail: any = ref([])

const openDetails = async (reference: string, respondent: string, address: string, subject: string) => {
  details.value = true
  referenceDetail.value = reference
  respondentDetail.value = respondent
  addressDetail.value = address
  subjectDetail.value = subject

  await getOutgoingDetails()
  await getOutgoingDocLog()
  await getOutgoingActionLog()
}

const getOutgoingDetails = async () => {
  try {
    const response = await api.get('/api/GetOutgoingDetails/' + referenceDetail.value.substring(8))
    const data = response.data

    if (data !== undefined) {
      fromDetail.value = data.result
      detailsDetail.value = data.result2.length > 1 ? data.result2 : null
      attachmentDetail.value = data.result3.length > 1 ? data.result3 : null
    }
  } catch {
    fromDetail.value = ''
    detailsDetail.value = ''
    attachmentDetail.value = ''
  }
}

const getOutgoingDocLog = async () => {
  try {
    const response = await api.get('/api/GetOutgoingDocLog/' + referenceDetail.value.substring(8))
    const data = response.data
    if (data !== undefined) doclogDetail.value = data
  } catch {
    doclogDetail.value = []
  }
}

const getOutgoingActionLog = async () => {
  try {
    const response = await api.get('/api/GetOutgoingActionLog/' + referenceDetail.value.substring(8))
    const data = response.data
    if (data !== undefined) actionlogDetail.value = data
  } catch {
    actionlogDetail.value = []
  }
}

;(async () => {
  await getOutgoing()
  await searchByReference()

  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/released')
})()


</script>

<style lang="sass" scoped>


.table-loading
  font-size: 2rem
  margin-top: 3rem

.table
  font-family: 'Montserrat'
  font-size: 0.8rem
  text-transform: uppercase
  border-collapse: collapse
  margin: 2rem 0
  min-width: 24rem
  border-radius: 1rem 1rem 0 0
  overflow: hidden
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2)

.table thead tr
  background-color: #000406
  color: #ffffff
  text-align: left
  font-weight: bold

.table th, .table td
  padding: 1rem 2.5rem
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
  font-family: 'Montserrat'
  background-color: $background
  color: #ffffff
  margin: 1rem
  border: 1px solid #2F5972
  border-radius: 12rem

.detail-dialog__info
  font-size: 1.4rem
  font-family: 'Montserrat'

.detail-dialog__info--large
  font-family: 'OpenSans'
  font-size: 2.2rem

.detail-dialog__info--subinfo
  font-size: 1.2rem

.detail-dialog__info--detail
  font-size: 1.6rem
  font-family: 'Montserrat'
  font-weight: bold
  margin-left: 1rem

.detail-dialog__button
  margin-bottom: 2rem
</style>
