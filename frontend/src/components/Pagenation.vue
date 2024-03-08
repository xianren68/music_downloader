<template>
    <div class="PageNation">
        <div class="up" @click="changePage(current!-1)">
            👆
        </div>
        <div class="line"></div>
        <div class="page">{{current}}/{{Math.ceil(total/20)}}</div>
        <div class="line"></div>
        <div class="down" @click="changePage(current!+1)">
            👇
        </div>
    </div>
</template>

<script setup lang="ts">
import { throttle } from '@/utils'
defineProps({total:Number,current:Number})
const emit = defineEmits(['changePage'])
const changePage = throttle((page:number)=>{
    emit('changePage',page)
},500)
</script>

<style scoped lang="scss">
.PageNation{
    height: 90%;
    margin-left: 20px;
    margin-top:3%;
    width:40px;
    display: flex;
    flex-direction: column;
    align-items: center;
    .up,.down,.page {
        height: 40px;
        display: flex;
        align-items: center;
    }
    .up,.down{
        font-size: 30px;
        cursor: pointer;
    }
    .line{
        height: calc((100% - 120px)/2);
        width: 2px;
        background-color: aqua;
        opacity: 20%;
    }
}
</style>