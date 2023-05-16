<template lang="pug">

q-page(padding)
  div.flex.flex-center
    qrcode-stream(@decode="handleDecode")

</template>

<script setup lang="ts">
import { QrcodeStream } from 'vue-qrcode-reader'

const handleDecode = (text: string) => {
  console.log(text)
}
</script>
