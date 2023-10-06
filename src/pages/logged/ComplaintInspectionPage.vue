<template lang="pug">

q-page
  section.online
    q-icon(name="circle" size="1rem" :color="onlineColor")

  div.full-width.row.justify-between
    span.title Complaint - Add Inspection
    q-btn(flat size="md" label="Back" @click="gotoComplaintDashboard" icon="arrow_back").close-button

  div.container
    section.date
    section.sections
    section.structure-owner
    section.structure-owner-address
    section.lot-owner
    section.lot-owner-address
    section.telephone
    section.location-of-construction
    section.occupancy
    section.storey
    section.remarks

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

const gotoComplaintDashboard = () => {
  _pagewithtable.pagewithtable = false
  _currentpage.currentpage = 'complaint'
  router.push('/complaint')
}
</script>

<style lang="sass" scoped>
.container {  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-rows: 1fr 1fr 1fr 1fr 1fr 1fr 1fr 1fr 1fr;
  gap: 0px 0px;
  grid-auto-flow: row;
  grid-template-areas:
    "date structure-owner structure-owner"
    "sections structure-owner-address structure-owner-address"
    ". lot-owner lot-owner"
    ". lot-owner-address lot-owner-address"
    ". telephone telephone"
    ". location-of-construction location-of-construction"
    ". occupancy occupancy"
    ". storey storey"
    ". remarks remarks";
}

.date { grid-area: date; }

.sections { grid-area: sections; }

.structure-owner { grid-area: structure-owner; }

.structure-owner-address { grid-area: structure-owner-address; }

.lot-owner { grid-area: lot-owner; }

.lot-owner-address { grid-area: lot-owner-address; }

.telephone { grid-area: telephone; }

.location-of-construction { grid-area: location-of-construction; }

.occupancy { grid-area: occupancy; }

.storey { grid-area: storey; }

.remarks { grid-area: remarks; }

</style>
