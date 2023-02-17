<script lang="ts">
  import { Router, Link, Route } from "svelte-routing";
  import "./app.css";
  import Home from "./routes/Home.svelte";
  ///import Prueba from "./routes/Prueba.svelte";
  import ImagesOrganitation from "./routes/ImagesOrganitation.svelte";
  import FolderOrganitation from "./routes/FolderOrganitation.svelte";
  import QrGenerator from "./routes/QrGenerator.svelte";
  const list = [
    {
      title: "Home",
      url: "/",
      icon: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M6.827 6.175A2.31 2.31 0 015.186 7.23c-.38.054-.757.112-1.134.175C2.999 7.58 2.25 8.507 2.25 9.574V18a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18V9.574c0-1.067-.75-1.994-1.802-2.169a47.865 47.865 0 00-1.134-.175 2.31 2.31 0 01-1.64-1.055l-.822-1.316a2.192 2.192 0 00-1.736-1.039 48.774 48.774 0 00-5.232 0 2.192 2.192 0 00-1.736 1.039l-.821 1.316z" />
        <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 12.75a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0zM18.75 10.5h.008v.008h-.008V10.5z" />
      </svg>`,
      component: Home,
    },
    {
      title: "organizador",
      url: "/organizador",
      icon: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 4.875c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5A1.125 1.125 0 013.75 9.375v-4.5zM3.75 14.625c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5a1.125 1.125 0 01-1.125-1.125v-4.5zM13.5 4.875c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5A1.125 1.125 0 0113.5 9.375v-4.5z" />
        <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 6.75h.75v.75h-.75v-.75zM6.75 16.5h.75v.75h-.75v-.75zM16.5 6.75h.75v.75h-.75v-.75zM13.5 13.5h.75v.75h-.75v-.75zM13.5 19.5h.75v.75h-.75v-.75zM19.5 13.5h.75v.75h-.75v-.75zM19.5 19.5h.75v.75h-.75v-.75zM16.5 16.5h.75v.75h-.75v-.75z" />
      </svg>
      `,
      component: FolderOrganitation,
    },
    {
      title: "Organizacion Manual",
      url: "/ImagesManual",
      icon: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 4.875c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5A1.125 1.125 0 013.75 9.375v-4.5zM3.75 14.625c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5a1.125 1.125 0 01-1.125-1.125v-4.5zM13.5 4.875c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5A1.125 1.125 0 0113.5 9.375v-4.5z" />
        <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 6.75h.75v.75h-.75v-.75zM6.75 16.5h.75v.75h-.75v-.75zM16.5 6.75h.75v.75h-.75v-.75zM13.5 13.5h.75v.75h-.75v-.75zM13.5 19.5h.75v.75h-.75v-.75zM19.5 13.5h.75v.75h-.75v-.75zM19.5 19.5h.75v.75h-.75v-.75zM16.5 16.5h.75v.75h-.75v-.75z" />
      </svg>
      `,
      component: ImagesOrganitation,
    },
    {
      title: "Qr",
      url: "/Qr",
      icon: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 4.875c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5A1.125 1.125 0 013.75 9.375v-4.5zM3.75 14.625c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5a1.125 1.125 0 01-1.125-1.125v-4.5zM13.5 4.875c0-.621.504-1.125 1.125-1.125h4.5c.621 0 1.125.504 1.125 1.125v4.5c0 .621-.504 1.125-1.125 1.125h-4.5A1.125 1.125 0 0113.5 9.375v-4.5z" />
        <path stroke-linecap="round" stroke-linejoin="round" d="M6.75 6.75h.75v.75h-.75v-.75zM6.75 16.5h.75v.75h-.75v-.75zM16.5 6.75h.75v.75h-.75v-.75zM13.5 13.5h.75v.75h-.75v-.75zM13.5 19.5h.75v.75h-.75v-.75zM19.5 13.5h.75v.75h-.75v-.75zM19.5 19.5h.75v.75h-.75v-.75zM16.5 16.5h.75v.75h-.75v-.75z" />
      </svg>
      `,
      component: QrGenerator,
    },
  ];
  let width:number = 0;
  let click:boolean = true;
  export let url = "";
</script>

<svelte:window bind:innerWidth={width} />

<main class="relative h-screen w-screen">
  <Router {url}>
    {#if click}
      <nav
        class="absolute inset-y-0 left-0 h-full {width > 1090
          ? 'open-animation-w-2-12 w-2/12'
          : 'open-animation-w-3-12 w-3/12'} bg-red-700 p-1 "
      >
        <div class=" grid grid-cols-1 gap-1">
          <button
            class="transition duration-700 ease-in-out pl-2 pt-1 pb-1 w-10  hover:bg-yellow-500"
            on:click={() => (click = !click)}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M3.75 5.25h16.5m-16.5 4.5h16.5m-16.5 4.5h16.5m-16.5 4.5h16.5"
              />
            </svg>
          </button>
          {#each list as item, i}
            <Link
              class="bg-orange-700 px-5  py-6 ease-in duration-300 hover:bg-yellow-500 grid grid-cols-2 gap-0"
              to={item.url}
            >
              {@html item.icon}
              <span class="w-1 h-6">{item.title}</span></Link
            >
          {/each}
        </div>
      </nav>
    {:else}
      <nav
        class="absolute inset-y-0 left-0 h-screen w-1/12 {width > 1090
          ? 'close-animation-w-2-12 w-2/12'
          : 'close-animation-w-3-12 w-3/12'} bg-red-700 p-1"
      >
        <div class=" grid grid-cols-1 gap-1">
          <button
            class="transition duration-700 ease-in-out pl-2 pt-1 pb-1 w-10  hover:bg-yellow-500"
            on:click={() => (click = !click)}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M3.75 5.25h16.5m-16.5 4.5h16.5m-16.5 4.5h16.5m-16.5 4.5h16.5"
              />
            </svg>
          </button>
          {#each list as item, i}
            <Link
              class="bg-orange-700 px-5  py-6 ease-in duration-300 hover:bg-yellow-500 grid grid-cols-1 gap-0"
              to={item.url}
            >
              {@html item.icon}
            </Link>
          {/each}
        </div>
      </nav>
    {/if}
    <div
      class="absolute inset-y-0 right-0 object-right h-screen {width > 1090
        ? click
          ? 'w-10/12'
          : 'w-11/12'
        : click
        ? 'w-9/12'
        : 'w-11/12'}
        {width > 1090
           ? click
          ? 'pc-open-animation-w-2-12'
          : 'pc-close-animation-w-2-12'
        : click
        ? 'pc-open-animation-w-3-12'
        : 'pc-close-animation-w-3-12'}
        "
    >
      {#each list as item, i}
        <Route path={item.url} component={item.component} />
      {/each}
    </div>
  </Router>
</main>
