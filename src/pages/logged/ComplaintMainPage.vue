<template lang="pug">

q-page(padding)
  section.online
    q-icon(name="circle" size="1rem" :color="onlineColor")

  div.full-width.row.justify-between
    span.title Complaint Tracking
    q-btn(flat size="md" label="Back" @click="gotoComplaint" icon="arrow_back").close-button

  div.fit.row.wrap.justify-center.search-area
    component(:is="docForm" text="Search" v-model:value="searchValue" :width=30 :mobileWidth=14 icon="search" @keypress.enter="filterTable")

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
              th Edit
          tbody
            tr(v-for="(item, index) in complaintList.result" :key="item").table-content-group
              td(@click="changeStatus(item, complaintList.result4[index])" style="cursor: pointer") {{item}}
              td
                q-btn(color="button" size="md" icon="visibility" :ripple="false" @click="getComplaintSpecific(item,false, true)").button-view
              td
                q-btn(color="button" size="md" icon="edit" :ripple="false" @click="getComplaintSpecific(item, true, true)").button-view

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
              th Update
          tbody
            tr(v-for="(item, index) in complaintList.result" :key="item").table-content-group
              td(v-if="complaintList.result5[index].includes('-') || complaintList.result5[index] === '0'" style="background-color: rgba(128, 21, 21, 0.45)") {{item}}
              td(v-else) {{ item }}

              td(v-if="complaintList.result5[index].includes('-') || complaintList.result5[index] === '0'" style="background-color: rgba(128, 21, 21, 0.45)") {{complaintList.result2[index]}}
              td(v-else) {{ complaintList.result2[index] }}

              td(v-if="complaintList.result5[index].includes('-') || complaintList.result5[index] === '0'" style="background-color: rgba(128, 21, 21, 0.45)") {{complaintList.result3[index]}}
              td(v-else) {{ complaintList.result3[index] }}

              td(v-if="complaintList.result5[index].includes('-') || complaintList.result5[index] === '0'" style="background-color: rgba(128, 21, 21, 0.45); cursor: pointer" @click="changeStatus(item, complaintList.result4[index])") {{complaintList.result4[index]}}
              td(v-else @click="changeStatus(item, complaintList.result4[index])" style="cursor: pointer") {{complaintList.result4[index]}}

              td(v-if="complaintList.result5[index].includes('-') || complaintList.result5[index] === '0'" style="background-color: rgba(128, 21, 21, 0.45)") {{complaintList.result5[index]}}
              td(v-else) {{complaintList.result5[index]}}

              td(v-if="complaintList.result5[index].includes('-') || complaintList.result5[index] === '0'" style="background-color: rgba(128, 21, 21, 0.45)")
                q-btn(rounded size="sm" color="button" label="show" :ripple="false" @click="getComplaintSpecific(item, false, true)").button-view
              td(v-else)
                q-btn(rounded size="sm" color="button" label="show" :ripple="false" @click="getComplaintSpecific(item, false, true)").button-view

              td(v-if="complaintList.result5[index].includes('-') || complaintList.result5[index] === '0'" style="background-color: rgba(128, 21, 21, 0.45)")
                q-btn(rounded size="sm" color="button" label="edit" :ripple="false" @click="getComplaintSpecific(item, true, true)").button-view
              td(v-else)
                q-btn(rounded size="sm" color="button" label="edit" :ripple="false" @click="getComplaintSpecific(item, true, true)").button-view

//- q-dialog(full-width full-height v-model="dialog" transition-show="flip-right" transition-hide="flip-left").dialog
//-   q-card.dialog-card.text-white
//-     q-card-section
//-       section(v-if="quasar.screen.width > 500").fit.row.wrap.justify-between.items-center.content-center.text-center.q-card--section
//-         div.column
//-           component(:is="docInfo" label="Complaint Code" :value="dialogCode" )
//-         div.column
//-           component(:is="docInfo" label="Complaint Type" :value="dialogType" )
//-         div.column
//-           component(:is="docInfo" label="Received Date" :value="dialogReceivedDate" )
//-         div.column
//-           component(:is="docInfo" label="Status" :value="dialogStatus" )
//-       section(v-else).fit.column.wrap.justify-center.items-center.content-center.text-center.q-card--section
//-         div.column
//-           component(:is="docInfo" label="Complaint Code" :value="dialogCode" )
//-         div.column
//-           component(:is="docInfo" label="Complaint Type" :value="dialogType" )
//-         div.column
//-           component(:is="docInfo" label="Received Date" :value="dialogReceivedDate" )

//-       section.full-width.wrap.column.wrap.justify-center.items-center.content-center.text-center
//-         div.padded
//-           component(:is="docInfo" label="Complaintant" :value="dialogName" )
//-         div.padded
//-           component(:is="docInfo" label="Complaintant Contact" :value="dialogContact" )
//-         div.padded
//-           component(:is="docInfo" label=" Complaintant Location" :value="dialogLocation" )
//-         div.padded
//-           component(:is="docInfo" label="Details" :value="dialogDetails" wide)

//-         div.padded
//-           component(:is="docInfo" label="Transaction Date" :value="dialogDateTransacted")
//-         div.padded
//-           component(:is="docInfo" label="Respodent" :value="dialogRespondentName")
//-         div.padded
//-           component(:is="docInfo" label="Respodent Location" :value="dialogRespondentLocation")
//-         div.padded
//-           component(:is="docInfo" label="Respodent Contact" :value="dialogRespondentContact")

//-       section.fit.row.wrap.justify-around.items-center.content-center.button-area
//-         component(:is="docButton" text="OK" @click="dialog=false")

q-dialog(full-width full-height v-model="dialog" transition-show="flip-right" transition-hide="flip-left").dialog
  q-card.dialog-card.text-white
    q-card-section.grid
      div.code
        component(:is="docInfo" label="Complaint Code" :value="dialogCode" )
      div.type
        component(:is="docInfo" label="Complaint Type" :value="dialogType" )
      div.received
        component(:is="docInfo" label="Received Date" :value="dialogReceivedDate" )
      div.status
        component(:is="docInfo" label="Status" :value="dialogStatus" left)
      div.generate
        component(:is="docPDF2" title="Sample Document PDF" text="Generate PDF" date="July 12, 2023")

      div.complaintant
        component(:is="docInfo" label="Complaintant" :value="dialogName" left)
      div.contact
        component(:is="docInfo" label="Complaintant Contact" :value="dialogContact" left)
      div.location
        component(:is="docInfo" label=" Complaintant Location" :value="dialogLocation" left)
      div.details
        component(:is="docInfo" label="Details" :value="dialogDetails" wide left)

      div.respondent
        component(:is="docInfo" label="Respodent" :value="dialogRespondentName" left)
      div.transaction
        component(:is="docInfo" label="Transaction Date" :value="dialogDateTransacted")

      div.respondent-contact
        component(:is="docInfo" label="Respodent Contact" :value="dialogRespondentContact" left)
      div.respondent-location
        component(:is="docInfo" label="Respodent Location" :value="dialogRespondentLocation" left)

      section.fit.row.wrap.justify-around.items-center.content-center.button-area
        component(:is="docButton" text="OK" @click="dialog=false")


q-dialog(full-width full-height v-model="dialogEdit" transition-show="flip-right" transition-hide="flip-left" ).dialog
  q-card.dialog-card.text-white
    q-card-section
      section(v-if="quasar.screen.width > 500").fit.row.wrap.justify-between.items-center.content-center.text-center.q-card--section
        div.column
          component(:is="docInfo" label="Complaint Code" :value="dialogCode")
        div.column
          component(:is="docInfo" label="Complaint Type" :value="dialogType")
        div.column
          component(:is="docInfo" label="Received Date" :value="dialogReceivedDate")
        div.column
          component(:is="docInfo" label="Status" :value="dialogStatus")
      section(v-else).fit.column.wrap.justify-center.items-center.content-center.text-center.q-card--section
        div.column
          component(:is="docInfo" label="Complaint Code" :value="dialogCode")
        div.column
          component(:is="docInfo" label="Complaint Type" :value="dialogType")
        div.column
          component(:is="docInfo" label="Received Date" :value="dialogReceivedDate")

      section.full-width.wrap.column.wrap.justify-center.items-center.content-center.text-center
        div.padded
          component(:is="docInfoEdit" label="Complaintant" v-model:value="dataName" @blur="recordChange('name')")
        div.padded
          component(:is="docInfoEdit" label="Complaintant Contact" v-model:value="dataContact" @blur="recordChange('contact')")
        div.padded
          component(:is="docInfoEdit" label=" Complaintant Location" v-model:value="dataLocation" @blur="recordChange('location')")
        div.padded
          component(:is="docInfoEdit" label="Details" v-model:value="dataDetails" wide @blur="recordChange('details')")

        div.padded
          component(:is="docInfo" label="Transaction Date" :value="dialogDateTransacted" )
        div.padded
          component(:is="docInfoEdit" label="Respodent" v-model:value="dataRespondentName" @blur="recordChange('respondent-name')")
        div.padded
          component(:is="docInfoEdit" label="Respodent Location" v-model:value="dataRespondentLocation" @blur="recordChange('respondent-location')")
        div.padded
          component(:is="docInfoEdit" label="Respodent Contact" v-model:value="dataRespondentContact" @blur="recordChange('respondent-contact')")

      section.fit.row.wrap.justify-around.items-center.content-center.button-area
        component(:is="docButton" text="OK" @click="dialogEdit=false")


q-dialog(full-width v-model="dialogStatusEdit" transition-show="flip-right" transition-hide="flip-left").dialog
  q-card.dialog-card.text-white(style="height: auto" )
    q-card-section
      section.full-height.column.wrap.justify-center.items-center.content-center.q-card--section
        div.padded
          component(:is="docInfo" label="Current Status" :value="dialogStatus")
        div.padded.text-center
          component(:is="docLabel" text="New Status")
          q-select(v-if="$q.screen.width > 500" dark rounded outlined v-model="dialogNewStatus" :options="statusList" input-class="select" behavior="menu" ).select
          q-select(v-else dark rounded outlined v-model="dialogNewStatus" :options="statusList" input-class="select" behavior="dialog" ).select
        div.padded
          component(:is="docInfoEdit" label="Remarks" v-model:value="statusRemarks" )

      section.fit.row.wrap.justify-around.items-center.content-center.button-area
        component(:is="docButton" text="OK" @click="postChangeStatus(dialogNewStatus)")
        //- @click="dialogStatusEdit=false"

</template>

<script setup lang="ts">
import { ref, Ref } from 'vue'
import { date, useQuasar, LocalStorage } from 'quasar'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'
import { useCurrentPage } from 'stores/currentpage'
import { useIsDemo } from 'stores/isdemo'
import { useEmployeeName } from 'stores/employeename'
import { useEmployeeID } from 'stores/employeeid'

const router = useRouter()
const quasar = useQuasar()
const _currentpage = useCurrentPage()
const _isdemo = useIsDemo()
const _employeename = useEmployeeName()
const _employeeid = useEmployeeID()

let onlineColor = ref('')

import docButton from 'components/docButton.vue'
import docForm from 'components/docForm.vue'
import docInfo from 'components/docInfo.vue'
import docInfoEdit from 'components/docInfoEdit.vue'
import docLabel from 'components/docLabel.vue'
import docPDF2 from 'components/docPDF2.vue'
import { parseDate } from 'pdf-lib'

let searchValue = ref('')
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
let dialogNewStatus = ref('')

let dataName = ref('')
let dataContact = ref('')
let dataLocation = ref('')
let dataDetails = ref('')
let dataRespondentName = ref('')
let dataRespondentContact = ref('')
let dataRespondentLocation = ref('')
let dataStatus = ref('')

let dialogEdit = ref(false)
let dialogStatusEdit = ref(false)

let statusList: Ref<string[]> = ref([])
let statusListTagword: Ref<string[]> = ref([])
let statusListTagcode: Ref<string[]> = ref([])
let statusLabel = ref('')
let statusRemarks = ref('')

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
      complaintList.value.result5 = arrayRemainingDays.toString()
      return true
    } else return false
  } else return false
}

const getComplaintList2 = async () => {
  try {
    const response = await api.get('/api/GetComplaintList2')
    const data = response.data.length !== 0 ? response.data : null

    if (data !== null) {
      const result = data.result
      if (result.length > 0) {
        let arrayRemainingDays: string[] = []
        let arrayStatus: string[] = []

        complaintList.value.result = data.result
        complaintList.value.result2 = data.result2
        complaintList.value.result3 = data.result3

        for (let i in data.result4) {
          const [result, result2] = await getLatestStatusNameIndividual2(data.result[i])
          arrayStatus.push(result)
          // const expirationDate =
          console.log('result2', result2)
          // const expirationDate = parseDate(result2) + 15
          const expirationDate = new Date(result2)
          expirationDate.setDate(expirationDate.getDate() + 15)
          console.log('expirationDate', expirationDate)
          const remainingDays = (await calculateRemainingDays(result2)).toString()
          // arrayRemainingDays.push(arrayStatus[i].toString().includes('SERVED') ? remainingDays : '')

          if (arrayStatus[i].toString() === 'FIRST NOTICE OF VIOLATION SERVED') {
            arrayRemainingDays.push(remainingDays)
          } else if (arrayStatus[i].toString().includes('WORK STOPPAGE ORDER SERVED')) {
            arrayRemainingDays.push(remainingDays)
          } else {
            arrayRemainingDays.push('')
          }
        }

        complaintList.value.result4 = arrayStatus
        complaintList.value.result5 = arrayRemainingDays

        return true
      } else return false
    } else return false
  } catch {
    return false
  }
}

const getLatestStatusNameIndividual = async (code: string): Promise<string> => {
  try {
    const response = await api.get('/api/GetLatestStatusNameIndividual/' + code)
    const data = response.data.length !== 0 ? response.data : null

    // if (data !== null) {
    //   return data.result.toString()
    // } else {
    //   return ''
    // }

    return data?.result.toString() || ''
  } catch {
    return ''
  }
}

const getLatestStatusNameIndividual2 = async (code: string): Promise<[string, string]> => {
  try {
    const response = await api.get('/api/GetLatestStatusNameIndividual2/' + code)
    const data = response.data.length !== 0 ? response.data : null

    return [data?.result.toString(), data?.result2.toString()] || ['','']
  } catch {
    return ['', '']
  }
}

const getComplaintListFiltered = async (code: string) => {
  const response = await api.get('/api/GetComplaintListFiltered2/' + code)
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    const result = data.result
    if (result.length > 0) {
      let arrayRemainingDays: string[] = []
      let arrayStatus: string[] = []

      complaintList.value.result = data.result
      complaintList.value.result2 = data.result2
      complaintList.value.result3 = data.result3

      for (let i in data.result4) {
        arrayStatus.push(await getLatestStatusNameIndividual(data.result[i]))
        const remainingDays = (await calculateRemainingDays(data.result4[i])).toString()
        arrayRemainingDays.push(arrayStatus[i].toString().includes('SERVED') ? remainingDays : '0')
      }
      complaintList.value.result4 = arrayStatus
      complaintList.value.result5 = arrayRemainingDays

      return true
    } else return false
  } else return false
}

const calculateRemainingDays = async (expiry: string): Promise<number> => {
  const expiryDate = date.formatDate(expiry, 'DDD')
  const today = date.formatDate(new Date(), 'DDD')

  if (expiryDate >= today) {
    const remainingDays = parseInt(expiryDate) - parseInt(today)
    const currentDayOfWeek = new Date().getDay()
    const remainingWeekdays = Math.max(remainingDays - Math.floor(remainingDays / 7) * 2, 0)
    console.log('remainingDays', remainingDays)
    let adjustedRemainingDays = remainingWeekdays
    if (currentDayOfWeek === 6) {
      adjustedRemainingDays = Math.max(remainingWeekdays - 1, 0)
    } else if (currentDayOfWeek === 0) {
      adjustedRemainingDays = Math.max(remainingWeekdays - 2, 0)
    }
    console.log('adjustedRemainingDays', adjustedRemainingDays)
    return adjustedRemainingDays
  } else return parseInt(expiryDate) - parseInt(today)
}

const calculateRemainingDays2 = async (status: string): Promise<number> => {
  const today = date.formatDate(new Date(), 'DDD')
  let expiryDate

  if (status === 'FIRST NOTICE OF VIOLATION SERVED') {
    expiryDate = 15
    console.log('parseInt(today)', parseInt(today))

    const remainingDays = expiryDate + parseInt(today)
    const currentDayOfWeek = new Date().getDay()
    const remainingWeekdays = Math.max(remainingDays - Math.floor(remainingDays / 7) * 2, 0)
    console.log('remainingDays', remainingDays)
    let adjustedRemainingDays = remainingWeekdays
    if (currentDayOfWeek === 6) {
      adjustedRemainingDays = Math.max(remainingWeekdays - 1, 0)
    } else if (currentDayOfWeek === 0) {
      adjustedRemainingDays = Math.max(remainingWeekdays - 2, 0)
    }
    console.log('adjustedRemainingDays', adjustedRemainingDays)
    return adjustedRemainingDays
  } else return parseInt(expiryDate) - parseInt(today)
}

const getComplaintSpecific = async (code: string, edit: boolean, show: boolean) => {
  const response = await api.get('/api/GetComplaintSpecific/' + code)
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    dialogCode.value = code
    dialogType.value = data.result
    dialogName.value = data.result2
    dialogContact.value = data.result3
    dialogLocation.value = data.result4
    dialogReceivedDate.value = date.formatDate(data.result5, 'MMMM D, YYYY')
    dialogDetails.value = data.result6
    dialogRespondentName.value = data.result7
    dialogRespondentContact.value = data.result8
    dialogRespondentLocation.value = data.result9
    dialogStatus.value = data.result10
    dialogDateTransacted.value = date.formatDate(data.result11, 'MMMM D, YYYY')

    if (show) {
      if (edit) dialogEdit.value = true
      else dialog.value = true
    }
  }
  await recordData()
}

const changeStatus = async (code: string, status: string) => {
  await getComplaintSpecific(code, false, false)
  await getStatusList(status)

  if (dialogNewStatus.value === '') statusLabel.value = 'SELECT STATUS'
  else statusLabel.value = ''

  dialogStatus.value = status
  dialogStatusEdit.value = true
}

const getMaxStatusID = async (code: string): Promise<number> => {
  const response = await api.get('/api/GetMaxStatusID/' + code)
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    return data.result
  } else return 0
}

const postUpdateStatusID = async (id: number, date: string, code: string): Promise<boolean> => {
  try {
    const response = await api.post('/api/PostUpdateStatusID', {
      data: id.toString(),
      data2: date,
      data3: code,
    })
    const data = response.data

    if (data.includes('Success')) {
      return true
    } else return false
  } catch {
    return false
  }
}

const filterTable = async () => {
  if (searchValue.value.length > 0) {
    nodata.value = !(await getComplaintListFiltered(searchValue.value))
  } else {
    nodata.value = !(await getComplaintList2())
  }
}

const getStatusList = async (exception: string) => {
  const response = await api.get('/api/GetStatusList/' + exception)
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    for (let i in data.result) {
      statusList.value.push(data.result[i])
      statusListTagword.value.push(data.result2[i])
      statusListTagcode.value.push(data.result3[i])
    }
  }
}

let newStatusId = 0
let newStatusTagword = ''
let newStatusTagcode = ''
const getStatusSpecific = async (whereabout: string): Promise<boolean> => {
  try {
    const response = await api.get('/api/GetStatusSpecific/' + whereabout)
    const data = response.data.length !== 0 ? response.data : null

    if (data !== null) {
      newStatusId = data.result
      newStatusTagword = data.result2
      newStatusTagcode = data.result3

      return true
    } else return false
  } catch {
    return false
  }
}

const postStatus = async (code: string, receiveddate: string, newstatus: string, tagcode: string, tagword: string, receivedby: string, details: string): Promise<boolean> => {
  const encodedStatus = newstatus.replace('/', '~')
  let expirationGap = 0

  switch (newstatus) {
    case 'FIRST NOTICE OF VIOLATION SERVED':
      expirationGap  = 15
      break
    case 'FINAL NOTICE OF VIOLATION SERVED':
      expirationGap  = 15
      break
    case 'SECOND NOTICE OF VIOLATION AND WORK STOPPAGE ORDER SERVED':
      expirationGap = 5
      break
    case 'CALL FOR DIALOGUE':
      expirationGap = 5
      break
    case 'FOR NOTICE OF HEARING':
      expirationGap = 5
      break
    case 'CASE FILED':
      expirationGap = 15
      break
    case 'FINAL EXECUTORY':
      expirationGap = 15
      break
    case 'RESOLUTION ORDER':
      expirationGap = 15
      break
    case 'HEARING':
      expirationGap = 15
      break
    default:
      expirationGap = 0
      break
  }

  const expirationDate = parseInt(date.formatDate(receiveddate, 'DDD')) + expirationGap
  const base = new Date (new Date().getFullYear(), 0, 1)
  const expirationDateCombine = date.addToDate(base, { days: expirationDate })
  const expirationDateFormat = date.formatDate(expirationDateCombine, 'YYYY-MM-DD')

  const response = await api.post('/api/PostStatus', {
    data: code,
    data2: receiveddate,
    data3: encodedStatus,
    data4: tagcode,
    data5: tagword,
    data6: receivedby,
    data7: details,
    data8: expirationDateFormat
  })
  const data = response.data.length !== 0 ? response.data : null

  return data !== null
}

const postChangeStatus = async (newstatus: string) => {
  const today = new Date()
  const formattedDate = date.formatDate(today, 'YYYY-MM-DD')
  const dialogEncodedStatus = dialogNewStatus.value.replace('/', '~')
  const getStatusSpecificBool = await getStatusSpecific(dialogEncodedStatus)

  if (getStatusSpecificBool) {
    const encodedStatus = newstatus.replace('/', '~')

    const postStatusBool = await postStatus(dialogCode.value, formattedDate, encodedStatus, newStatusTagcode, newStatusTagword, _employeename.employeename, statusRemarks.value)
    if (postStatusBool) {
      await postUpdateStatusID(newStatusId, formattedDate, dialogCode.value)
      await refreshData()
    }
  }

  dialogStatusEdit.value = false
}

const postUserLog = async (userid: number, table: string, column: string, old_value: string, new_value: string, dateOfEdit: string) => {
  const response = await api.post('/api/PostUserLog', {
    data: userid,
    data2: table,
    data3: column,
    data4: old_value,
    data5: new_value,
    data6: dateOfEdit,
  })
  const data = response.data.length !== 0 ? response.data : null

  return data !== null
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

const postEditLog = async (table: string, column: string, old_value: string, new_value: string, date: string) => {
  const response = await api.post('/api/PostEditLog', {
    data: table,
    data2: column,
    data3: old_value,
    data4: new_value,
    data5: date,
  })
  const data = await response.data

  if (data.includes('Success') === false) {
    await postEditLog(table, column, old_value, new_value, date)
  }
}

const postUpdate = async (value: string, code: string, column: string) => {
  const response = await api.post('/api/PostUpdate', {
    data: value,
    data2: code,
    data3: column,
  })
  const data = await response.data

  if (data.includes('Success') === false) {
    await postUpdate(value, code, column)
  }
}

const recordData = async () => {
  dataName.value = dialogName.value
  dataContact.value = dialogContact.value
  dataLocation.value = dialogLocation.value
  dataDetails.value = dialogDetails.value
  dataRespondentName.value = dialogRespondentName.value
  dataRespondentContact.value = dialogRespondentContact.value
  dataRespondentLocation.value = dialogRespondentLocation.value
  dataStatus.value = dialogStatus.value
}

const recordChange = async (value: string) => {
  const today = new Date()
  const formattedDate = date.formatDate(today, 'YYYY-MM-DD HH:mm:ss')

  switch (value) {
    case 'name':
      if (dataName.value !== dialogName.value) {
        await postUserLog(_employeeid.employeeid, 'complaint_info', 'complaintant_name', dataName.value, dialogName.value, formattedDate)
        // postEditLog('complaint_info', 'complaintant_name', dataName.value, dialogName.value, formattedDate)
        await postUpdate(dataName.value, dialogCode.value, value)
        break
      }
    case 'contact':
      if (dataContact.value !== dialogContact.value) {
        await postUserLog(_employeeid.employeeid, 'complaint_info', 'complaintant_contact', dataContact.value, dialogContact.value, formattedDate)
        // postEditLog('complaint_info', 'complaintant_contact', dataContact.value, dialogContact.value, formattedDate)
        await postUpdate(dataContact.value, dialogCode.value, value)
        break
      }
    case 'location':
      if (dataLocation.value !== dialogLocation.value) {
        await postUserLog(_employeeid.employeeid, 'complaint_info', 'locationOfconstruction', dataContact.value, dialogContact.value, formattedDate)
        // postEditLog('complaint_info', 'locationOfconstruction', dataLocation.value, dialogLocation.value, formattedDate)
        await postUpdate(dataLocation.value, dialogCode.value, value)
        break
      }
    case 'details':
      if (dataDetails.value !== dialogDetails.value) {
        await postUserLog(_employeeid.employeeid, 'complaint_info', 'details', dataContact.value, dialogContact.value, formattedDate)
        // postEditLog('complaint_info', 'details', dataDetails.value, dialogDetails.value, formattedDate)
        await postUpdate(dataDetails.value, dialogCode.value, value)
        break
      }
    case 'respondent-name':
      if (dataRespondentName.value !== dialogRespondentName.value) {
        await postUserLog(_employeeid.employeeid, 'respondent_info', 'respondent_name', dataContact.value, dialogContact.value, formattedDate)
        // postEditLog('respondent_info', 'respondent_name', dataRespondentName.value, dialogRespondentName.value, formattedDate)
        await postUpdate(dataRespondentName.value, dialogCode.value, value)
        break
      }
    case 'respondent-contact':
      if (dataRespondentContact.value !== dialogRespondentContact.value) {
        await postUserLog(_employeeid.employeeid, 'respondent_info', 'respondent_contact', dataContact.value, dialogContact.value, formattedDate)
        // postEditLog('respondent_info', 'respondent_contact', dataRespondentContact.value, dialogRespondentContact.value, formattedDate)
        await postUpdate(dataRespondentContact.value, dialogCode.value, value)
        break
      }
    case 'respondent-location':
      if (dataRespondentLocation.value !== dialogRespondentLocation.value) {
        await postUserLog(_employeeid.employeeid, 'respondent_info', 'respondent_location', dataContact.value, dialogContact.value, formattedDate)
        // postEditLog('respondent_info', 'respondent_location', dataRespondentLocation.value, dialogRespondentLocation.value, formattedDate)
        await postUpdate(dataRespondentLocation.value, dialogCode.value, value)
        break
      }
    default:
      break
  }
}

const refreshData = async () => {
  if (_isdemo.isdemo) fillupOffline()
  else {
    if (await getComplaintList2()) nodata.value = false
    else nodata.value = true
  }
}

;(async () => {
  checkOnline()
  await refreshData()
})()
</script>

<style lang="sass" scoped>

.table-loading
  font-size: 2rem
  margin-top: 3rem

.button-view
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
  // width: 100vw
  // height: 100vh

.dialog-card::-webkit-scrollbar
    display: none

.dialog-card__label
  font-family: 'Inter'
  font-weight: 300
  font-size: 1rem

.dialog-card__data
  font-family: 'Inter'
  font-weight: 500
  font-size: 1.2rem
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

.q-card--section
  padding: 2rem

.padded
  padding-bottom: 1rem


.select
  width: auto
  max-width: 200rem
  font-family: 'Inter'
  font-weight: 500
  font-size: 1.2rem
  text-align: center
  white-space: nowrap
  overflow: hidden
  text-overflow: ellipsis

.grid
  display: grid
  grid-template-columns: 1fr 1fr 1fr
  grid-template-rows: 0fr
  gap: 2rem 0px
  grid-template-areas: "code type received" "status status generate" "complaintant complaintant complaintant" "contact contact contact" "location location location" "details details details" "respondent respondent transaction" "respondent-contact respondent-contact respondent-contact" "respondent-location respondent-location respondent-location" "button button button"
  padding: 2rem

.code
  grid-area: code
  justify-self: start
  align-self: center
.type
  grid-area: type
.received
  grid-area: received
  justify-self: end
  align-self: center
.generate
  grid-area: generate
  justify-self: end
  align-self: center
  margin-right: 1rem
.status
  grid-area: status
  justify-self: start
  align-self: center
.complaintant
  grid-area: complaintant
  justify-self: start
  align-self: center
.contact
  grid-area: contact
  justify-self: start
  align-self: center
.location
  grid-area: location
  justify-self: start
  align-self: center
.details
  grid-area: details
  justify-self: start
  align-self: center
.respondent
  grid-area: respondent
  justify-self: start
  align-self: center
.transaction
  grid-area: transaction
  justify-self: end
  align-self: center
.respondent-contact
  grid-area: respondent-contact
  justify-self: start
  align-self: center
.respondent-location
  grid-area: respondent-location
  justify-self: start
  align-self: center
.button-area
  padding-top: 2rem
  grid-area: button


@media screen and (max-width: 500px)
  .section
    margin-bottom: 1.2rem

  .nodata--text
    font-family: 'Inter'
    font-weight: 300
    font-size: 1.4rem
    padding: 1rem 4.5rem
</style>
