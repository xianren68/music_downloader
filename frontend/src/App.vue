<template>
  <div class="main">
    <Aside></Aside>
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
import Aside from '@/components/Aside.vue'

</script>

<style lang="scss" scoped>
.main {
  height: 100%;
  display: flex;
  overflow: hidden;
  position: relative;
  background: linear-gradient(to right, #2c3e50,#bdc3c7);
  &::before{
    content: '';
        position: absolute;
        top: -20px; /* 使玻璃效果更突出，可根据需要调整 */
        left: -20px;
        right: -20px;
        bottom: -20px;
        background: inherit;
        filter: blur(10px); /* 使用模糊滤镜创建玻璃效果 */
        z-index: -20; /* 将伪元素置于底层，以使内容显示在上层 */
  }

}
</style>
