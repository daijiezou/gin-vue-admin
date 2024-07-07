<template>
  <h3>工地详情</h3>
</template>

<script setup>
import {useRoute} from 'vue-router'
const route = useRoute()
// 打印query参数
console.log(route.params)
</script>



<style scoped lang="scss">

</style>