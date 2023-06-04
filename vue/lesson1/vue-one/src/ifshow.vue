<template>
<p3>条件判断渲染</p3>
<div v-if="ifflag">if条件渲染你能看见我吗</div>
<div v-else>ifelse条件渲染你能看见我吗</div>
<div v-show="showflag">show条件渲染你能看见我吗</div>
<div v-for="item in names"> {{ item }} </div>
<div v-for="(value,key) in userInfo"> {{ key }} : {{ value }} </div>
<button v-on:click="addCount">按钮</button>
<div>{{ count }}</div>
</template>



<script>
export default {
    data(){
        return {
            showflag : false,
            ifflag : false,
            names: ["ABBD","DNFFF","FFWFWF"],
            userInfo: {
                name: "jmingzhang",
                age: "28",
                sex: "m",
            },
            count: 0,
        }
    },
    methods : {
        addCount(){
            this.count = this.count + 1;
        }
    }


}
</script>