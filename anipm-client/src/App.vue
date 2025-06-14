<template>
  <el-menu mode="horizontal" router :default-active="$route.path">
    <el-menu-item index="/">首页</el-menu-item>
    <el-menu-item index="/hello">Hello</el-menu-item>
  </el-menu>
  <router-view />
</template>
