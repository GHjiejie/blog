<template>
  <div class="header-container">
    <div class="left">
      <svg-icon iconClass="icon-blog-logo" className="icon"></svg-icon>
    </div>
    <Navbar></Navbar>
    <div class="right">
      Jie
    </div>
  </div>
</template>

<script setup>
import Navbar from '@/components/Header/components/Navbar/index.vue'
</script>
<style scoped lang="scss">
.header {
  height: 50px;
  line-height: 50px;
  text-align: center;
  background-color: #409eff;
  color: #fff;
}
</style>
