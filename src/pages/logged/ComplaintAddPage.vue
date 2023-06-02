<template lang="pug">

q-page(padding)
  section.online
    q-icon(name="circle" size="1rem" :color="onlineColor")

  div.full-width.row.justify-between
    span.title Complaint - Add New
    q-btn(flat size="md" label="Back" @click="gotoComplaintDashboard" icon="arrow_back").close-button

  div.entry-group
    div.complaint-group
      section.label-and-input
        component(:is="docLabel" text="Source Type:").label--spaced
        q-btn-dropdown(unelevated color="grey-10" :label="sourceEntry" size="1rem" :content-style="{ fontSize: '1.2rem', fontFamily: 'Inter' }")
          q-list
            q-item(clickable v-close-popup @click="sourceEntry = 'DCR/888'")
              q-item-section
                q-item-label DCR/888

            q-item(clickable v-close-popup @click="sourceEntry = 'EMAIL'")
              q-item-section
                q-item-label EMAIL

            q-item(clickable v-close-popup @click="sourceEntry = 'MOTU PROPRIO'")
              q-item-section
                q-item-label MOTU PROPRIO

            q-item(clickable v-close-popup @click="sourceEntry = 'WALK-IN'")
              q-item-section
                q-item-label WALK-IN
                div.full-width.row.justify-between

      section.label-and-input
        component(:is="docLabel" text="Date Received:").label--spaced
        q-date(flat v-model="receivedDate" minimal color="$button" @click="formatDate").calendar
        component(v-if="formattedReceivedDate.length > 0" :is="docLabel" :text="formattedReceivedDate").inputs__label--date
        component(v-else :is="docLabel" text="No Date Selected").inputs__label--date

      section.label-and-input
        component(:is="docLabel" text="Complaint Name:").label--spaced
        component(:is="docInputEntry" v-model:value="complaintName" )

      section.label-and-input
        component(:is="docLabel" text="Complaint Contact:").label--spaced
        component(:is="docInputEntry" v-model:value="complaintContact" )

      section.label-and-input
        component(:is="docLabel" text="Complaint Location:").label--spaced
        component(:is="docInputEntry" v-model:value="complaintLocation" )

      section.label-and-input
        component(:is="docLabel" text="Detail of Complaint:").label--spaced
        component(:is="docTextArea" v-model:value="complaintDetail" )

    div.respondent-group
      section.label-and-input
        component(:is="docLabel" text="Respondent Name:").label--spaced
        component(:is="docInputEntry" v-model:value="respondentName")

      section.label-and-input
        component(:is="docLabel" text="Respondent Contact:").label--spaced
        component(:is="docInputEntry" v-model:value="respondentContact")

      section.label-and-input
        component(:is="docLabel" text="Respondent Location:").label--spaced
        component(:is="docInputEntry" v-model:value="respondentLocation")

  div.flex.flex-center
    component(:is="docButton" text="Save" @click="saveData")

q-dialog(v-model="dialog" transition-show="flip-right" transition-hide="flip-left").dialog
  q-card.dialog-card.text-white.flex.flex-center
    q-card-section.dialog-card__section
      div.dialog-title-area.column.justify-center.items-center
        span.dialog-card__title {{ dialogTitle }}
        span.dialog-card__info {{ dialogMessage }}
        doc-button(text="OK" @click="dialog=false")
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { date, LocalStorage } from 'quasar'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'

import { useCurrentPage } from 'stores/currentpage'
import { usePageWithTable } from 'stores/pagewithtable'
import { useIsDemo } from 'stores/isdemo'

import docButton from 'components/docButton.vue'
import docTextArea from 'components/docTextArea.vue'
import docInputEntry from 'components/docInputEntry.vue'
import docInput from 'components/docInput.vue'
import docLabel from 'components/docLabel.vue'

let sourceEntry = ref('Select Source')
let sourceEntryID = ref(0)
let receivedDate = ref('')
let formattedReceivedDate = ref('')
let complaintName = ref('')
let complaintContact = ref('')
let complaintLocation = ref('')
let complaintDetail = ref('')
let respondentName = ref('')
let respondentContact = ref('')
let respondentLocation = ref('')

let dialog = ref(false)
let dialogTitle = ref('')
let dialogMessage = ref('')

const router = useRouter()
let _currentpage = useCurrentPage()
let _pagewithtable = usePageWithTable()
const _isdemo = useIsDemo()

let onlineColor = ref('')

const formatDate = () => {
  formattedReceivedDate.value = date.formatDate(Date.parse(receivedDate.value), 'MMMM D, YYYY')
}

let missingDetails: string[] = []
const checkComplete = () => {
  missingDetails = []

  if (!sourceEntry.value || sourceEntry.value === 'Select Source') missingDetails.push('source type')
  if (!receivedDate.value) missingDetails.push('date received')
  if (!complaintName.value) missingDetails.push('complaint name')
  if (!complaintContact.value) missingDetails.push('complaint contact')
  if (!complaintDetail.value) missingDetails.push('complaint detail')
  if (!complaintLocation.value) missingDetails.push('complaint location')
  if (!respondentName.value) missingDetails.push('respondent name')
  if (!respondentContact.value) missingDetails.push('respondent contact')
  if (!respondentLocation.value) missingDetails.push('respondent location')

  if (missingDetails.length > 0) return true
  else return false
}

const getSourceIDFromDatabase = async () => {
  const encodedSource = sourceEntry.value.replace('/', '~')
  const response = await api.get('/api/GetSourceID/' + encodedSource)
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    sourceEntryID.value = data.result
  }
}

const getSourceIDFromLOCAL = async () => {
  switch (sourceEntry.value) {
    case 'Email':
      sourceEntryID.value = 1
      break
    case 'DCR/888':
      sourceEntryID.value = 2
      break
    case 'Walk-In':
      sourceEntryID.value = 3
      break
    default:
      sourceEntryID.value = 4
  }
}

const postRespondent = async (name: string, contact: string, location: string): Promise<boolean> => {
  const response = await api.post('/api/PostRespondent', {
    data: name,
    data2: contact,
    data3: location,
  })
  const data = response.data.length !== 0 ? response.data : null

  if (data.includes('Success')) return true
  else return false
}

const postRespodentLOCAL = async (name: string, contact: string, location: string): Promise<boolean> => {
  const respodentData = {
    respondentname: name,
    respondentcontact: contact,
    respondentlocation: location,
  }
  LocalStorage.set('respondent', respodentData)

  return true
}

const getLatestRespondent = async (): Promise<number> => {
  const response = await api.get('/api/GetLatestRespondentID')
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    const result = data.result
    return result > 0 ? result : 0
  } else return 0
}

const getMaxComplaintCode = async (): Promise<string> => {
  const response = await api.get('/api/GetMaxComplaintCode')
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) return data.result
  else return ''
}

const getMaxComplaintCodeLOCAL = async (): Promise<string> => {
  const response = LocalStorage.getItem('complaintcode')
  const data = response?.toString() ?? ''

  if (data !== null) return data
  else return ''
}

const generateNewComplaintCode = async (code: string) => {
  const today = new Date()
  const currentYear = date.formatDate(today, 'YY')
  const series = parseInt(code.substring(5)) + 1
  const newSeries = series.toString().padStart(4, '0')

  return `${currentYear}-${sourceEntryID.value}-${newSeries}`
}

const postComplaint = async (code: string, complaintid: number, complaintname: string, complaintcontact: string, datereceived: string, location: string, details: string, infoid: number): Promise<boolean> => {
  const response = await api.post('/api/PostComplaint', {
    data: code,
    data2: complaintid,
    data3: complaintname,
    data4: complaintcontact,
    data5: datereceived,
    data6: location,
    data7: details,
    data8: infoid,
  })
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) return true
  else return false
}

const postComplaintLOCAL = async (code: string, complaintid: number, complaintname: string, complaintcontact: string, datereceived: string, location: string, details: string, infoid: number): Promise<boolean> => {
  const response = {
    code: code,
    complaintid: complaintid,
    complaintname: complaintname,
    complaintcontact: complaintcontact,
    datereceived: datereceived,
    location: location,
    details: details,
    infoid: infoid,
  }
  LocalStorage.set('complaint', response)

  return true
}

const saveData = async () => {
  if (checkComplete() === false) {
    if (_isdemo) {
      await getSourceIDFromLOCAL()

      if ((await postRespodentLOCAL(respondentName.value, respondentContact.value, respondentLocation.value)) === true) {
        console.log('1')
        const latestRespondent = 0
        const maxComplaint = await getMaxComplaintCodeLOCAL()
        const newComplaint = await generateNewComplaintCode(maxComplaint)

        if (await postComplaintLOCAL(newComplaint, sourceEntryID.value, complaintName.value, complaintDetail.value, receivedDate.value, complaintLocation.value, complaintDetail.value, latestRespondent)) openDialog('Success', 'Successfully Saved Complaint')
        else openDialog('Error', 'Failed to Save Complaint')
      } else {
        openDialog('Error', 'Failed to Save Respondent')
      }
    } else {
      await getSourceIDFromDatabase()

      if ((await postRespondent(respondentName.value, respondentContact.value, respondentLocation.value)) === true) {
        const latestRespondent = await getLatestRespondent()
        const maxComplaint = await getMaxComplaintCode()
        const newComplaint = await generateNewComplaintCode(maxComplaint)

        if (await postComplaint(newComplaint, sourceEntryID.value, complaintName.value, complaintDetail.value, receivedDate.value, complaintLocation.value, complaintDetail.value, latestRespondent)) openDialog('Success', 'Successfully Saved Complaint')
        else openDialog('Error', 'Failed to Save Complaint')
      } else {
        openDialog('Error', 'Failed to Save Respondent')
      }
    }
  }
}

const openDialog = (title: string, message: string) => {
  dialog.value = true
  dialogTitle.value = title
  dialogMessage.value = message
}

const checkOnline = () => {
  if (_isdemo.isdemo) onlineColor.value = 'red'
  else {
    onlineColor.value = 'green'
  }
}

const gotoComplaintDashboard = () => {
  _pagewithtable.pagewithtable = false
  _currentpage.currentpage = 'complaint'
  router.push('/complaint')
}

;(async () => {
  if (_currentpage.currentpage !== undefined) router.push(_currentpage.currentpage)
  else router.push('/dashboard')

  checkOnline()
})()
</script>

<style lang="sass" scoped>
.entry-group
  display: flex
  flex-direction: column
  flex-wrap: wrap
  justify-content: center
  align-items: center
  margin-top: 2rem
  padding: 1rem 1rem

.complaint-group
  flex-grow: 1
  padding: 2rem
  // border: 1px solid white

.respondent-group
  flex-grow: 1

.label--spaced
  padding-right: 1rem

.row
  margin-bottom: 0.9rem

.label-and-input
  display: flex
  flex-direction: column
  flex-wrap: wrap
  justify-content: center
  align-items: center
  align-content: center
  padding-bottom: 1rem

@media screen and (max-width: 500px)
  .entry-group
    display: flex
    flex-direction: column
    flex-wrap: wrap
    justify-content: center
    align-items: center
    align-content: center
</style>
