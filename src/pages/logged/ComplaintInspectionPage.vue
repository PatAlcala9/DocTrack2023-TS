<template lang="pug">

q-page(padding)
  section.online
    q-icon(name="circle" size="1rem" :color="onlineColor")

  div.full-width.row.justify-between
    span.title Complaint - Add Inspection
    q-btn(flat size="md" label="Back" @click="gotoComplaintDashboard" icon="arrow_back").close-button

  div.container
    div.left
      section.section.date
        span.form--label Date Received
        q-date(flat v-model="calendarDate" minimal color="$button" @click="formatDate").calendar
        component(v-if="docDate.length > 0" :is="docLabel" :text="docDate" ).form--label
        component(v-else :is="docLabel" text="No Date Selected").form--label

        section.sections
          component(:is="docList" text="Access" v-model:modelValue="sectionsList" :options="sectionsOption" :alert="redAccess" ).login__username--input

    div.right
      section.section.structure-owner
        component(:is="docLabel" text="Structure Owner:").label--spaced
        component(:is="docInputEntry" v-model:value="structureOwner" alignment="left" width=100)

      section.section.structure-owner-address
        component(:is="docLabel" text="Structure Owner Address:").label--spaced
        component(:is="docInputEntry" v-model:value="structureOwnerAddress" alignment="left" width=100)

      section.section.lot-owner
        component(:is="docLabel" text="Lot Owner:").label--spaced
        component(:is="docInputEntry" v-model:value="lotOwner" alignment="left" width=100)

      section.section.lot-owner-address
        component(:is="docLabel" text="Lot Owner Address:").label--spaced
        component(:is="docInputEntry" v-model:value="lotOwnerAddress" alignment="left" width=100)

      section.section.telephone
        component(:is="docLabel" text="Phone Number:").label--spaced
        component(:is="docInputEntry" v-model:value="phoneNo" alignment="left" width=100)

      section.section.location-of-construction
        component(:is="docLabel" text="Location of Construction:").label--spaced
        component(:is="docInputEntry" v-model:value="locationOfConstruction" alignment="left" width=100)

      section.section.occupancy
        component(:is="docLabel" text="Use of Occupancy:").label--spaced
        component(:is="docInputEntry" v-model:value="useOfOccupancy" alignment="left" width=100)

      section.section.storey
        component(:is="docLabel" text="Number of Story:").label--spaced
        component(:is="docInputEntry" v-model:value="noOfStorey" alignment="left" width=100)

  div.flex.flex-center.button-area
    component(:is="docButton" text="Save" @click="saveData")
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { date, useQuasar } from 'quasar'
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

const router = useRouter()
const quasar = useQuasar()
const _currentpage = useCurrentPage()
const _pagewithtable = usePageWithTable()
const _isdemo = useIsDemo()

let onlineColor = ref('')
let sectionsList = ref<string[]>([])
let sectionsOption = ref([
  {
    label: 'Administrative Fines',
    value: 'SECTION 212. Administrative Fines',
  },
  {
    label: 'Building Permits',
    value: 'SECTION 301. Building Permits',
  },
  {
    label: 'Certificate of Occupancy',
    value: 'SECTION 309. Certificate of Occupancy',
  },
  {
    label: 'Sizes and Dimensions of Courts',
    value: 'SECTION 804. Sizes and Dimensions of Courts',
  },
  {
    label: 'Windows Openings',
    value: 'SECTION 808. Windows Openings',
  },
  {
    label: 'Pedestrian Protection',
    value: 'SECTION 1106. Pedestrian Protection',
  },
  {
    label: 'Excavation, Foundations, and Retaining Walls',
    value: 'SECTION 1202. Excavation, Foundations, and Retaining Walls',
  },
])

let calendarDate = ref('')
let docDate = ref('')
let structureOwner = ref('')
let structureOwnerAddress = ref('')
let lotOwner = ref('')
let lotOwnerAddress = ref('')
let complaintName = ref('')
let phoneNo = ref('')
let locationOfConstruction = ref('')
let useOfOccupancy = ref('')
let noOfStorey = ref('')

const formatDate = () => {
  docDate.value = date.formatDate(Date.parse(calendarDate.value), 'MMMM D, YYYY')
}

const postInspectionSections = async (inspectionid: number, sectionid: number): Promise<boolean> => {
  const response = await api.post('/api/PostInspectionSections', {
    data: inspectionid,
    data2: sectionid
  })
  const data = response.data.length !== 0 ? response.data : null

  if (data.includes('Success')) return true
  else return false
}

const postInspection = async (structureOwner: string, soAddress: string, lotOwner: string, loAddress: string, phone: string, location: string, occupancy: string, storey: number) => {
  const response = await api.post('/api/PostInspection', {
    data: structureOwner,
    data2: soAddress,
    data3: lotOwner,
    data4: loAddress,
    data5: phone,
    data6: location,
    data7: occupancy,
    data8: storey
  })
  const data = response.data.length !== 0 ? response.data : null

  if (data.includes('Success')) return true
  else return false
}

const gotoComplaintDashboard = () => {
  _pagewithtable.pagewithtable = false
  _currentpage.currentpage = 'complaint'
  router.push('/complaint')
}
</script>

<style lang="sass" scoped>
.container
  display: grid
  grid-template-columns: 1fr 6fr
  grid-template-rows: 1fr
  gap: 0px 10px
  grid-auto-flow: row
  grid-template-areas: "left right"

.left
  display: grid
  grid-template-columns: 1fr
  grid-template-rows: 1fr
  grid-template-areas: "date"
  grid-area: left

.right
  display: grid
  grid-template-columns: 1fr
  grid-template-rows: 1fr 1fr
  gap: 0px 0px
  grid-template-areas: "structure-owner" "structure-owner-address" "lot-owner" "lot-owner-address" "telephone" "location-of-construction" "occupancy" "storey" "remarks"
  grid-area: right

.date
  grid-area: date
.sections
  grid-area: sections
.structure-owner
  grid-area: structure-owner
.structure-owner-address
  grid-area: structure-owner-address
.lot-owner
  grid-area: lot-owner
.lot-owner-address
  grid-area: lot-owner-address
.telephone
  grid-area: telephone
.location-of-construction
  grid-area: location-of-construction
.occupancy
  grid-area: occupancy
.storey
  grid-area: storey
.remarks
  grid-area: remarks

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

.section
  display: flex
  flex-direction: column
  // background-color: rgba(10, 10, 35, 0.52)
  // background-color: rgba(17, 25, 40, 0.82)
  //backdrop-filter: blur(1.6px) saturate(173%)
  //background-color: rgba(17, 25, 40, 0.8)
  background-color: rgba(12, 21, 42, 0.45)
  backdrop-filter: blur(9px) saturate(150%)
  border-radius: 0.6rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  padding: 1.2rem 2rem 1.2rem 1.2rem
  margin-bottom: 1rem

.button-area
  margin-top: 2rem

</style>
