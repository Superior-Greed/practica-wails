<script lang="ts">
  import { Task } from "../class/task";
  import type {controller} from "../../wailsjs/go/models";

  import {
    Images,
    Carpet,
    TransferImageToFolder,
  } from "../../wailsjs/go/controller/Image.js";
  import LoadingStop from "./generics/loading.svelte";

  let imagenes_list: controller.ImagesRouteBase64[] = [];
  let is_final_image = false;
  let carpet: string = "";

  let task = new Task(
    "se a guardado con exito",
    3,
    "h-28 w-2/4 p-4 bg-lime-600 border-lime-400 border-2",
    "text-2xl text-white",
    true
  );

  let options = [
    {
      value: 0,
      carpet: "Comedia",
      class: " bg-amber-600 border-amber-900 hover:bg-amber-400 ",
    },
    {
      value: 1,
      carpet: "Miedo",
      class: "  bg-neutral-600 border-neutral-900 hover:bg-neutral-400 ",
    },
    {
      value: 2,
      carpet: "Sexy",
      class: " bg-pink-600 border-pink-900 hover:bg-pink-400 ",
    },
    {
      value: 3,
      carpet: "Alegria",
      class: " bg-emerald-600 border-emerald-900 hover:bg-emerald-400 ",
    },
    {
      value: 4,
      carpet: "Triste",
      class: " bg-cyan-600 border-cyan-900 hover:bg-cyan-400 ",
    },
    {
      value: 5,
      carpet: "Idefinido",
      class: " bg-indigo-600 border-indigo-900 hover:bg-indigo-400 ",
    },
  ];

  //trae la informacion basica de la imagenes y previsualiza
  function ImagesFiles() {
    imagenes_list = [];
    is_final_image = true;
    Images().then((result) => {
      imagenes_list = result;
      is_final_image = false;
    });
  }

  //trae la ruta donde se depositaran las imagenes
  function RuteFinal() {
    carpet = "";
    Carpet().then((result) => {
      carpet = result;
    });
  }

  //muestra el toast artificial
  async function Display(text: string, resp: boolean) {
    task.value = true;
    if (resp) {
      task.text = text;
      task.class_div = "h-28 w-2/4 p-4 bg-lime-600 border-lime-400 border-2";
    } else {
      task.class_div = "h-28 w-2/4 p-4 bg-red-600 border-red-400 border-2";
      task.text = text;
    }
    task.value = false;
    setTimeout(() => {
      task.value = true;
    }, task.time * 1000);
  }
  //organiza imagen por imagen y crea el folder
  async function Folder(namefolder: object, direction: string, rute: string) {
    imagenes_list = imagenes_list.filter((x) => x.url != direction);
    if ((direction.trim() != "", rute.trim() != "")) {
      const r = await TransferImageToFolder(namefolder.carpet, direction, rute);
      console.log(r);
      const data = await r;
      console.log(data);
      Display(data.text,data.value)
    }
  }
</script>

<LoadingStop bind:isloading={is_final_image}>
  <div class="h-1/5 w-full p-4 flex ">
    <button
      class="border-2 rounded-md hover:bg-sky-400 border-blue-600  hover:border-blue-400 hover:text-blue-700 text-blue-300 w-32 h-20 mr-4"
      on:click={ImagesFiles}>imagenes</button
    >
    <button
      class="border-2 rounded-md hover:bg-sky-400 border-blue-600  hover:border-blue-400 hover:text-blue-700 text-blue-300 w-32 h-20 mr-4"
      on:click={RuteFinal}>ruta final</button
    >
    <p
      class="border-4 rounded-md border-blue-600 w-full h-20 text-blue-300 pt-5"
    >
      <label for="">Ruta final: </label>{carpet}
    </p>

    <div
      hidden={task.value}
      class="absolute top-2 right-0 h-20 w-2/4 {task.class_div}"
    >
      <p class="m-auto justify-center {task.class_p}">{task.text}</p>
    </div>
  </div>
  <div
    slot="content"
    class="snap-x snap-mandatory h-4/5 w-full overflow-scroll "
  >
    {#each imagenes_list as file}
      <div class="snap-center h-full shadow-2xl pt-4">
        <div class="flex h-full w-full pb-6 pl-1 pr-1">
          <img
            loading="lazy"
            class="border-4 border-red-600 h-full w-8/12 m-1 "
            src={file.url_server}
            alt="d"
          />
          <div
            class="m-1 p-1 overflow-y-auto h-full w-4/12 grid grid-cols-1 grid-rows-6 gap-1 border-2 border-amber-400"
          >
            {#each options as option}
              <button
                on:click={Folder(option, file.url, carpet)}
                disabled={carpet.trim() === ""}
                class="px-4 py2 {option.class} border-2 "
              >
                {option.carpet}
              </button>
            {/each}
          </div>
        </div>
      </div>
    {/each}
  </div>
</LoadingStop>

<style>
</style>
