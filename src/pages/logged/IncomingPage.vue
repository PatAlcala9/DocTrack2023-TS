<template lang="pug">
q-page(padding)
  div.full-width.row.justify-between
    span.title Incoming Commmunications - New Entry
    q-btn(flat size="md" label="Back" @click="gotoMenu" icon="arrow_back").close-button

  section.fit.row.items-center.justify-center
    div.column
      span.inputs__label Received Date:
      q-date(flat v-model="receivedDate" minimal color="$button" @click="sample").calendar
      span(v-if="formattedReceivedDate.length > 0").inputs__label--date {{ formattedReceivedDate }}
      span(v-else).inputs__label--date No Date Selected

    div.inputs.column
      span.inputs__label Source:
      //- input.inputs__input
      component(:is="docInput" width="40" alignment="left" transform="initial")

      span.inputs__label Subject:
      //- input.inputs__input
      component(:is="docInput" width="40" alignment="left" transform="initial")

      span.inputs__label Details:
      //- input.inputs__input
      component(:is="docTextArea" width="40")

      span.inputs__label Attachments:
      //- input.inputs__input
      component(:is="docTextArea" width="40")

  section.button-area.fit.row.items-center.justify-center
    component(:is="docButton" text="Save")

  section.list-area.column.text-center
    span.list-area__span Show List of Incomings
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { date } from 'quasar'
import { useRouter } from 'vue-router'

import docButton from 'components/docButton.vue'
import docTextArea from 'components/docTextArea.vue'
import docInput from 'components/docInput.vue'

let router = useRouter()
let receivedDate = ref('')
let formattedReceivedDate = ref('')

const sample = () => {
  formattedReceivedDate.value = date.formatDate(Date.parse(receivedDate.value), 'MMMM D, YYYY')
}

const gotoMenu = () => {
  router.push('/dashboard')
}
</script>

<style lang="sass" scoped>

.calendar
  background-color: rgba(8,25,35, 0.75)
  border: 1px solid rgba(255, 255, 255, 0.525)
  font-family: 'OpenSans'
  font-size: 1.6rem
  border-radius: 12px
  ::v-deep .q-date__calendar-item
    padding-top: 0.6rem
    // padding-left: 0.5rem
    // padding-right: 0.5rem
  ::v-deep .block
    font-size: 1.25rem




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

.dialog-card
  background-color: #021926

.dialog-content
  display: grid
  grid-template-columns: 1fr 1fr 1fr
  grid-template-rows: 1fr 1fr
  gap: 0px 0px
  grid-template-areas: "search search search" "table table table"
  // justify-items: end

.dialog-content-search
  font-family: 'Raleway'
  grid-area: search
  justify-self: stretch
  align-self: start
  height: 4em


.dialog-content-table
  // grid-area: table
  // justify-self: center
  // align-self: center
  margin-top: 2rem


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
</style>
