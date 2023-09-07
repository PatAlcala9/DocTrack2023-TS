<template lang="pug">

div.full-width.column.wrap.justify-start.items-start.content-start
  div(v-for="version in versions" :key="version")
    br
    span.version {{ version }}
    span.info(v-for="(info, index) in infos[version]" :key="info" v-html="formatInfo(info, 0)")
</template>

<script setup lang="ts">
const versions = ['0.8', '0.7', '0.6', '0.5', '0.4', '0.3', '0.2', '0.1', 'Undocumented']
const infos = {
  '0.8': ['09/07/2023 - Added individual fetching of complaint latest status for accuracy', '09/07/2023 - Fixed incorrect generation of complaint code', '09/07/2023 - Updated Backend Dependencies', '09/05/2023 - Improved way of getting complaint list'],
  '0.7': ['09/04/2023 - Fixed saving of complaint', '09/04/2023 - Redirected unready pages to temp file', "09/04/2023 - Fixed error when Complaint Type has '/'", '09/04/2023 - Simplified function to enhance performance'],
  '0.6': ['08/25/2023 - Fixed storage usage', "08/25/2023 - Updated Backend Dependencies", '08/25/2023 - Added a dialog to show version changes', '08/25/2023 - Added new Versioning component', '08/24/2023 - Fixed Demo mode authentication checking', '08/24/2023 - Changed text alignment', '08/24/2023 - Changed width measurement to percentage'],
  '0.5': ['08/23/2023 - Rearranged widescreen component', '08/23/2023 - Added Authentication Checking', '08/23/2023 - Imported Loading plugin', '08/23/2023 - Added loading animation', '08/23/2023 - Upgraded Backend from 1.20 to 1.21 for better performance', '08/23/2023 - New CheckConnection function'],
  '0.4': ['08/19/2023 - Each login needs authentication', '08/19/2023 - Minor format modification', '08/19/2023 - Added authentication encryption', '08/18/2023 - New Authentication storage', '08/18/2023 - Minor unnecessary code cleanup'],
  '0.3': ['08/18/2023 - Minor formatting fixes', '08/18/2023 - Added Request API If an interruption occurs or if the initial API request fails', '08/18/2023 - API request retry library added', '08/18/2023  - Increased retries to 3 if first attempt fails', '08/18/2023 - Minor changes to Backend'],
  '0.2': ['08/15/2023 - Added mouse hover animation', '08/15/2023 - Sample data included to demonstrate QR Scanner accuracy', "08/10/2023 - Fix API request problem with '/'", '08/10/2023 - Fix Invalid Complaint Display', '08/10/2023 - Fix query to list complaints', '08/10/2023 - Used new logo', '08/10/2023 - Updated Version and Depedencies'],
  '0.1': ['07/20/2023 - Refactor refresh as a function', '07/20/2023 - Added Status Saving when Dialog Closes', '07/20/2023 - Backend updated to fetch specific status', '07/12/2023 - Finished PDF Generation'],
  Undocumented: ['07/07/2023-11/28/2023 - Update (Undocumented Changes)', '11/27/2023 - Creation of DocTrack 2023 and Upload to temporary cloud host'],
}
const formatInfo = (info: string, isFirst: number) => {
  if (isFirst) {
    return info
  } else {
    return '<br>* ' + info
  }
}
</script>

<style lang="sass">
.version
  font-size: 2rem

.info
  font-size: 1.2rem
</style>
