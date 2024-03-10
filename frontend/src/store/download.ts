import {defineStore} from 'pinia'
import {ref} from 'vue'
import {DownloadInt} from '@/type'
export const DownStore = defineStore('down', ()=>{
    const size = ref(0)
    const downList = ref<Map<string,DownloadInt>>(new Map())
    return {downList,size}
})