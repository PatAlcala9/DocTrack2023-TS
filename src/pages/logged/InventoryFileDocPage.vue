<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title File List
    q-btn(flat size="md" label="Back" @click="gotoInventory" icon="arrow_back").close-button

  div.fit.row.wrap.justify-start.items-start.content-start.tab-area
    section(v-if="tabSelected === 'incoming'")
      span.tab__selected Incoming
      span(@click="switchTab('incoming')").tab Outgoing
    section(v-else)
      span(@click="switchTab('outgoing')").tab Incoming
      span.tab__selected Outgoing
  //-   section.column
  //-     span Search
  //-     q-input(dense filled v-model="searchValue" bg-color="white" :input-style="{ fontSize: '1.2rem', fontFamily: 'OpenSans' }" @keydown.enter="getIncomingUsingValue")

  div.flex.flex-center
    section(v-if="tabSelected === 'incoming'").dialog-content-table
      table.table
        thead
          tr
            th Entry Code
            th Source
            th Subject
            th Date Received
            th Add Data
        tbody
          tr(v-for="(item, index) in incomingList.result" :key="item").table-content-group
            td {{item}}
            td {{incomingList.result2[index]}}
            td {{incomingList.result3[index]}}
            td {{incomingList.result4[index]}}
            td
              q-btn(color="button" icon="visibility" :ripple="false").button-view
              //- q-btn(v-if="showText === false" v-else color="button" label="View" :ripple="false" @mouseleave="mouseLeave").button-view

    section(v-else).dialog-content-table
      table.table
        thead
          tr
            th Reference Number
            th Respondent
            th Subject
            th Date Released
            th Add Data
        tbody
          tr(v-for="(item, index) in outgoingList.result" :key="item").table-content-group
            td {{item}}
            td {{outgoingList.result2[index]}}
            td {{outgoingList.result3[index]}}
            td {{outgoingList.result4[index]}}
            td {{outgoingList.result5[index]}}
            td
              q-btn(color="button" icon="visibility" :ripple="false" @click="openDetails(item, outgoingList.result2[index], outgoingList.result3[index], outgoingList.result4[index])").button-view
              //- q-btn(v-if="showText === false" v-else color="button" label="View" :ripple="false" @mouseleave="mouseLeave").button-view



</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { api } from 'boot/axios'
import { useCurrentPage } from 'stores/currentpage'
import { usePageWithTable } from 'stores/pagewithtable'

import docButton from 'components/docButton.vue'
import docLabel from 'components/docLabel.vue'

const _pagewithtable = usePageWithTable()
const _currentpage = useCurrentPage()
const router = useRouter()

let tabSelected = ref('incoming')

type Incoming = {
  result: string
  result2: string
  result3: string
  result4: string
}
let incomingList = ref({} as Incoming)

type Outgoing = {
  result: string
  result2: string
  result3: string
  result4: string
  result5: string
}
let outgoingList = ref({} as Outgoing)

const switchTab = (value: string) => {
  if (value === 'incoming') tabSelected.value = 'outgoing'
  else tabSelected.value = 'incoming'
}

const gotoInventory = () => {
  _pagewithtable.pagewithtable = false
  _currentpage.currentpage = 'inventory'
  router.push('/inventory')
}
</script>

<style lang="sass" scoped>
.title
  grid-area: title
  justify-self: start
  align-self: stretch
  padding: 0 0 0 1rem
  margin: 0 0 -2rem 0

.logo
  width: 12rem
  height: auto
  opacity: 0.9
  margin-right: 2rem

.name
  grid-area: name
  justify-self: start
  align-self: center

.login
  grid-area: login
  justify-self: center
  align-self: center
  padding: 1.4rem
  border: 1px solid rgba(255, 255, 255, 0.125)
  border-radius: 2rem
  background-color: #021926
  margin-top: -10rem
  // backdrop-filter: blur(16px) saturate(180%)

.inquiry
  grid-area: inquiry
  justify-self: start
  align-self: center

.inquiry-text
  font-family: 'Raleway'
  font-size: 1.2rem
  text-decoration: underline
  padding: 2rem
  cursor: pointer
  color: #ffffff

.username-area
  margin: 1rem

.username-label
  font-family: 'Raleway'
  font-size: 1.4rem
  color: #ffffff

.username-input
  font-family: 'Raleway'
  font-size: 1.3rem
  border-radius: 0.6rem
  text-align: center
  text-transform: uppercase

.password-area
  margin: 1rem

.password-label
  @extend .username-label

.password-input
  @extend .username-input

.button-area
  padding-top: 1rem

.davao
  grid-area: davao
  justify-self: end
  align-self: center

.davaologo
  width: 18rem
  height: auto
  opacity: 0.4

.inquiry-label
  font-family: 'Raleway'
  font-size: 2.1rem
  color: #ffffff

.dialog
  width: 100vw

.dialog-title-area
  padding: 0.1rem 0  1rem 0.1rem

.dialog-title
  font-family: 'Raleway'
  font-size: 1.4rem

.dialog-subtitle
  font-family: 'Raleway'
  font-size: 1.2rem

.inputs
  margin-left: 2rem

.inputs__label
  font-size: 1.2rem

.inputs__input
  font-family: 'OpenSans'
  font-size: 1.3rem
  border-radius: 0.6rem
  // text-align: center
  margin-bottom: 1rem
  width: 50vw

.inputs__label--date
  @extend .inputs__label
  font-family: 'OpenSans'
  margin-top: 0.5rem
  font-size: 1.4rem

.button-area
  margin-top: 2rem

.list-area
  padding-top: 2rem
  font-size: 0.5rem

.list-area__span
  cursor: pointer
  text-decoration: underline
  font-size: 1.2rem

.form-area
  margin-top: 2rem

.requiredWarning
  color: red

.buttons-area
  margin: 2rem 0

.button
  margin-left: 1rem

.table
  font-family: 'Montserrat'
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
  padding: 1rem 1rem
  width: 20rem

.table tbody tr
  border-bottom: 1px solid #dddddd

.table tbody tr:nth-of-type(even)
  background-color: #1C4157

.table tbody tr:last-of-type
  border-bottom: 2px solid #000406

.tab-area
  margin-top: 2rem

.tab
  font-family: 'Raleway'
  font-size: 1.2rem
  padding: 1rem
  border: 1px solid white
  cursor: pointer

.tab__selected
  @extend .tab
  font-family: 'RalewayBold'
  background-color: rgba(255, 255, 255, 0.1)
</style>
