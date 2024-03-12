import {defineStore} from 'pinia'
import {ref} from 'vue'
import {Abslist} from '@/type'
export const SearchStore = defineStore('search', ()=>{
    const searchText = ref('')
    const total = ref(1)
    const page = ref(1)
    const ResultList = ref<Abslist[]>([])
    return {searchText,ResultList,total,page}
}   
)