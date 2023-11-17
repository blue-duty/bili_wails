<script setup lang="ts">

import { reactive } from 'vue'
import { useRouter } from 'vue-router'
import { OpenDir, GetClipboard, DownloadVedio, OpenFileExplorer } from '../../wailsjs/go/main/App'
import { useRouteStore } from '@/stores/counter'

const data = reactive({
  dir: "",
  clipboard: "",
  status: "未开始...",
  proc: 0,
  inputing: false
})

const bili_url = useRouteStore()
const router = useRouter()

function getDir() {
  OpenDir().then((result: string) => {
    data.dir = result
  })
}

function getClipboard() {
  GetClipboard().then((result: string) => {
    data.clipboard = result
  })
}

function download() {
  if (data.dir === "") {
    data.status = "请选择保存位置"
    return
  }

  if (data.clipboard === "") {
    data.status = "请粘贴视频链接"
    return
  }
  data.status = "下载中..."

  DownloadVedio(data.clipboard, data.dir).then((result: string) => {
    // console.log(result)
    data.status = result
    data.proc = 100
  })

  OpenFileExplorer(data.dir + "/download").then((result: any) => {
    // console.log(result)
  })
}

// document.addEventListener('contextmenu', function (event) {
//   event.preventDefault();
// });

// 添加对于Ctrl+V的监听
document.addEventListener('keydown', function (event) {
  if (event.ctrlKey && event.key === 'v') {
    // getClipboard()
    event.preventDefault();
    if (data.clipboard) {
      if (data.inputing) {
        GetClipboard().then((result: string) => {
          data.clipboard = data.clipboard + result
        })
      } else {
        GetClipboard().then((result: string) => {
          data.clipboard = result
        })
      }
    } else {
      GetClipboard().then((result: string) => {
        data.clipboard = result
      })
    }
  }

  // delete
  if (event.key === 'Delete') {
    event.preventDefault();
    data.clipboard = ""
  }
});

//  设置bili_url
function setBiliUrl() {
  bili_url.setPath(data.clipboard)
  router.push("/download")
}


</script>

<template>
  <main>
    <div class="container">
      <div class="image-container">
        <img src="../assets/images/bili_wails.png" alt="logo" width="250" height="226" />
      </div>
      <div class="form">
        <div class="item">
          <label for="">视频链接</label>
          <!--  监听鼠标聚焦事件 -->
          <input type="text" v-model="data.clipboard" @focus="data.inputing = true" @blur="data.inputing = false"
            @keyup.enter="setBiliUrl" />
          <button @click="getClipboard" :class="{ 'hiddenbutton': data.clipboard }">粘贴</button>
        </div>
        <div class="item">
          <label for="">保存位置</label>
          <!-- 无边框 -->
          <input type="text" v-model="data.dir" readonly style="pointer-events: none; border: none;">
          <button @click="getDir">选择位置</button>
        </div>
        <div>
          <label class="status">
            {{ data.status }}
          </label>
          <div class="g-container">
            <div class="g-progress" :style="{ 'width': data.proc + '%' }"></div>
          </div>
        </div>
      </div>
      <button class="download" @click="download">下载视频</button>
    </div>
  </main>
</template>

<style lang="scss">
main {
  height: 100%;
  width: 100%;
}

.container {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.image-container {
  // margin: auto;
  margin: 5%;
  /* Adjust the margin as needed */
}

.g-container {
  width: 100%;
  height: 25px;
  border-radius: 5px;
  background: #ccc;
}

.g-progress {
  height: inherit;
  border-radius: 5px;
  background: #0f0;
}

.hiddenbutton {
  display: none;
}

.status {
  font-size: 14px;
  font-family: 'Courier New', Courier, monospace;
  text-wrap: nowrap;
  text-align: center;
}

.form {
  width: 60%;
  display: flex;
  flex-direction: column;
}

.item {
  width: 100%;
  display: flex;
  /* padding-bottom: 32px; */
  height: 40px;
  line-height: 32px;
  justify-content: center;
  position: relative;
}

.item label {
  font-size: 14px;
  font-family: 'Courier New', Courier, monospace;
  text-wrap: nowrap;
  // line-height: 40px;
  margin: 0;
}

.item input {
  width: 100%;
  height: 1.5rem;
  border-radius: 5px;
  text-indent: 2px;
  background-color: #ccc;
  /* pointer-events: none; */
  margin-left: 1rem;
  /* width: calc(100% - 30px); */
}



.item button {
  position: absolute;
  right: 3px;
  /* Adjust the right position as needed */
  top: 50%;
  transform: translateY(-70%);
  border: none;
  line-height: 1.5rem;
  height: 1.5rem;
  background-color: rgba(64, 190, 91, 0.485);
  border-radius: 5px 5px;
  text-wrap: nowrap;
  color: white;
  padding: 0 10px;
  cursor: pointer;
  text-align: center;
  font-size: 14px;
}

.download {
  position: fixed;
  bottom: 10px;
  right: 10px;
  padding: 10px;
  background-color: rgb(55, 151, 246);
  border-radius: 10px;
  color: white;
  cursor: pointer;
  font-size: 12px;
  font-family: 'Courier New', Courier, monospace;
}

.download:hover {
  background-color: #1563c3;
}
</style>
