<script lang="ts">
  import {
    Carpet,
    AddFolderImage,
    TransferImages,
    AllFolderImage,
  } from "../../wailsjs/go/controller/Image";
  import LoadingStop from "../routes/generics/loading.svelte";
  import { models } from "../../wailsjs/go/models";
  import { TaskComponent } from "../class/task";
  import Tasks from "../routes/generics/task.svelte";

  let loading: boolean = false;
  let folder_inicial: string = "";
  let folder_final: string = "";
  let folder_save: models.Folder = new models.Folder();
  let list_folder: models.Folder[];
  let name: string = "";
  let description: string = "";
  let rute: string = "";
  let termination_image: string = "";
  let width;
  let view_list = true;

  let task = new TaskComponent(
    "se a guardado con exito",
    3,
    "h-28 w-2/4 p-4 bg-lime-600 border-lime-400 border-2",
    "text-2xl text-white",
    true,
    1,
  );

  AllFolderImage().then((data) => {
    list_folder = data;
    console.log(list_folder);
  });

  function RuteFinal(option: boolean) {
    Carpet().then((result) => {
      if (option) {
        folder_inicial = result;
      } else {
        folder_final = result;
      }
    });
  }

  async function Transfer() {
    const response = await TransferImages(folder_inicial, folder_final);
    const data = await response;
    console.log(data);
  }

  async function SaveFolder() {
    if (
      name.trim() != "" &&
      description.trim() != "" &&
      rute.trim() != "" &&
      termination_image.trim() != ""
    ) {
      folder_save.name = name;
      folder_save.description = description;
      folder_save.rute = rute.toUpperCase().trim().split(" ").join("");
      folder_save.termination_image = (
        termination_image.trim().split(" ").join("_") + "_"
      ).toLowerCase();
      if (
        list_folder.filter(
          (x) =>
            x.termination_image == folder_save.termination_image ||
            x.rute == folder_save.rute
        ).length == 0
      ) {
        const response = await AddFolderImage(folder_save);
        const data = await response;
        folder_save = data;
        list_folder.push(folder_save);
        name = description = rute = termination_image = "";

        task.type_task = 1;
        task.text = "folder guardado con exito";
      } else {
        task.type_task = 2;
        task.text = "uno de los valores existe";
      }
    } else {
      task.type_task = 2;
      task.text = "llene todos los campos";
    }
    task.value =false
  }
</script>

<svelte:window bind:innerWidth={width} />
<LoadingStop bind:isloading={loading}>
  <div class="w-full h-1/6 flex">
    <div class="w-3/6 m-2 pr-2">
      <button
        on:click={() => RuteFinal(true)}
        class=" p-2 rounded-md border-2 bg-blue-300 hover:bg-blue-200 border-blue-500 hover:border-blue-300 text-blue-700 hover:text-blue-900 "
        >Folder inicial</button
      >
      <button
        on:click={() => RuteFinal(false)}
        class=" p-2 rounded-md border-2 bg-blue-300 hover:bg-blue-200 border-blue-500 hover:border-blue-300 text-blue-700 hover:text-blue-900"
        >Folder destino</button
      >

      <button
        on:click={() => Transfer()}
        class=" p-2 rounded-md border-2 bg-blue-300 hover:bg-blue-200 border-blue-500 hover:border-blue-300 text-blue-700 hover:text-blue-900"
        >Guardar</button
      >
    </div>
    <div
      class="w-2/6 h-20 m-2 p-2 rounded-md border-2 border-blue-500 text-blue-700 overflow-x-auto text-left"
    >
      <p class="">
        <label for="">Inicial: </label>
        {folder_inicial}
      </p>
      <p class="border-t-2 mt-2 w-full border-t-yellow-400">
        <label for="">Final: </label>
        {folder_final}
      </p>
    </div>
    <div
      class="w-1/6 h-20 m-2 p-2 rounded-md border-2 border-blue-500 text-blue-700 overflow-x-auto text-left"
    >
      <input
        checked
        on:change={(e) => {
          view_list = e.target.checked;
        }}
        type="checkbox"
        name=""
        id=""
      /> Abrir formulario de creacion
    </div>
    <Tasks bind:task />
  </div>
  {#if view_list}
    <div
      class="w-full  {width < 1000
        ? 'h-4/6'
        : 'h-3/6 '} p-2 text-black grid grid-rows-4 grid-flow-col gap-2"
    >
      <div class="row-start-1 row-span-4 h-full w-full">
        <label
          class=" uppercase tracking-wide text-green-700 text-xs font-bold"
          for="grid-first-name"
        >
          Descripcion
        </label>
        <textarea
          placeholder="Descripcion"
          class=" w-full py-10 rounded-md max-h-full"
          name=""
          id=""
          cols="30"
          rows="10"
          bind:value={description}
        />
      </div>
      <div class="h-full w-full">
        <label
          class=" uppercase tracking-wide text-green-700 text-xs font-bold"
          for="grid-city"
        >
          Nombre
        </label>
        <input
          type="text"
          placeholder="Nombre"
          class="py-4 w-full rounded-md"
          bind:value={name}
        />
      </div>
      <div class="h-full w-full ">
        <label
          class=" uppercase tracking-wide text-green-700 text-xs font-bold"
          for="grid-state"
        >
          Ruta
        </label>
        <input
          type="text"
          placeholder="Nombre de Ruta se unira todo el texto final"
          class="py-4 w-full rounded-md"
          bind:value={rute}
        />
      </div>
      <div class="h-full w-full ">
        <label
          class=" uppercase tracking-wide text-green-700 text-xs font-bold"
          for="grid-zip"
        >
          Nombre formato de imagen
        </label>
        <input
          type="text"
          placeholder="Nombre formato cada espacio agrega un _"
          class="py-4 w-full rounded-md"
          bind:value={termination_image}
        />
      </div>
      <div class="h-full w-full relative py-1">
        <button
          on:click={() => SaveFolder()}
          class="float-right  py-4 w-4/5 border-2 hover:border-4 bg-emerald-100 hover:border-emerald-500 border-emerald-400 text-emerald-700 rounded-md "
          >Guardar</button
        >
      </div>
    </div>
  {:else}
    <div class="w-full h-5/6 pt-4 overflow-x-auto">
      {#each list_folder as folder}
        <div class="grid grid-cols-3 gap-4 p-4 m-2 border-4 border-amber-600">
          <div>
            <p><label for="">Nombre: </label>{folder.name}</p>
          </div>
          <div>
            <p><label for="">Carpeta: </label>/{folder.rute}</p>
          </div>
          <div>
            <p>
              <label for=""
                >Nombre que deve conterner la imagen:
              </label>{folder.termination_image}image.png
            </p>
          </div>
          <div class="col-span-3">
            <p>
              <label for=""
                >Descripcion de la carpeta:
              </label>{folder.description}
            </p>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</LoadingStop>

<style>
</style>
