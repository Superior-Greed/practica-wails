<script>
  import { TaskComponent } from "../../class/task";
  
  export let task = new TaskComponent(
    "se a guardado con exito",
    3,
    "h-28 w-2/4 p-4 bg-lime-600 border-lime-400 border-2",
    "text-2xl text-white",
    true,
    1,
  );

  function Display() {
    //task.value = false;
    setTimeout(() => {
      task.value = true;
    }, task.time * 1000);
  }

  $:{
    if(!task.value){
        switch(task.type_task){
            case 1:
                task.class_div = "h-28 w-2/4 p-4 bg-lime-600 border-lime-400 border-2";
                Display();
                break;
            case 2:
                task.class_div = "h-28 w-2/4 p-4 bg-red-600 border-red-400 border-2";
                Display();
                break;
            case 3:
                task.class_div = "h-28 w-2/4 p-4 bg-yellow-600 border-yellow-400 border-2";
                Display()
                break;
        }
    }
  }
</script>
<div
hidden={task.value}
class="absolute top-2 right-0 h-20 w-2/4 {task.class_div}"
>
<p class="m-auto justify-center {task.class_p}">{task.text}</p>
</div>
<style></style>