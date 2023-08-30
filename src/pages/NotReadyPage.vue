<template lang="pug">

q-page(padding).flex.flex-center
  h2 This page is not ready
  component(:is="DocButton" label="aa")

</template>

<script setup lang="ts">
import DocButton from 'components/docButton.vue'
</script>
