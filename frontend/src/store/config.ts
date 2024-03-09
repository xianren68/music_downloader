import {defineStore} from 'pinia'
import {ref} from 'vue'
import type { allConfig } from '@/type'
export const ConfigStore = defineStore('config', ()=>{
    const config = ref<allConfig>({Save:{Path:'',IsSortByAuthor:false}})
    return {config}
})