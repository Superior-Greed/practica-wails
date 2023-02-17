<script lang="ts">
  import Loading from "./generics/loading.svelte";
  import Qr from "./generics/Qr.svelte";
  import {
    ReturnImage,
    ImageUrl,
    ImageUrlList,
  } from "../../wailsjs/go/controller/Qr";

  let isloading = false;
  let content_text: string;
  var options_list = [];

  async function addQr() {
    if (content_text.trim() != "") {
      const request = await ReturnImage();
      const data = await request;
      options_list = [
        ...options_list,
        {
          backgroundImage: data,
          text: content_text,
          whiteMargin: false,
          size: 400,
        },
      ];
    }
  }
  /**
   * ejemplo de formato json list
   * {
  "data": [
    {
      "id": 0,
      "name": "Miller Madden"
    },
    {
      "id": 1,
      "name": "Mckee Wilkins"
    },
    {
      "id": 2,
      "name": "Hutchinson Oneal"
    },
    {
      "id": 3,
      "name": "Petersen Hendrix"
    },
    {
      "id": 4,
      "name": "Lesa Ochoa"
    }
  ],
  "images": {
    "url": [
  "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/batman-peliculas-orden-fotogramas-1616515164.jpg?crop=0.6914082358922217xw:1xh;center,top&resize=1200:*",
      "https://i.blogs.es/638a6a/the-batman-movie-review/840_560.jpg",
     "https://collectible506.com/wp-content/uploads/2019/09/BatmanDay2019_Blogroll-1_5d7c133dde7fb5.63394104.jpg",
     "https://i0.wp.com/codigoespagueti.com/wp-content/uploads/2022/02/DC-Comics-Cuanto-tiempo-lleva-activo-Bruce-Wayne-como-el-Caballero-de-la-Noche-en-The-Batman.png?fit=1280%2C720&quality=80&ssl=1",
      "https://m.media-amazon.com/images/M/MV5BMDdmMTBiNTYtMDIzNi00NGVlLWIzMDYtZTk3MTQ3NGQxZGEwXkEyXkFqcGdeQXVyMzMwOTU5MDk@._V1_FMjpg_UX1000_.jpg"
    ],
    "name": [
      "CandiceEspinoza.png",
      "IrwinParker.png",
      "DoreenPreston.png",
      "BryantHolmes.png",
      "LawrenceCallahan.png"
    ]
  }
}
  */

  /** 
   * json single
   * {
   * "data": "nose"
   * "image":{
   * "url":"https://i.blogs.es/638a6a/the-batman-movie-review/840_560.jpg",
   * "name":"batman.png"
   * }
   * }
   * **/

  async function AddJsonQr(is_list: boolean) {
    if (content_text.trim() != "") {
      if (is_list) {
        let json = JSON.parse(content_text);
        const request = await ImageUrlList(json.images.url, json.images.name);
        const data = await request;
        for (let index = 0; index < data.length; index++) {
          options_list = [
            ...options_list,
            {
              backgroundImage: data[index],
              text: JSON.stringify(json.data[index]),
              whiteMargin: false,
              size: 400,
            },
          ];
        }
      } else {
        let json = JSON.parse(content_text);
        const request = await ImageUrl(json.image.url, json.image.name);
        const data = await request;
        options_list = [
          ...options_list,
          {
            backgroundImage: data,
            text: JSON.stringify(json.data),
            whiteMargin: false,
            size: 400,
          },
        ];
      }
    }
  }
</script>

<Loading bind:isloading>
  <div class="w-full h-2/6 p-4">
    <textarea
      class="w-full h-4/5 p-8 text-black"
      name=""
      id=""
      cols="30"
      rows="10"
      bind:value={content_text}
    />
    <div class="h-1/5">
      <button
        on:click={() => addQr()}
        class="p-4 border-2 border-emerald-700 text-emerald-500"
        >Agregar imagen</button
      >
      <button
        on:click={() => AddJsonQr(false)}
        class="p-4 border-2 border-amber-700 text-amber-500"
        >Agregar Json</button
      >
      <button
        on:click={() => AddJsonQr(true)}
        class="p-4 border-2 border-amber-700 text-amber-500"
        >Agregar lista Json</button
      >
    </div>
  </div>
  <div class="w-full h-4/6 overflow-x-auto m-auto p-4">
    <div class="grid grid-cols-2 gap-4">
      {#each options_list as option}
        <Qr bind:option />
      {/each}
    </div>
  </div>
</Loading>
<style></style>
