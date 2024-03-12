<template>
    <div class="setting">
        <div class="save">
            <p class="title">下载设置</p>
            <div class="savePath">
                <p class="label">保存路径：</p>
                <input type="text" class="input" v-model="configStore.config.Save.Path">
                <button class="select" @click="selectPath">选择目录</button>
            </div>
            <div class="sort">
                <p class="label">下载分类:</p>
                <input type="radio"  :checked="configStore.config.Save.IsSortByAuthor" name="sort" @change="configStore.config.Save.IsSortByAuthor=true">
                <label class="label" for="sort">按作者分类</label>
                <input type="radio" :checked="!configStore.config.Save.IsSortByAuthor" name="notsort" @change="configStore.config.Save.IsSortByAuthor=false"> 
                <label class="label" for="notsort">不分类</label>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import {OpenDirDialog,SaveConfig } from '@/../wailsjs/go/main/App'
import {onUnmounted,toRaw} from 'vue'
import {ElMessage} from 'element-plus'
import { ConfigStore } from '@/store'
const configStore = ConfigStore()
// 选择目录
const selectPath = async () => {
    // webview自带的api无法获取文件夹绝对路径，需要使用wails的api
    const path = await OpenDirDialog()
    configStore.config.Save.Path = path
}
// 向后台同步配置
onUnmounted(async ()=>{
    if (!await SaveConfig(JSON.stringify(toRaw(configStore.config)))){
        ElMessage.error('保存配置失败')
    }
})
</script>

<style lang="scss" scoped>
.setting {
    height: 100%;
    padding: 0 20px;
    box-sizing: border-box;
    border-top: 1px solid rgba(0, 0, 0, 0);
    flex-grow: 1;
    .label {
        margin-right: 10px;
        font-size: 14px;
        color: #999;
    }
    .savePath {
        display: flex;
        height: 40px;
        align-items: center;
        margin-top: 10px;
        width: 600px;
        .input {
            width: 60%;
            border-bottom: #000 1px solid;
        }
    }
    .sort{
        display: flex;
        height: 40px;
        align-items: center;
        margin-top: 8px;
        input {
            margin-left: 10px;
        }
        label{
            margin-left: 5px;
        }
    }
}

.title {
    margin-top: 20px;
    font-size: 16px;
    font-weight: 600;
    color: #999;
}

.save {
    margin-bottom: 20px;
}
</style>