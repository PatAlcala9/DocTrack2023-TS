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

  div(v-if="nodata")
    span No Data Found

  div(v-else)
    section(v-if="quasar.screen.width <= 500").flex.flex-center
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

    section(v-else).flex.flex-center
      section(v-if="complaintList.result !== ''").dialog-content-table
        table.table
          thead
            tr
              th Complain Code
              th Type
              th Name
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
              td
                q-btn(color="button" icon="visibility" :ripple="false" @click="getComplaintSpecific(item)").button-view

q-dialog(v-model="dialog" transition-show="flip-right" transition-hide="flip-left").dialog
  q-card.dialog-card.text-white.flex.flex-center
    q-card-section.dialog-card__section
      div.dialog-title-area.column.justify-center.items-center
        span.dialog-card__label Complaint Code:
        span.dialog-card__data {{ dialogCode }}
        span.dialog-card__label Complaint Type:
        span.dialog-card__data {{ dialogType }}
        span.dialog-card__label Received Date:
        span.dialog-card__data {{ dialogReceivedDate }}

        span.dialog-card__label Complaintant:
        span.dialog-card__data {{ dialogName }}
        span.dialog-card__label Complaintant Contact:
        span.dialog-card__data {{ dialogContact }}
        span.dialog-card__label Complaintant Location:
        span.dialog-card__data {{ dialogLocation }}
        span.dialog-card__label Details:
        span.dialog-card__data {{ dialogDetails }}

        span.dialog-card__label Details:
        span.dialog-card__data {{ dialogDetails }}

        span.dialog-card__label Transaction Date :
        span.dialog-card__data {{ dialogDateTransacted }}

        span.dialog-card__label Respodent:
        span.dialog-card__data {{ dialogRespondentName }}
        span.dialog-card__label Respodent Location:
        span.dialog-card__data {{ dialogRespondentLocation }}
        span.dialog-card__label Respodent Contact:
        span.dialog-card__data {{ dialogRespondentContact }}

        doc-button(text="OK" @click="dialog=false")
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
let nodata = ref(true)

let dialog = ref(false)
let dialogCode = ref('')
let dialogType = ref('')
let dialogName = ref('')
let dialogContact = ref('')
let dialogReceivedDate = ref('')
let dialogLocation = ref('')
let dialogDetails = ref('')
let dialogRespondentName = ref('')
let dialogRespondentContact = ref('')
let dialogRespondentLocation = ref('')
let dialogStatus = ref('')
let dialogDateTransacted = ref('')

type Complaint = {
  result: string
  result2: string
  result3: string
  result4: string
  result5: string
}
let complaintList = ref({} as Complaint)

const gotoComplaint = () => {
  _currentpage.currentpage = 'complaint'
  router.push('/complaint')
}

const getComplaintList = async () => {
  const response = await api.get('/api/GetComplaintList')
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    const result = data.result
    if (result.length > 0) {
      let arrayRemainingDays: string[] = []

      complaintList.value.result = data.result
      complaintList.value.result2 = data.result2
      complaintList.value.result3 = data.result3
      complaintList.value.result4 = data.result4

      for (let item of data.result5) {
        const remainingDays = (await calculateRemainingDays(item)).toString()
        arrayRemainingDays.push(data.result4.toString().includes('SERVED') ? remainingDays : '')
      }

      console.log('arrayRemainingDays:', arrayRemainingDays)
      complaintList.value.result5 = arrayRemainingDays.toString()

      return true
    } else return false
  } else return false
}

const calculateRemainingDays = async (expiry: string): Promise<number> => {
  const expiryDate = date.formatDate(expiry, 'DDD')
  const today = date.formatDate(new Date(), 'DDD')

  if (expiryDate > today) {
    const remainingDays = parseInt(expiryDate) - parseInt(today)

    const currentDayOfWeek = new Date().getDay()
    const remainingWeekdays = Math.max(remainingDays - Math.floor(remainingDays / 7) * 2, 0)

    let adjustedRemainingDays = remainingWeekdays
    if (currentDayOfWeek === 6) {
      adjustedRemainingDays = Math.max(remainingWeekdays - 1, 0)
    } else if (currentDayOfWeek === 0) {
      adjustedRemainingDays = Math.max(remainingWeekdays - 2, 0)
    }
    return adjustedRemainingDays
  } else return 0
}

const getComplaintSpecific = async (code: string) => {
  const response = await api.get('/api/GetComplaintSpecific/' + code)
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    dialogCode.value = code
    dialogType.value = data.result
    dialogName.value = data.result2
    dialogContact.value = data.result3
    dialogLocation.value = data.result4
    dialogReceivedDate.value = data.result5
    dialogDetails.value = data.result6
    dialogRespondentName.value = data.result7
    dialogRespondentContact.value = data.result8
    dialogRespondentLocation.value = data.result9
    dialogStatus.value = data.result10
    dialogDateTransacted.value = data.result11

    dialog.value = true
  }
}

;(async () => {
  if (await getComplaintList()) nodata.value = false
  else nodata.value = true
})()
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
  padding: 1rem 2.5rem
  width: 20rem

.table tbody tr
  border-bottom: 1px solid #dddddd
  font-weight: 350

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


.doc-log-area
  margin-top: 2rem

.section
  margin-bottom: 0

.dialog
  width: 100vw

.dialog-card
  width: 100vw
  height: 100vh

.dialog-card__label
  font-family: 'Inter'
  font-weight: 300
  font-size: 1rem

.dialog-card__data
  font-family: 'Inter'
  font-weight: 500
  font-size: 1.1rem
  padding-bottom: 2rem

@media screen and (max-width: 500px)
  .section
    margin-bottom: 1.2rem
</style>
