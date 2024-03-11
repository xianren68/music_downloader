<template>
    <div class="download">
        <div class="head">
            <div class="title">下载队列</div>
            <div class="quantity"><span>{{downStore.size}}</span>个下载任务</div>
        </div>
        <div class="content">
            <div class="item" v-for="i in downStore.downList.values()" :key="i.id">
                <div class="info">
                    <div>音乐名：{{i.name}}</div>
                    <div>专辑：{{i.album}}</div>
                    <div>作者：{{i.author}}</div>
                </div>
                <div class="progress" v-if="i.isStart">
                    <el-progress :percentage="i.progress" color="#008c8c" />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { DownStore } from '@/store'
import {watch,ref} from 'vue'
const downStore = DownStore()
// const values = ref(downStore.downList.values())
// watch (()=>downStore.downList,()=>{
//     values.value = downStore.downList.values()
// })
</script>

<style lang="scss" scoped>
.download {
    flex-grow: 1;
    padding: 40px 20px;
    .head {
        display: flex;
        justify-content: space-between;
        align-items: center;
        .title {
            font-size: 24px;
            font-weight: bold;
        }
        .quantity {
            font-size: 14px;
            color: #999;
            span {
                color: #333;
            }
        }
    }
    .content{
        margin-top: 20px;
        height: 85%;
        overflow: scroll;
        overflow-x: hidden;
        &::-webkit-scrollbar {
            display: none;
        }
        .item{
            height: 40px;
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            align-items: center;
            padding: 10px 0;
            border-bottom: 1px solid #eee;
            .info{
                display: flex;
                justify-content: space-around;
            }
            .progress{
                width: 100%;
            }
        }
    }
}
</style>