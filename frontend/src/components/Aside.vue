<template>
    <div class="aside">
        <div class="item" v-for="i in routes" :key="i.name" :class="{active:i.name == route.name}"
        @click="jump(i.name as string)">
            <svg class="icon">
                <use v-bind:xlink:href="'#'+i.meta.iconName"></use>
            </svg>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRoute,useRouter} from 'vue-router'
// 路由列表
const router = useRouter()
const routes = router.getRoutes()
const route = useRoute()
const jump = (name:string)=>{
    router.push({name})
}
</script>

<style scoped lang="scss">
.aside {
    height: calc(100% - 20px);
    width:60px;
    padding-top:20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    .item {
        width: 30px;
        height: 30px;
        margin-top:24px;
        cursor: pointer;
        .icon{
            height: 100%;
            width: 100%;
            transition: all .3s;
        }
        &:nth-child(3){
            position: absolute;
            bottom: 20px;
        }
        &.active{
            .icon {
                transform: scale(1.3);
            }
        }
    }
}
</style>