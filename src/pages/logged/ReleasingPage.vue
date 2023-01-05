<template lang="pug">

q-page(padding)
  div.full-width.row.justify-between
    span.title Releasing
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

  div.fit.row.wrap.justify-start.items-start.content-start.search-area
    section.column
      span Search
      q-input(dense filled v-model="searchValue" bg-color="white" :input-style="{ fontSize: '1.2rem', fontFamily: 'OpenSans' }" @keydown.enter="getIncomingUsingValue")
    section.column
      span Search By
      //- input
      q-btn-dropdown(unelevated color="grey-10" :label="searchByValue" size="1rem" :content-style="{ fontSize: '1.2rem', fontFamily: 'Raleway' }")
        q-list
          q-item(clickable v-close-popup @click="searchByCode" @keypress="formatApplication($event)")
            q-item-section
              q-item-label ENTRY CODE

          q-item(clickable v-close-popup @click="searchByValue = 'DATE'")
            q-item-section
              q-item-label DATE

          q-item(clickable v-close-popup @click="searchByValue = 'NAME'")
            q-item-section
              q-item-label NAME

          q-item(clickable v-close-popup @click="searchByValue = 'SUBJECT'")
            q-item-section
              q-item-label SUBJECT

  div.flex.flex-center
    section(v-if="releasingList.result !== ''").dialog-content-table
      table.table
        thead
          tr
            th Entry Code
            th Received Date
            th Name
            th Subject
            th Details
        tbody
          tr(v-for="(item, index) in releasingList.result" :key="item").table-content-group
            td {{item}}
            td {{releasingList.result2[index]}}
            td {{releasingList.result3[index]}}
            td {{releasingList.result4[index]}}
            td
              q-btn(color="button" icon="visibility" :ripple="false" @click="openDetails(item, incomingList.result2[index], incomingList.result3[index], incomingList.result4[index])").button-view
              //- q-btn(v-if="showText === false" v-else color="button" label="View" :ripple="false" @mouseleave="mouseLeave").button-view

    section(v-else).table-loading.column.items-center
      span Loading Contents
      q-spinner-orbit(color="white" size="4em" style="margin-top: 2rem")

</template>

<script setup lang="ts">
import { ref } from 'vue'
import { api } from 'boot/axios'
import { useRouter } from 'vue-router'

import docButton from 'components/docButton.vue'

const router = useRouter()

let searchValue = ref('')
let searchByValue = ref('')

type Releasing = {
  result: string
  result2: string
  result3: string
  result4: string
}
let releasingList = ref({} as Releasing)
let showText = ref(false)

const gotoMenu = () => {
  router.push('/dashboard')
}

let details = ref(false)
let entryCodeDetail = ref('')
let dateReceivedDetail = ref('')
let nameDetail = ref('')
let subjectDetail = ref('')
let detailsDetail = ref('')
let attachmentDetail = ref('')
let noteDetail = ref('')

const openDetails = async (entryCode: string, date: string, sourcename: string, subjectdetail: string) => {
  details.value = true
  entryCodeDetail.value = entryCode
  dateReceivedDetail.value = date
  nameDetail.value = sourcename
  subjectDetail.value = subjectdetail

  await getIncomingDetails()
  await getIncomingDocLog()
  await getIncomingActionLog()
}
</script>

<style lang="sass" scoped>
// .calendar
//   background-color: rgba(8,25,35, 0.75)
//   border: 1px solid rgba(255, 255, 255, 0.525)
//   font-family: 'OpenSans'
//   font-size: 1.6rem
//   border-radius: 12px
//   ::v-deep .q-date__calendar-item
//     padding-top: 0.6rem
//     // padding-left: 0.5rem
//     // padding-right: 0.5rem
//   ::v-deep .block
//     font-size: 1.25rem

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
</style>
