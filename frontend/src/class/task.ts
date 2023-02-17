export class Task{
  
    text:string;
    time:number;
    class_div:string;
    class_p: string;
    value:boolean;

    constructor(_text:string,_time:number,_class_div:string,_class_p: string,_value:boolean){
      this.text = _text
      this.time = _time
      this.class_div = _class_div
      this.class_p = _class_p
      this.value = _value
    }
  }

export class TaskComponent{
  text:string;
  time:number;
  class_div:string;
  class_p: string;
  value:boolean;
  type_task:number;

  constructor(_text:string,_time:number,_class_div:string,_class_p: string,_value:boolean,_type_task:number){
    this.text = _text
    this.time = _time
    this.class_div = _class_div
    this.class_p = _class_p
    this.value = _value
    this.type_task = _type_task
  }
}