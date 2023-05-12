<template lang="pug">

div
  qrcode-stream(@decode="onDecode")

</template>

<script setup lang="ts">
import { QrcodeStream } from 'vue-qrcode-reader'

const onDecode = () => {
  console.log()
}
</script>
