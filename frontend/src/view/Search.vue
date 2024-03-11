<template>
    <div class="search">
        <div class="searchbox">
                <input type="text" placeholder="请输入您要搜索的内容..." v-model="searchText">
                <button @click="search">搜索</button>
        </div>
        <div class="result">
            <div class="list">
                <div class="head" v-if="searchStore.ResultList.length != 0">
                    <div class="name">名称</div>
                    <div class="author">作者</div>
                    <div class="albums">专辑</div>
                    <div class="download"></div>
                </div>
                <div class="item" v-for="(i,index) in searchStore.ResultList" :key="index">
                    <div class="name">{{i.NAME}}</div>
                    <div class="author">{{i.ARTIST}}</div>
                    <div class="albums">{{i.ALBUM}}</div>
                    <div class="download"><button @click="e=>download(index)">下载</button></div>
                </div>
            </div>
            <Pagenation :total="total" :current="pageNo" @changePage="changePage" v-if="searchStore.ResultList.length != 0"></Pagenation>
        </div>
    </div>
</template>

<script setup lang="ts">
import {SearchResult} from '@/type'
import { Search,Download } from '@/../wailsjs/go/main/App'
import {ref,onBeforeUnmount} from 'vue'
import { ElMessage } from 'element-plus'
import {SearchStore,DownStore} from '@/store'
import Pagenation from '@/components/Pagenation.vue'
const total = ref(0)
const pageNo = ref(1)
const searchStore = SearchStore()
const downloadStore = DownStore()
const searchText = ref(searchStore.searchText)
// 发送请求
const search = async () => {
    if(searchText.value === ""){
        return
    }
    // 前端发送有跨域问题，到后端去请求
    const res = await Search(pageNo.value,searchText.value)
    if(res === ""){
        ElMessage.error("请求出错")
        return
    }
    const result:SearchResult = JSON.parse(res)
    if(result.abslist.length === 0){
        ElMessage.info("没有找到相关歌曲")
    }
    total.value = +result.TOTAL
    searchStore.ResultList = result.abslist
}
// 切换页码
const changePage = (value:number) => {
    if(value < 1 || value > Math.ceil(total.value / 20)){
        return
    }
    pageNo.value = value
    search()
}
// 下载
const download = (index:number) => {
    const info = searchStore.ResultList![index]
    // 从队列中去除
    searchStore.ResultList.splice(index,1)
    Download(info.NAME!,info.ARTIST!,info.MUSICRID!.split("_")[1],info.ALBUM!)
    const obj = {
        id:info.MUSICRID!.split("_")[1]!,
        name:info.NAME!,
        album:info.ALBUM!,
        isStart:false,
        author:info.AARTIST!,
        progress:0,
        
    }
    // 添加到下载队列
    downloadStore.downList.set(obj.id,obj)
    downloadStore.size++
}
onBeforeUnmount(() => {
    searchStore.searchText = searchText.value
    console.log(downloadStore.downList)
})

</script>

<style scoped lang="scss">
.search {
    height: 100%;
    width:calc(100% - 60px);
    padding-top:1px;
    display: flex;
    flex-direction: column;
    .searchbox{
        margin-top: 50px;
        margin-left:90px;
        display: flex;
        align-items: center;
        input{
            box-sizing: border-box;
            width: 400px;
            height: 30px;
            border: 1px solid #ccc;
            border-right: none;
            border-radius: 5px 0 0 5px;
            padding-left: 10px;
        }
        button{
            height: 30px;
            box-sizing: border-box;
            width:80px;
            text-align: center;
            border:1px solid #ccc;
            background-color:rgba(234,233,235,.6);
            border-radius:0 5px 5px 0;
            filter: blur(.3px);
        }

    }
    .result{
        width: 100%;
        flex-grow: 1;
        display: flex;
        .list {
            margin-left: 60px;
            margin-top: 20px;
            width: 60%;
            &>div{
                display: flex;
                height: 25px;
                line-height: 25px;
                font-size: 14px;
                color: #cfc;
                &>div{
                    padding-left: 30px;
                    box-sizing: border-box;
                    white-space: nowrap;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    width: 25%;
                    button {
                        background-color: rgba(0,0,0,0);
                        font-size: 12px;
                        margin-top: 2px;
                        color:#fff;
                        border: none;
                        border-bottom: 1px solid #fcf;
                        
                    }
                }
            }
            .head{
                color: #fff;
            }
        }
    }
}
</style>