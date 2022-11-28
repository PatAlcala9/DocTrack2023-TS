<template lang="pug">

q-page(padding)
  //- div.full-width.row.justify-between
  //-   span.title Registration
  //-   q-btn(flat size="md" label="Back" @click="gotoHome" icon="arrow_back").close-button
  section.login.column.items-center
    span.login__username--label Fullname
    input(v-model="ifullname").login__username--input

    span.login__username--label Username
    input(v-model="iusername").login__username--input

    span.login__username--label Password
    input(v-model="ipassword" type="password").login__username--input

    span.login__username--label Access
    q-option-group(dark v-model="accessList" :options="accessOption" color="indigo-9" type="checkbox").login__username--option

    q-btn(outline padding="0.4rem 3rem" rounded color="white" label="Register" size="lg" @click="saveAccount")

q-dialog(v-model="dialog" persistent full-width full-height transition-show="scale" transition-hide="scale")
  q-card.dialog-card.text-white
    q-card-section
      div.dialog-title-area.row.justify-between
        span.dialog-title {{dialogMessage}}
        q-btn(flat size="md" label="close" v-close-popup).dialog-close

</template>

<script setup lang="ts">
import { api } from 'boot/axios'
import { ref } from 'vue'
import { encrypt } from 'src/js/OCBO'

let ifullname = ref('')
let iusername = ref('')
let ipassword = ref('')
let ipasswordEncrypted = ref('')
let accessList: any = ref([])
let accessOption = ref([
  {
    label: 'Incoming',
    value: 'is_incoming',
  },
  {
    label: 'Outgoing',
    value: 'is_outgoing',
  },
  {
    label: 'Releasing',
    value: 'is_releasing',
  },
  {
    label: 'Inventory',
    value: 'is_inventory',
  },
  {
    label: 'Other Documents',
    value: 'is_otherdocuments',
  },
])
let dialog = ref(false)
let dialogMessage = ref('')

const saveAccount = async () => {
  let ipasswordEncrypted = await encrypt(ipassword.value.toUpperCase())
  let iincoming = accessList.value.includes('is_incoming') ? 1 : 0
  let ioutgoing = accessList.value.includes('is_outgoing') ? 1 : 0
  let ireleasing = accessList.value.includes('is_releasing') ? 1 : 0
  let iinventory = accessList.value.includes('is_inventory') ? 1 : 0
  let iothers = accessList.value.includes('is_otherdocuments') ? 1 : 0
  const response = await api.post('/api/SaveAccount/' + ifullname.value.toUpperCase() + '/' + iusername.value.toUpperCase() + '/' + ipasswordEncrypted + '/' + iincoming + '/' + ioutgoing + '/' + ireleasing + '/' + iinventory + '/' + iothers)
  const data = response.data

  if (data.includes('Success')) {
    dialog.value = true
    dialogMessage.value = 'Successfully Registered'
  } else {
    dialog.value = true
    dialogMessage.value = 'Failed to Register'
  }
}
</script>

<style lang="sass" scoped>

.dialog
  font-family: 'Raleway'

.detail-dialog__info
  font-size: 2rem

.detail-dialog__info--large
  font-family: 'OpenSans'
  font-size: 3rem

.detail-dialog__info--subinfo
  font-size: 1.2rem

.login
  margin: 1rem
  font-family: 'Raleway'

.login__username--label
  font-family: 'Raleway'
  font-size: 1.4rem

.login__username--input
  font-family: 'OpenSans'
  font-size: 1.3rem
  border-radius: 0.6rem
  text-align: center
  margin-bottom: 1.4rem
  text-transform: uppercase

.login__username--option
  @extend .login__username--input
  text-transform: capitalize

.dialog-card
  background-color: blue
</style>
