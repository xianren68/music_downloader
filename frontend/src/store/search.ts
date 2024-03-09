import {defineStore} from 'pinia'
import {ref} from 'vue'
import {Abslist} from '@/type'
export const SearchStore = defineStore('search', ()=>{
    const searchText = ref('')
    const ResultList = ref<Abslist[]>([])
    return {searchText,ResultList}
}   
)