<template lang="pug">

div.full-width.column.wrap.justify-start.items-start.content-start
  div(v-for="version in versions" :key="version")
    br
    span.version {{ version }}
    span.info(v-for="(info, index) in infos[version]" :key="info" v-html="formatInfo(info, 0)")
</template>

<script setup lang="ts">
const versions = ['0.14', '0.13', '0.12', '0.11', '0.10', '0.9', '0.8', '0.7', '0.6', '0.5', '0.4', '0.3', '0.2', '0.1', 'Undocumented']
const infos = {
  '0.14': ['10/17/2023 - Adjusted UI', '10/17/2023 - Added Detection for fake QR', '10/17/2023 - Added new API for Backend to access QR Data', '10/16/2023 - Added data to be displayed', '10/16/2023 - Fixed display data for Scanned QR'],
  '0.13': ['10/12/2023 - Additional data in PDF', '10/11/2023 - Added complaint code to inspection data', '10/11/2023 - Fixed bug on saving inspection', '10/10/2023 - Added saving of Inspection', '10/09/2023 - Added dialogs on Inspection', '10/09/2023 - Sectionized Date input', '10/09/2023 - New APIs on Backend', '10/06/2023 - Grid design on UI', '10/06/2023 - Added Inspection Entry'],
  '0.12': ['10/05/2023 - Added counting of expiration on backend', '10/03/2023 - Finished Complaint highlighting when empty', '10/03/2023 - Fixed minor bug on backend', '09/29/2023 - Registration highlighted to red when empty', '09/29/2023 - Added color variables for fast color switching', '09/29/2023 - Complaint Add highlighted to red when empty (partially)'],
  '0.11': ['09/27/2023 - Added Highligh of expired data', '09/27/2023 - Revised complaint list filter', '09/26/2023 - Fixed Expiration countdown', '09/26/2023 - Added Pop-up when missing details', '09/25/2023 - Added recording of Employee Id', '09/25/2023 - Fixed lowercase value on saved data', '09/25/2023 - Changed checking of connection to 1 minute', '09/25/2023 - Added Edit log'],
  '0.10': ['09/22/2023 - Fixed Respondent saving', '09/18/2023 - Changed name to DDMS', '09/16/2023 - Added API for work stoppage', '09/16/2023 - Used the new logo for PDF'],
  '0.9': ['09/11/2023 - New Background', '09/11/2023 - Patched Security', '09/11/2023 - Removed Mobile Bug', '09/11/2023 - Changed mode to SPA', '09/11/2023 - Updated NPM', '09/11/2023 - Used new logo on README', '09/11/2023 - Fixed Browser detection and add loading when switching pages'],
  '0.8': ['09/07/2023 - Added individual fetching of complaint latest status for accuracy', '09/07/2023 - Fixed incorrect generation of complaint code', '09/07/2023 - Updated Backend Dependencies', '09/05/2023 - Improved way of getting complaint list'],
  '0.7': ['09/04/2023 - Fixed saving of complaint', '09/04/2023 - Redirected unready pages to temp file', "09/04/2023 - Fixed error when Complaint Type has '/'", '09/04/2023 - Simplified function to enhance performance'],
  '0.6': ['08/25/2023 - Fixed storage usage', '08/25/2023 - Updated Backend Dependencies', '08/25/2023 - Added a dialog to show version changes', '08/25/2023 - Added new Versioning component', '08/24/2023 - Fixed Demo mode authentication checking', '08/24/2023 - Changed text alignment', '08/24/2023 - Changed width measurement to percentage'],
  '0.5': ['08/23/2023 - Rearranged widescreen component', '08/23/2023 - Added Authentication Checking', '08/23/2023 - Imported Loading plugin', '08/23/2023 - Added loading animation', '08/23/2023 - Upgraded Backend from 1.20 to 1.21 for better performance', '08/23/2023 - New CheckConnection function'],
  '0.4': ['08/19/2023 - Each login needs authentication', '08/19/2023 - Minor format modification', '08/19/2023 - Added authentication encryption', '08/18/2023 - New Authentication storage', '08/18/2023 - Minor unnecessary code cleanup'],
  '0.3': ['08/18/2023 - Minor formatting fixes', '08/18/2023 - Added Request API If an interruption occurs or if the initial API request fails', '08/18/2023 - API request retry library added', '08/18/2023  - Increased retries to 3 if first attempt fails', '08/18/2023 - Minor changes to Backend'],
  '0.2': ['08/15/2023 - Added mouse hover animation', '08/15/2023 - Sample data included to demonstrate QR Scanner accuracy', "08/10/2023 - Fix API request problem with '/'", '08/10/2023 - Fix Invalid Complaint Display', '08/10/2023 - Fix query to list complaints', '08/10/2023 - Used new logo', '08/10/2023 - Updated Version and Depedencies'],
  '0.1': ['07/20/2023 - Refactor refresh as a function', '07/20/2023 - Added Status Saving when Dialog Closes', '07/20/2023 - Backend updated to fetch specific status', '07/12/2023 - Finished PDF Generation'],
  Undocumented: ['11/28/2022 - 07/07/2023 - Updates (No Description Changes)', '11/27/2022- Creation of DocTrack 2023 and Upload to temporary cloud host'],
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
