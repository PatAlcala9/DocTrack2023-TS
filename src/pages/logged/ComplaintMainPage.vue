<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Complaint Tracking
    q-btn(flat size="md" label="Back" @click="gotoComplaint" icon="arrow_back").close-button

  div.fit.row.wrap.justify-start.items-start.content-start.search-area
    section.column.section
      component(:is="docLabel" text="Search")
      q-input(dense filled v-model="searchValue" bg-color="white" :input-style="{ fontSize: '1.2rem', fontFamily: 'OpenSans' }" @keydown.enter="getIncomingUsingValue")
    section.column.section
      component(:is="docLabel" text="Search By")
      //- input
      q-btn-dropdown(unelevated color="grey-10" :label="searchByValue" size="1rem" :content-style="{ fontSize: '0.9rem', fontFamily: 'Inter' }")
        q-list
          q-item(clickable v-close-popup @click="searchByValue = 'CODE'")
            q-item-section
              q-item-label COMPLAINT CODE

          q-item(clickable v-close-popup @click="searchByValue = 'RESPONDENT'")
            q-item-section
              q-item-label RESPONDENT

          q-item(clickable v-close-popup @click="searchByValue = 'INSPECTOR'")
            q-item-section
              q-item-label INSPECTOR

          q-item(clickable v-close-popup @click="searchByValue = 'LOCATION'")
            q-item-section
              q-item-label LOCATION

          q-item(clickable v-close-popup @click="searchByValue = 'STATUS'")
            q-item-section
              q-item-label STATUS

          q-item(clickable v-close-popup @click="searchByValue = 'REMAINING DAYS'")
            q-item-section
              q-item-label REMAINING DAYS

  div(v-if="quasar.screen.width <= 500").flex.flex-center
    section(v-if="complaintList.result !== ''").dialog-content-table
      table.table
        thead
          tr
            th Complain Code
            th Details
        tbody
          tr(v-for="(item, index) in complaintList.result" :key="item").table-content-group
            td {{item}}
            td
              q-btn(color="button" icon="visibility" :ripple="false" ).button-view

  div(v-else).flex.flex-center
    section(v-if="complaintList.result !== ''").dialog-content-table
      table.table
        thead
          tr
            th Complain Code
            th Respondent
            th Inspector
            th Location
            th Status
            th Remaining Days
            th Details
        tbody
          tr(v-for="(item, index) in complaintList.result" :key="item").table-content-group
            td {{item}}
            td {{complaintList.result2[index]}}
            td {{complaintList.result3[index]}}
            td {{complaintList.result4[index]}}
            td {{complaintList.result5[index]}}
            td {{complaintList.result6[index]}}
            td
              q-btn(color="button" icon="visibility" :ripple="false" ).button-view
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { date, useQuasar } from 'quasar'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'
import { useCurrentPage } from 'stores/currentpage'

const router = useRouter()
const quasar = useQuasar()
const _currentpage = useCurrentPage()

import docButton from 'components/docButton.vue'
import docTextArea from 'components/docTextArea.vue'
import docInput from 'components/docInput.vue'
import docLabel from 'components/docLabel.vue'

let searchValue = ref('')
let searchByValue = ref('')

type Complaint = {
  result: string
  result2: string
  result3: string
  result4: string
  result5: string
  result6: string
}
let complaintList = ref({} as Complaint)

const gotoComplaint = () => {
  _currentpage.currentpage = 'complaint'
  router.push('/complaint')
}

const pushSampleData = () => {
  let sampleData = {
    result: ['23-1-0052', '23-2-0041'],
    result2: ['JUAN DELA CRUZ', 'HARRY POTTER'],
    result3: ['MR. BEAN', 'SPIDER-MAN'],
    result4: ['GOTHAM CITY', 'MARS'],
    result5: ['FOR NOTICE OF VIOLATION SERVING', 'FIRST NOTICE OF VIOLATION SERVED'],
    result6: ['5', '3'],
  }
  complaintList.value = sampleData
}
pushSampleData()
</script>

<style lang="sass" scoped>

.table-loading
  font-size: 2rem
  margin-top: 3rem

.table
  font-family: 'Inter'
  font-weight: 100
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
  padding: 1rem 4rem
  width: 20rem

.table tbody tr
  border-bottom: 1px solid #dddddd
  font-weight: 250

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

.section
  margin-bottom: 0

@media screen and (max-width: 500px)
  .section
    margin-bottom: 1.2rem
</style>
