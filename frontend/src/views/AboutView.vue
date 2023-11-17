<script setup lang="ts">
import { onMounted, reactive } from 'vue';
import { BiliLogin, GetCookie } from '../../wailsjs/go/main/App';

// import { useI18n } from "vue-i18n";

// import comeonGif from "@/assets/images/comeon.gif";

// const { t } = useI18n();

const data = reactive({
  urlsrc: "",
});

//  通过获取的base64数据创建图片
function createImg() {
  BiliLogin().then((result: string) => {
    data.urlsrc = result;
  });

  // 登录校验
  GetCookie()
}

// 页面加载完成后执行
onMounted(() => {
  createImg();
});


</script>

<template>
  <div class="about">
    <!-- 展示登录二维码 -->
    <!-- <img class="title" :src="data.urlsrc" alt="title" /> -->
    <div>
      <input type="checkbox" id="scales" name="scales" />
      <label for="scales">Scales</label>
    </div>
  </div>
</template>

<style lang="scss">
.about {
  .title {
    margin: 30px auto 10px;
    font-size: 38px;
    color: #a150b5;
    text-align: center;
  }

  .content {
    position: relative;
    margin: 36px 20px;

    .comeon {
      position: absolute;
      left: 26px;
      top: 38px;
      max-width: 66%;

      img {
        width: 220px;
        height: 180px;
      }
    }

    .info {
      margin: 0 0 0 33%;
      font-size: 24px;
      text-align: left;

      .info-item {
        margin-bottom: 10px;

        .name {
          line-height: 40px;
          font-size: 28px;
          color: #6d6363;
        }

        .link {
          line-height: 30px;
          font-size: 20px;
          color: #5f6c86;
        }
      }
    }
  }

  .thank {
    height: 68px;
    line-height: 68px;
    margin: 36px auto;
    text-align: center;
    font-size: 40px;
  }
}
</style>
