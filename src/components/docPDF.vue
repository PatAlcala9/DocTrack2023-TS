<template lang="pug">

div.flex.flex-center
  component(:is="docButton" @click="createPDF" text="Create Sample PDF")
  component(:is="docQR" :text="sampleQR" :size=qrSize)

</template>

<script setup lang="ts">
import { PDFDocument, StandardFonts, PageSizes, PDFField } from 'pdf-lib'
import { decryptAES, encryptAES } from 'src/js/OCBO'
import docQR from 'components/docQR.vue'
import docButton from 'components/docButton.vue'
import { ref } from 'vue'
// import htmlToImage from 'html-to-image'

export interface Props {
  title: string
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Generated Document',
})

const randomData1 = encryptAES(Math.random().toString()).substring(4, 10)
const randomData2 = encryptAES('some random string').substring(25)
const sampleQR = ref(encryptAES(`**SCAN ME USING OCBO DOCTRACK** QrId::${encryptAES(randomData1)}??${encryptAES(randomData2)}`))
const qrSize = 200

// htmlToImage.toPng(#qr)
// .then((dataUrl) => {
//   const link = document.createElement('a');
//   link.download = 'my-component.png'; // replace with your desired file name
//   link.href = dataUrl;
//   link.click();
// });

// const createRandomData = () => {

// }

const createPDF = async () => {
  const url = 'assets/wso.pdf'
  const existingPdfBytes = await fetch(url).then(res => res.arrayBuffer())
  const pdfDoc = await PDFDocument.load(existingPdfBytes)
  const page = pdfDoc.addPage(PageSizes.Letter)

  const font = await pdfDoc.embedFont(StandardFonts.Helvetica)
  const textSize = 14
  page.drawText('Hello World!', { x: 50, y: 50, size: textSize, font })
  page.drawText('Falcon Heavy', { x: 120, y: 560, size: textSize, font })
  page.drawText('Saturn IV', { x: 120, y: 500, size: textSize, font })
  page.drawText('Delta IV Heavy', { x: 340, y: 560, size: textSize, font })
  page.drawText('Space Launch System', { x: 340, y: 500, size: textSize, font })

  pdfDoc.setTitle(props.title)

  const pdfBytes = await pdfDoc.save()
  const pdfBlob = new Blob([pdfBytes], { type: 'application/pdf' })
  const link = document.createElement('a')
  link.href = URL.createObjectURL(pdfBlob)
  window.open(link.href)
}
</script>