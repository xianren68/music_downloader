<template>
    <div class="PageNation">
        <div class="up" @click="changePage(false)">
            👆
        </div>
        <div class="line"></div>
        <div class="page">{{searchStore.page}}/{{Math.ceil(searchStore.total/20)}}</div>
        <div class="line"></div>
        <div class="down" @click="changePage(true)">
            👇
        </div>
    </div>
</template>

<script setup lang="ts">
import { throttle } from '@/utils'
import { SearchStore } from '@/store'
const searchStore = SearchStore()
const emit = defineEmits(['changePage'])
const changePage = throttle((type:boolean)=>{
    const value = searchStore.page
    if(type){
        searchStore.page = searchStore.page == Math.ceil(searchStore.total/20)?searchStore.page:searchStore.page+1
    }else {
        searchStore.page = searchStore.page == 1?searchStore.page:searchStore.page-1
    }
    if(value != searchStore.page){
        emit('changePage')
        console.log('cfdf')
    }
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