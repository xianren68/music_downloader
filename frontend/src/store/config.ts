import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { allConfig } from '@/type'
import { GetConfig } from '@/../wailsjs/go/main/App'
export const ConfigStore = defineStore('config', () => {
    const config = ref<allConfig>({ Save: { Path: '', IsSortByAuthor: false } })
    // 配置信息
    GetConfig().then(res => {
        config.value = JSON.parse(res)
    })
    return { config }
})