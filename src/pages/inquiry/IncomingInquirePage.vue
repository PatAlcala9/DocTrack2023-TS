<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Inquiry for Received Documents
    q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button

  div.fit.row.wrap.justify-center.search-area
    component(:is="docForm" text="Search" v-model:value="searchValue" :width=24 :mobileWidth=14 icon="search" @keypress.enter="filterTable")

  //- div.fit.row.wrap.justify-start.items-start.content-start.search-area
  //-   section.column
  //-     span Search
  //-     q-input(dense filled v-model="searchValue" bg-color="white" :input-style="{ fontSize: '1.2rem', fontFamily: 'OpenSans' }" @keydown.enter="getIncomingUsingValue")
  //-   section.column
  //-     span Search By
  //-     //- input
  //-     q-btn-dropdown(unelevated color="grey-10" :label="searchByValue" size="1rem" :content-style="{ fontSize: '1.2rem', fontFamily: 'Raleway' }")
  //-       q-list
  //-         q-item(clickable v-close-popup @click="searchByCode" @keypress="formatApplication($event)")
  //-           q-item-section
  //-             q-item-label ENTRY CODE

  //-         q-item(clickable v-close-popup @click="searchByValue = 'DATE'")
  //-           q-item-section
  //-             q-item-label DATE

  //-         q-item(clickable v-close-popup @click="searchByValue = 'NAME'")
  //-           q-item-section
  //-             q-item-label NAME

  //-         q-item(clickable v-close-popup @click="searchByValue = 'SUBJECT'")
  //-           q-item-section
  //-             q-item-label SUBJECT

  //- div.flex.flex-center
  //-   //- section(v-if="incomingList.result !== ''").dialog-content-table
  //-   //-   table.table
  //-   //-     thead
  //-   //-       tr
  //-   //-         th Entry Code
  //-   //-         th Received Date
  //-   //-         th Name
  //-   //-         th Subject
  //-   //-         th Details
  //-   //-     tbody
  //-   //-       tr(v-for="(item, index) in incomingList.result" :key="item").table-content-group
  //-   //-         td {{item}}
  //-   //-         td {{incomingList.result2[index]}}
  //-   //-         td {{incomingList.result3[index]}}
  //-   //-         td {{incomingList.result4[index]}}
  //-   //-         td
  //-   //-           q-btn(color="button" icon="visibility" :ripple="false" @click="openDetails(item, incomingList.result2[index], incomingList.result3[index], incomingList.result4[index])").button-view
  //-   //-           //- q-btn(v-if="showText === false" v-else color="button" label="View" :ripple="false" @mouseleave="mouseLeave").button-view

  //-   //- section(v-else).table-loading.column.items-center
  //-   //-   span Loading Contents
  //-   //-   q-spinner-orbit(color="white" size="4em" style="margin-top: 2rem")
  //-   section.column.text-center(style="font-size: 1.2rem")
  //-     span Cannot display table
  //-     span No Connection on Server

  div(v-if="nodata").fit.row.wrap.justify-center
    span.flex.flex-center.nodata--text No Data Found

  div(v-else).fit.row.wrap.justify-center
    section(v-if="quasar.screen.width <= 500").flex.flex-center
      section(v-if="incomingList.result !== ''").dialog-content-table
        table.table
          thead
            tr
              th Complain Code
              th Details
              th Edit
          tbody
            tr(v-for="(item, index) in incomingList.result" :key="item").table-content-group
              td(@click="changeStatus(item, incomingList.result4[index])" style="cursor: pointer") {{item}}
              td
                q-btn(color="button" size="md" icon="visibility" :ripple="false" @click="getComplaintSpecific(item,false, true)").button-view
              td
                q-btn(color="button" size="md" icon="edit" :ripple="false" @click="getComplaintSpecific(item, true, true)").button-view

    section(v-else).flex.flex-center
      section(v-if="incomingList.result !== ''").dialog-content-table
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
              td {{ item }}
              td {{ incomingList.result2[index] }}
              td {{ incomingList.result3[index] }}
              td( style="cursor: pointer" ) {{incomingList.result4[index]}}
              td
                q-btn(rounded size="sm" color="button" label="edit" :ripple="false" @click="openDetails(item, incomingList.result2[index], incomingList.result3[index], incomingList.result4[index])").button-view


q-dialog(full-width v-model="details" transition-show="flip-right" transition-hide="flip-left").dialog
  q-card.dialog-card.text-white(style="height: auto" )
    q-card-section.full-width.column.items-center
      section.full-width.row.justify-between
        //- span.detail-dialog__info--large {{entryCodeDetail}}
        //- span.detail-dialog__info--large {{dateReceivedDetail}}
        component(:is="docInfo" label="Complaint Code" :value="entryCodeDetail")
        component(:is="docInfo" label="Date" :value="dateReceivedDetail")
        component(:is="docInfo" label="Name" :value="nameDetail")
      //- span(style="padding: 2rem; font-weight: bold;").detail-dialog__info--large {{nameDetail}}
      section.full-width.row
        component(:is="docInfo" label="Subject:" :value="subjectDetail")
        //- span.detail-dialog__info Subject:
        //-   span.detail-dialog__info--detail {{subjectDetail}}
      section.full-width.row
        component(:is="docInfo" label="Details:" :value="detailsDetail")
        //- span.detail-dialog__info Details:
        //-   span.detail-dialog__info--detail {{detailsDetail}}

    q-card-section.full-width.column.items-center
      section.full-width.column.justify-around
        div(v-if="attachmentDetail !== ''")
          component(:is="docInfo" label="Attachments:" :value="attachmentDetail")
        //- span(v-if="attachmentDetail !== null").detail-dialog__info--subinfo Attachments:
          span.detail-dialog__info--subdetail {{attachmentDetail}}
        div(v-if="noteDetail !== ''")
          component(:is="docInfo" label="Notes:" :value="noteDetail")
        //- span(v-if="noteDetail !== null").detail-dialog__info--subinfo Notes:
        //-   span.detail-dialog__info--subdetail {{noteDetail}}

    q-card-section(v-if="doclogEmpty === false").doc-log-area.full-width.column.items-center
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

    q-card-section(v-else).doc-log-area.full-width.column.items-center
      span.detail-dialog__info No Document Logs Found

    q-card-section(v-if="actionlogEmpty === false").full-width.column.items-center
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

    q-card-section(v-else).full-width.column.items-center
      span.detail-dialog__info No Actions Logs Found

    q-card-actions(align="center")
      component(:is="docButton" text="Close" v-close-popup).detail-dialog__button
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { api } from 'boot/axios'
import { useRouter } from 'vue-router'
import { useCurrentPage } from 'stores/currentpage'
import { useIsLogged } from 'stores/islogged'
import { useQuasar } from 'quasar'

import docButton from 'components/docButton.vue'
import docForm from 'components/docForm.vue'
import docInfo from 'components/docInfo.vue'
import docInfoEdit from 'components/docInfoEdit.vue'
import docLabel from 'components/docLabel.vue'
import docPDF2 from 'components/docPDF2.vue'
import { parseDate } from 'pdf-lib'

const router = useRouter()
const quasar = useQuasar()
const _currentpage = useCurrentPage()
const _islogged = useIsLogged()

let searchValue = ref('')
let searchByValue = ref('')
let nodata = ref(true)

type Incoming = {
  result: string
  result2: string
  result3: string
  result4: string
}
let incomingList = ref({} as Incoming)
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
    nodata.value = false
  } else nodata.value = true
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

const gotoHome = () => {
  let homepage = ''
  if (_islogged.islogged === true) homepage = 'dashboard'
  else homepage = '/'

  _currentpage.currentpage = homepage
  router.push(homepage)
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

let doclogEmpty = ref(false)
const getIncomingDocLog = async () => {
  try {
    const response = await api.get('/api/GetIncomingDocLog/' + entryCodeDetail.value)
    const data = response.data

    if (data.result.length > 0) {
      doclogDetail.value = data
      doclogEmpty.value = false
    } else doclogEmpty.value = true
  } catch {
    doclogDetail.value = []
    doclogEmpty.value = true
  }
}

let actionlogEmpty = ref(false)
const getIncomingActionLog = async () => {
  try {
    const response = await api.get('/api/GetIncomingActionLog/' + entryCodeDetail.value)
    const data = response.data
    if (data.result.length > 0) {
      actionlogDetail.value = data
      actionlogEmpty.value = false
    } else actionlogEmpty.value = true
  } catch {
    actionlogDetail.value = []
    actionlogEmpty.value = true
  }
}

;(async () => {
  await getIncoming()
  await searchByCode()

  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/incominginquire')
})()
</script>

<style lang="sass" scoped>


.table-loading
  font-size: 2rem
  margin-top: 3rem

// .table
//   font-family: 'Montserrat'
//   font-size: 0.8rem
//   text-transform: uppercase
//   border-collapse: collapse
//   margin: 2rem 0
//   min-width: 25rem
//   border-radius: 1rem 1rem 0 0
//   overflow: hidden
//   box-shadow: 0 0 20px rgba(0, 0, 0, 0.2)

// .table thead tr
//   background-color: #000406
//   color: #ffffff
//   text-align: left
//   font-weight: bold

// .table th, .table td
//   padding: 1rem 2.5rem
//   width: 20rem

// .table tbody tr
//   border-bottom: 1px solid #dddddd

// .table tbody tr:nth-of-type(even)
//   background-color: #1C4157

// .table tbody tr:last-of-type
//   border-bottom: 2px solid #000406

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
  background-color: $background
  color: #ffffff
  margin: 1rem
  border: 1px solid #2F5972
  border-radius: 12rem

.detail-dialog__info
  font-size: 1.4rem
  font-family: 'Montserrat'

.detail-dialog__info--large
  font-family: 'Montserrat'
  font-size: 2.2rem

.detail-dialog__info--subinfo
  font-family: 'Montserrat'
  font-size: 1.2rem

.detail-dialog__info--detail
  font-size: 1.6rem
  font-family: 'Montserrat'
  font-weight: bold
  margin-left:1rem

.detail-dialog__info--subdetail
  font-size: 1.4rem
  font-family: 'Montserrat'
  margin-left:1rem

.detail-dialog__button
  margin-bottom: 2rem

.doc-log-area
  margin-top: 2rem
</style>
