<template>
    <div class="download">
        <div class="head">
            <div class="title">下载队列</div>
            <div class="quantity"><span>{{downStore.downList.size}}</span>个下载任务</div>
        </div>
        <div class="content">

        </div>
    </div>
</template>

<script setup lang="ts">
import { DownStore } from '@/store'
const downStore = DownStore()
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
}
</style>