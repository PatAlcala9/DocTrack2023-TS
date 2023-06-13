<template lang="pug">

q-page(padding)
  section.online
    q-icon(name="circle" size="1rem" :color="onlineColor")

  div.full-width.row.justify-between
    span.title Complaint Tracking
    q-btn(flat size="md" label="Back" @click="gotoComplaint" icon="arrow_back").close-button

  div.fit.row.wrap.justify-center.search-area
    component(:is="docForm" text="Search" v-model:value="searchValue" :width=30 :mobileWidth=14 icon="search")

  div(v-if="nodata").fit.row.wrap.justify-center
    span.flex.flex-center.nodata--text No Data Found

  div(v-else).fit.row.wrap.justify-center
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
import { date, useQuasar, LocalStorage } from 'quasar'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'
import { useCurrentPage } from 'stores/currentpage'
import { useIsDemo } from 'stores/isdemo'

const router = useRouter()
const quasar = useQuasar()
const _currentpage = useCurrentPage()
const _isdemo = useIsDemo()

let onlineColor = ref('')

import docButton from 'components/docButton.vue'
import docTextArea from 'components/docTextArea.vue'
import docInput from 'components/docInput.vue'
import docLabel from 'components/docLabel.vue'
import docForm from 'components/docForm.vue'

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

const checkOnline = () => {
  if (_isdemo.isdemo) onlineColor.value = 'red'
  else {
    onlineColor.value = 'green'
  }
}

const fillupOffline = () => {
  // dialogCode.value = code
  // dialogType.value = data.result
  // dialogName.value = data.result2
  // dialogContact.value = data.result3
  // dialogLocation.value = data.result4
  // dialogReceivedDate.value = data.result5
  // dialogDetails.value = data.result6
  // dialogRespondentName.value = data.result7
  // dialogRespondentContact.value = data.result8
  // dialogRespondentLocation.value = data.result9
  // dialogStatus.value = data.result10
  // dialogDateTransacted.value = data.result11
  const complaint = LocalStorage.getItem('complaint')
}

;(async () => {
  checkOnline()

  if (_isdemo) fillupOffline()
  else {
    if (await getComplaintList()) nodata.value = false
    else nodata.value = true
  }
})()
</script>

<style lang="sass" scoped>

.table-loading
  font-size: 2rem
  margin-top: 3rem

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

.nodata--text
  font-family: 'Inter'
  font-weight: 400
  font-size: 1.6rem
  margin-top: 6rem
  background-color: rgba(17, 25, 40, 0.8)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 1.2rem 9rem

@media screen and (max-width: 500px)
  .section
    margin-bottom: 1.2rem

  .nodata--text
    font-family: 'Inter'
    font-weight: 300
    font-size: 1.4rem
    padding: 1rem 4.5rem
</style>
