<!-- <template>
  <div>
    <div ref="qrCodeRegion"></div>
    <div>Decoded Text: {{ decodedText }}</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import Html5Qrcode from 'html5-qrcode';

const qrCodeRegion = ref<HTMLElement | null>(null);
const decodedText = ref('');

onMounted(() => {
  const qrCodeScanner = new Html5Qrcode('qr-code-full-region');

  qrCodeScanner.start({ facingMode: 'environment' }, (decodedText) => {
    // Update the decoded text
    decodedText.value = decodedText;
  });

  // Set the qrCodeRegion ref to the QR code region element
  qrCodeRegion.value = document.getElementById('qr-code-full-region');
});

onUnmounted(() => {
  // Stop the QR code scanning when the component is unmounted
  const qrCodeScanner = new Html5Qrcode('qr-code-full-region');
  qrCodeScanner.stop();
});
</script> -->
