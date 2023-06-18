<template lang="pug">

q-page(padding)
  section.online
    q-icon(name="circle" size="1rem" :color="onlineColor")


  div.full-width.row.justify-between
    span.title Complaint - Add New
    q-btn(flat size="md" label="Back" @click="gotoComplaintDashboard" icon="arrow_back").close-button

  div.entry-group
    div.complaint-group
      section.form
        section
          //- component(:is="docLabel" text="Source Type:").form--label
          component(:is="docSelection" text="Source Type" label="Select Type" :options="sourceEntryList" v-model:modelValue="sourceEntry")
        //- q-btn-dropdown(unelevated color="grey-10" :label="sourceEntry" size="1rem" :content-style="{ fontSize: '1.2rem', fontFamily: 'Inter' }")
        //-   q-list
        //-     q-item(clickable v-close-popup @click="sourceEntry = 'DCR/888'")
        //-       q-item-section
        //-         q-item-label DCR/888

        //-     q-item(clickable v-close-popup @click="sourceEntry = 'EMAIL'")
        //-       q-item-section
        //-         q-item-label EMAIL

        //-     q-item(clickable v-close-popup @click="sourceEntry = 'MOTU PROPRIO'")
        //-       q-item-section
        //-         q-item-label MOTU PROPRIO

        //-     q-item(clickable v-close-popup @click="sourceEntry = 'WALK-IN'")
        //-       q-item-section
        //-         q-item-label WALK-IN
        //-         div.full-width.row.justify-between

      section.form
        section.section
          span.form--label Date Received
          q-date(flat v-model="receivedDate" minimal color="$button" @click="formatDate").calendar
          component(v-if="formattedReceivedDate.length > 0" :is="docLabel" :text="formattedReceivedDate").form--label
          component(v-else :is="docLabel" text="No Date Selected").inputs__label--date.text-center
          //- component(:is="docCalendar" text="Date Received" v-model:modelValue="formattedReceivedDate" @click="formatDate")

      section.form
        section.section
          component(:is="docLabel" text="Complaint Name:").label--spaced
          component(:is="docInputEntry" v-model:value="complaintName" )

      section.form
        section.section
          component(:is="docLabel" text="Complaint Contact:").label--spaced
          component(:is="docInputEntry" v-model:value="complaintContact" )

      section.form
        section.section
          component(:is="docLabel" text="Complaint Location:").label--spaced
          component(:is="docInputEntry" v-model:value="complaintLocation" )

      section.form
        section.section
          component(:is="docLabel" text="Detail of Complaint:").label--spaced
          component(:is="docTextArea" v-model:value="complaintDetail" )

    div.respondent-group
      section.form
        section.section
          component(:is="docLabel" text="Respondent Name:").label--spaced
          component(:is="docInputEntry" v-model:value="respondentName")

      section.form
      section.section
        component(:is="docLabel" text="Respondent Contact:").label--spaced
        component(:is="docInputEntry" v-model:value="respondentContact")

      section.form
        component(:is="docLabel" text="Respondent Location:").label--spaced
        component(:is="docInputEntry" v-model:value="respondentLocation").label--spaced

    div.attachment-group
      component(:is="docList" text="Attachments" :options="attachmentList" v-model:modelValue="attachmentSelectedList").label--spaced

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
import { checkConnection } from 'src/js/functions'
import Dexie from 'dexie'

import { useCurrentPage } from 'stores/currentpage'
import { usePageWithTable } from 'stores/pagewithtable'
import { useIsDemo } from 'stores/isdemo'

import docButton from 'components/docButton.vue'
import docTextArea from 'components/docTextArea.vue'
import docInputEntry from 'components/docInputEntry.vue'
import docInput from 'components/docInput.vue'
import docLabel from 'components/docLabel.vue'
import docSelection from 'components/docSelection.vue'
import docCalendar from 'components/docCalendar.vue'
import docList from 'components/docList.vue'

const sourceEntryList = ['DCR/888', 'EMAIL', 'MOTU PROPRIO', 'LETTER']
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
let attachmentList = ref<{ label: string; value: string }[]>([])
let attachmentSelectedList = ref<string[]>([])

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

const getAttachmentListFromDatabase = async () => {
  try {
    const response = await api.get('/api/GetAttachmentList')
    const data = response.data.length !== 0 ? response.data : null
    if (data !== null)
      for (let i in data.result)
        attachmentList.value.push({
          label: data.result[i],
          value: data.result2[i],
        })
  } catch {
    attachmentList.value = []
  }
}

const generateNewComplaintCode = async (code: string) => {
  const today = new Date()
  const currentYear = date.formatDate(today, 'YY')

  if (code !== '') {
    const series = parseInt(code.substring(5)) + 1
    const newSeries = series.toString().padStart(4, '0')

    return `${currentYear}-${sourceEntryID.value}-${newSeries}`
  } else return `${currentYear}-${sourceEntryID.value}-0001`
}

const postComplaint = async (code: string, complaintid: number, complaintname: string, complaintcontact: string, datereceived: string, location: string, details: string, infoid: number, statusid: number): Promise<boolean> => {
  const response = await api.post('/api/PostComplaint', {
    data: code,
    data2: complaintid.toString(),
    data3: complaintname,
    data4: complaintcontact,
    data5: datereceived,
    data6: location,
    data7: details,
    data8: infoid.toString(),
    data9: statusid.toString(),
  })
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) return true
  else return false
}

const postComplaintLOCAL = async (code: string, complaintid: number, complaintname: string, complaintcontact: string, datereceived: string, location: string, details: string, infoid: number, statusid: number): Promise<boolean> => {
  const response = {
    code: code,
    complaintid: complaintid,
    complaintname: complaintname,
    complaintcontact: complaintcontact,
    datereceived: datereceived,
    location: location,
    details: details,
    infoid: infoid,
    statusid: statusid,
  }
  LocalStorage.set('complaint', response)

  return true
}

const postStatus = async (code: string, date: string, status: string, tagcode: string, tagword: string, receivedby: string, details: string): Promise<boolean> => {
  const response = await api.post('/api/PostStatus', {
    data: code,
    data2: date,
    data3: status,
    data4: tagcode,
    data5: tagword,
    data6: receivedby,
    data7: details,
  })
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) return true
  else return false
}

const postAttachment = async (code: string, ref: number): Promise<boolean> => {
  const response = await api.post('/api/PostAttachment', {
    data: code,
    data2: ref,
  })
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) return true
  else return false
}

const getLatestStatus = async (code: string): Promise<number> => {
  try {
    const response = await api.get('/api/GetLatestStatus/' + code)
    const data = response.data.length !== 0 ? response.data : null

    if (data !== null) {
      return data.result
    } else return 0
  } catch {
    return 0
  }
}

const showDialog = (title: string, message: string) => {
  dialog.value = true
  dialogTitle.value = title
  dialogMessage.value = message
}

const saveData = async () => {
  console.log('_isdemo.isdemo:', _isdemo.isdemo)
  if (checkComplete() === false) {
    if (_isdemo.isdemo) {
      await getSourceIDFromLOCAL()

      if ((await postRespodentLOCAL(respondentName.value, respondentContact.value, respondentLocation.value)) === true) {
        const latestRespondent = 0
        const maxComplaint = await getMaxComplaintCodeLOCAL()
        const newComplaint = await generateNewComplaintCode(maxComplaint)
        const newStatusID = await getLatestStatus(newComplaint)
        if (await postComplaintLOCAL(newComplaint, sourceEntryID.value, complaintName.value, complaintDetail.value, receivedDate.value, complaintLocation.value, complaintDetail.value, latestRespondent, newStatusID)) showDialog('Success', 'Successfully Saved Complaint')
        else showDialog('Error', 'Failed to Save Complaint')
      } else {
        showDialog('Error', 'Failed to Save Respondent')
      }
    } else {
      // for (let item of attachmentSelectedList.value) {
      //   console.log(item)
      // }
      if (await checkConnection()) {
        try {
          await getSourceIDFromDatabase()

          if ((await postRespondent(respondentName.value, respondentContact.value, respondentLocation.value)) === true) {
            const latestRespondent = await getLatestRespondent()
            const maxComplaint = await getMaxComplaintCode()
            const newComplaint = await generateNewComplaintCode(maxComplaint)

            if (await postStatus(newComplaint, receivedDate.value, 'ENCODED TO SYSTEM', '15', 'ENCODED', complaintName.value, '')) {
              const newStatusID = await getLatestStatus(newComplaint)
              if (await postComplaint(newComplaint, sourceEntryID.value, complaintName.value, complaintDetail.value, receivedDate.value, complaintLocation.value, complaintDetail.value, latestRespondent, newStatusID)) {
                for (let item of attachmentSelectedList.value) {
                  await postAttachment(newComplaint, parseInt(item))
                }
                showDialog('Success', 'Successfully Saved Complaint')
              }
              else showDialog('Error', 'Failed to Save Complaint')
            } else showDialog('Error', 'Failed to Save Complaint')
          } else showDialog('Error', 'Failed to Save Respondent')
        } catch {
          showDialog('Error', 'Failed to Save Complaint')
        }
      }
    }
  }
}

const createLOCALDATABASE = async () => {
  const db = new Dexie('ocbodoctracksys')

  interface source_complaint {
    source_complaintid: number
    source_desc: string
  }

  // interface complaint_info {
  //   complaint_infoid: number;
  //   complaint_code: string;
  //   source_complaintid: string;
  //   complaintant_name: string;
  //   complaintant_contact: number
  //   date_received: string;
  //   locationOfconstruction: string;
  //   details: string;
  //   respondent_infoid: number;
  //   complaintant_statusid: number;
  // }

  interface complaint_whereabouts {
    complaint_whereaboutsid: number
    whereabout: string
    tagword: string
    tagcode: string
  }

  db.version(1).stores({
    source_complaint: 'source_complaintid, source_desc',
    complaint_info: 'complaint_infoid, complaint_code, source_complaintid, complaintant_name, complaintant_contact, date_received, locationOfconstruction, details, respondent_infoid, complaintant_statusid',
    complaint_status: 'complaint_statusid, complaint_code, date_transacted, status, tagcode, tagword, received_by, remarks',
    respondent_info: 'respondent_infoid, respondent_name, respondent_contact, respondent_location',
    complaint_whereabouts: 'complaint_whereaboutsid, whereabout, tagword, tagcode',
  })

  db.open()
    .then(() => {
      const source_complaintData = db.table<source_complaint, number>('source_complaint')
      const source_complaintRows = [
        { source_complaintid: 1, source_desc: 'Email' },
        { source_complaintid: 2, source_desc: 'DCR/888' },
        { source_complaintid: 3, source_desc: 'Letter' },
        { source_complaintid: 4, source_desc: 'Motu Proprio' },
      ]
      source_complaintData.bulkPut(source_complaintRows)

      const complaint_whereaboutsData = db.table<complaint_whereabouts, number>('complaint_whereabouts')
      const complaint_whereaboutsRows = [
        { complaint_whereaboutsid: 1, whereabout: 'FOR INSPECTION AND VERIFICATION', tagword: 'INSPECTION', tagcode: '01' },
        { complaint_whereaboutsid: 2, whereabout: 'FIRST NOTICE OF VIOLATION SERVING', tagword: '1STNOVSERVING', tagcode: '02' },
        { complaint_whereaboutsid: 3, whereabout: 'FIRST NOTICE OF VIOLATION SERVED', tagword: '1STNOVSERVED', tagcode: '03' },
        { complaint_whereaboutsid: 4, whereabout: 'SECOND NOTICE OF VIOLATION AND WORK STOPPAGE ORDER SERVING', tagword: '2NDNOVSERVING', tagcode: '04' },
        { complaint_whereaboutsid: 5, whereabout: 'SECOND NOTICE OF VIOLATION AND WORK STOPPAGE ORDER SERVED', tagword: '2NDNOVSERVED', tagcode: '05' },
        { complaint_whereaboutsid: 6, whereabout: 'CALL OF DIALOGUE', tagword: '1STDIALOGUE', tagcode: '06' },
        { complaint_whereaboutsid: 7, whereabout: 'FOR NOTICE OF HEARING', tagword: 'NOTICEHEARING', tagcode: '07' },
        { complaint_whereaboutsid: 8, whereabout: 'NOTICE OF HEARING SCHEDULED', tagword: 'HEARINGSCHED', tagcode: '08' },
        { complaint_whereaboutsid: 9, whereabout: 'HEARING', tagword: 'HEARING', tagcode: '09' },
        { complaint_whereaboutsid: 10, whereabout: 'RESOLUTION/ORDER', tagword: 'RESORDER', tagcode: '10' },
        { complaint_whereaboutsid: 11, whereabout: 'FOR FILING OF CASE', tagword: 'FILINGCASE', tagcode: '11' },
        { complaint_whereaboutsid: 12, whereabout: 'CASE FILED', tagword: 'CASEFILED', tagcode: '12' },
        { complaint_whereaboutsid: 13, whereabout: 'ON GOING CASE', tagword: 'ONGOING', tagcode: '13' },
        { complaint_whereaboutsid: 14, whereabout: 'CASE CLOSED', tagword: 'CLOSED', tagcode: '14' },
        { complaint_whereaboutsid: 15, whereabout: 'ENCODED TO SYSTEM', tagword: 'ENCODED', tagcode: '15' },
      ]
      complaint_whereaboutsData.bulkPut(complaint_whereaboutsRows)

      return true
    })
    .catch(() => {
      return false
    })
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
  createLOCALDATABASE()
  getAttachmentListFromDatabase()
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

.row
  margin-bottom: 0.9rem

.form
  display: flex
  flex-direction: column
  flex-wrap: wrap
  justify-content: center
  align-items: center
  align-content: center
  text-align: center
  padding-bottom: 1rem

.form--label
  font-family: 'Inter'
  font-weight: 300
  font-size: 1.2rem
  color: #ffffff
  padding: 0.5rem

.calendar
  background-color: transparent
  border: 1px solid rgba(255, 255, 255, 0.525)
  font-family: 'Inter'
  font-size: 1.4rem
  border-radius: 12px
  :deep .q-date__calendar-item
    padding-top: 0.6rem
  :deep .block
    font-size: 1.1rem

.section
  display: flex
  flex-direction: column
  backdrop-filter: blur(1.6px) saturate(173%)
  // background-color: rgba(10, 10, 35, 0.52)
  // background-color: rgba(17, 25, 40, 0.82)
  background-color: rgba(17, 25, 40, 0.8)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 1.2rem 2rem 1.2rem 1.2rem


@media screen and (max-width: 500px)
  .entry-group
    display: flex
    flex-direction: column
    flex-wrap: wrap
    justify-content: center
    align-items: center
    align-content: center
</style>
