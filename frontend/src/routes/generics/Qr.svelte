<script lang="ts">
    import { AwesomeQR} from "awesome-qr";
    let qrCode;
    export let option;
  
    async function QrGenerator() {
    new AwesomeQR(option)
        .draw()
        .then((dataURL) => (qrCode = dataURL));
    }
    $: {
      QrGenerator();
    }
  </script>
  
    <img src={qrCode} alt="" />
  
  <style></style>
  