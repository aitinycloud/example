<script setup>
import ifshow from './ifshow.vue'
</script>

<template>
<p3>模板标签</p3>
<p>{{ msg }}</p>
<p>{{ number + 1 }}</p>
<p>{{ message.split("").reverse().join("") }}</p>
<p v-html="rawHtml"></p>
<div v-bind:id="myid" :class="myclass"></div>
<div v-bind="myobject" ></div>

<header>
    <div>
      <ifshow />
    </div>
  </header>

</template>

<script>
export default{
  data(){
    return {
      msg : "This is a msg",
      number : 19,
      message: "Hello world",
      rawHtml: "<a href='https://www.baidu.com'>百度</a>",
      myid : "myid",
      myclass : "myclass",
      myobject : {
        mytmpid : "myid",
        mytmpclass : "myclass",
      },
    }
  }
}
</script>