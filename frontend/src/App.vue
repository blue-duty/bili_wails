<script setup lang="ts">
// import { useI18n } from "vue-i18n";
import { reactive } from "vue";
import { GetAvatar, Minimize, Quit } from "../wailsjs/go/main/App";

// const { t, availableLocales: languages, locale } = useI18n();

const data = reactive({
  img: "",
});

const onclickMinimise = () => {
  Minimize();
};

const onclickQuit = () => {
  Quit();
};

function getImg() {
  // 获取用户头像
  GetAvatar().then((result: string) => {
    data.img = result;
  });
}

document.body.addEventListener("click", function (event) {
  event.preventDefault();
});
</script>

<template>
  <!-- Header -->
  <div class="header" style="--wails-draggable:drag">
    <!-- navigation -->
    <div class="nav">
      <router-link to="/">首页</router-link>
      <router-link to="/download">下载</router-link>
      <router-link to="/about">关于</router-link>
    </div>
    <!-- Menu -->
    <div class="menu">
      <!-- <div class="language">
        <div v-for="item in languages" :key="item" :class="{ active: item === locale }"
          @click="onclickLanguageHandle(item)" class="lang-item">
          {{ t("languages." + item) }}
        </div>
      </div> -->
      <div class="bar">
        <!-- 用户头像：圆形 -->
        <div style="width: 50px;height: 50px; border-radius: 50%; overflow: hidden;background-color: #98c9a5">
          <!-- 解决跨域问题 -->
          <img :src="data.img" alt="user" width="50" height="50" @click="getImg" />
        </div>
        <div class="bar-btn" @click="onclickMinimise">
          最小化
        </div>
        <div class="bar-btn" @click="onclickQuit">退出</div>
      </div>
    </div>
  </div>
  <!-- Page -->
  <div class="view">
    <router-view />
  </div>
</template>

<style lang="scss">
@import url("./assets/css/reset.css");
@import url("./assets/css/font.css");

html {
  width: 100%;
  height: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "JetBrainsMono";
  background-color: transparent;
}

#app {
  position: relative;
  // width: 900px;
  // height: 520px;
  height: 100%;
  width: 100%;
  background-color: #87CEFA;
  overflow: hidden;
}

.header {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: space-between;
  height: 50px;
  padding: 0 10px;
  background-color: #87ebfaa5;

  .nav {
    a {
      display: inline-block;
      min-width: 50px;
      height: 30px;
      line-height: 30px;
      padding: 0 5px;
      margin-right: 8px;
      background-color: #98c9a5;
      ;
      border-radius: 2px;
      text-align: center;
      text-decoration: none;
      color: #000000;
      font-size: 14px;
      white-space: nowrap;

      &:hover,
      &.router-link-exact-active {
        background-color: #1563c3;
        color: #ffffff;
      }
    }
  }

  .menu {
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    align-items: center;
    justify-content: space-between;

    .language {
      margin-right: 20px;
      border-radius: 2px;
      background-color: #c3c3c313;
      overflow: hidden;

      .lang-item {
        display: inline-block;
        min-width: 50px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        background-color: transparent;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;

        &:hover {
          background-color: #ff050542;
          cursor: pointer;
        }

        &.active {
          background-color: #ff050542;
          color: #ffffff;
          cursor: not-allowed;
        }
      }
    }

    .bar {
      display: flex;
      flex-direction: row;
      flex-wrap: nowrap;
      align-items: center;
      justify-content: flex-end;
      min-width: 150px;

      .bar-btn {
        display: inline-block;
        min-width: 80px;
        height: 30px;
        line-height: 30px;
        padding: 0 5px;
        margin-left: 8px;
        background-color: #98c9a5;
        border-radius: 2px;
        text-align: center;
        text-decoration: none;
        color: #000000;
        font-size: 14px;

        &:hover {
          background-color: #d7a8d8;
          color: #ffffff;
          cursor: pointer;
        }
      }
    }
  }
}

.view {
  position: absolute;
  top: 50px;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
}
</style>
