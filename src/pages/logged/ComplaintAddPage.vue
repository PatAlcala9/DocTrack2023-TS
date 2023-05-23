<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Complaint - Add New
    q-btn(flat size="md" label="Back" @click="gotoComplaintDashboard" icon="arrow_back").close-button

  div.entry-group
    div.complaint-group
      section.fit.row.wrap.justify-start.items-center.content-center
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
        component(:is="docInputEntry" v-model="complaintName" )

      section.label-and-input
        component(:is="docLabel" text="Complaint Contact:").label--spaced
        component(:is="docInputEntry" v-model="complaintContact" )

      section.label-and-input
        component(:is="docLabel" text="Complaint Location:").label--spaced
        component(:is="docInputEntry" v-model="complaintLocation" )

      section.label-and-input
        component(:is="docLabel" text="Detail of Complaint:").label--spaced
        component(:is="docTextArea" v-model="complaintDetail" )

    div.respondent-group
      section.label-and-input
        component(:is="docLabel" text="Respondent Name:").label--spaced
        component(:is="docInputEntry" v-model="respondentName")

      section.label-and-input
        component(:is="docLabel" text="Respondent Contact:").label--spaced
        component(:is="docInputEntry" v-model="respondentContact")

      section.label-and-input
        component(:is="docLabel" text="Respondent Location:").label--spaced
        component(:is="docInputEntry" v-model="respondentLocation")
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { date } from 'quasar'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'

import { useCurrentPage } from 'stores/currentpage'

import docButton from 'components/docButton.vue'
import docTextArea from 'components/docTextArea.vue'
import docInputEntry from 'components/docInput.vue'
import docLabel from 'components/docLabel.vue'

let sourceEntryList = ref<Array<string>>([])
let sourceEntry = ref('Select Source')
let receivedDate = ref('')
let formattedReceivedDate = ref('')
let showCalendar = ref(false)
let complaintName = ref('')
let complaintContact = ref('')
let complaintLocation = ref('')
let complaintDetail = ref('')
let respondentName = ref('')
let respondentContact = ref('')
let respondentLocation = ref('')

const router = useRouter()
let _currentpage = useCurrentPage()

const formatDate = () => {
  formattedReceivedDate.value = date.formatDate(Date.parse(receivedDate.value), 'MMMM D, YYYY')
}

const getSourcesFromDatabase = async () => {
  const response = await api.get('/api/GetSources')
  const data = response.data.length !== 0 ? response.data : null

  if (data !== null) {
    sourceEntryList.value = data.result
  }

  console.log('sourceEntryList:', sourceEntryList.value[0])
}

const openCalendar = () => {
  dateReceived.value = ''
  showCalendar.value = true
}

const gotoMenu = () => {
  _currentpage.currentpage = 'dashboard'
  router.push('/dashboard')
}

const gotoComplaintDashboard = () => {
  _currentpage.currentpage = 'complaint'
  router.push('/complaint')
}

;(async () => {
  await getSourcesFromDatabase()
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
