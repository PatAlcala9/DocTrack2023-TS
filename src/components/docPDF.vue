<template lang="pug">

div
  q-btn(@click="createPDF" label="PDF")

</template>

<script setup lang="ts">
import { PDFDocument, StandardFonts, PageSizes } from 'pdf-lib'

export interface Props {
  title: string
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Generated Document',
})

const fileName = 'hello-world.pdf'

const createPDF = async () => {
  const pdfDoc = await PDFDocument.create()
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
  window.open(link.href, fileName)
}
</script>
