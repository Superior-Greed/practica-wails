<script tag="ts">
  export let isloading = false;
</script>

<div class="container mx-auto h-screen">
  <slot />
  {#if isloading}
    <div
      class="flex border-4 loadin-animation rounded-full w-4/5 h-4/5 items-center justify-center m-auto"
    >
      <p class="text-9xl">loadin...</p>
    </div>
  {:else}
    <slot name="content" />
  {/if}
</div>

<style>
</style>
