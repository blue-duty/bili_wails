<script setup lang=ts>
import { onMounted, ref } from "vue";
import { GetVideoInfo } from "../../wailsjs/go/main/App";
import type { app } from '../../wailsjs/go/models';


const video_info = ref<app.VideoInfo>();
const loading = ref<boolean>(true);
const videos = ref<app.Episode[]>([]);

onMounted(() => {
    // 设置页面加载中
    GetVideoInfo("https://www.bilibili.com/video/BV19w411p7u2?spm_id_from=333.1007.tianma.3-4-10.click").then((res: app.VideoInfo) => {
        video_info.value = res;
        videos.value = res.eps;
        loading.value = false;
    });
})

function selectRow(ep: app.Episode) {
    // 选中行
    ep.selected = !ep.selected;
}

// 全选
function selectAll() {
    videos.value.forEach((ep: app.Episode) => {
        ep.selected = true;
    });
}

</script>

<template>
    <main class="main-container">
        <header class="header">
            <input type="text" placeholder="Search" class="search-bar" />
            <button class="search-button">Search</button>
        </header>
        <!-- Video information section -->
        <div v-if="loading" class="loading-overlay">
            <div class="loading-spinner"></div>
        </div>
        <div v-else class="video-info">
            <img :src="video_info?.pic" alt="logo" />
            <article class=" article">
                <h1 class="article-title">{{ video_info?.title }}</h1>
                <p class="article-meta">
                    <span style="margin-bottom: 5px;">{{ video_info?.pubdate }}</span>
                    <span>
                        <span>{{ video_info?.stat.view }}播放 </span>
                        <span>{{ video_info?.stat.danmaku }}弹幕 </span>
                        <span>{{ video_info?.stat.reply }}评论 </span>
                        <span>{{ video_info?.stat.favorite }}收藏 </span>
                        <span>{{ video_info?.stat.coin }}投币 </span>
                        <span>{{ video_info?.stat.share }}分享 </span>
                        <span>{{ video_info?.stat.like }}点赞 </span>
                    </span>
                </p>
                <p class="article-dirc">{{ video_info?.desc }}</p>
            </article>
            <div class="face">
                <img :src="video_info?.owner_face" alt="logo" />
                <p>{{ video_info?.owner_name }}</p>
            </div>
        </div>
        <!-- Table for file details -->
        <div v-if="!loading" class="file-table">
            <table>
                <thead>
                    <th>序号</th>
                    <th style="width: 400px;">标题</th>
                    <th>时长</th>
                    <th>音质</th>
                    <th>画质</th>
                </thead>
            </table>
            <div class="tbody-container">
                <table>
                    <tbody>
                        <tr v-for="ep in videos " :key="ep.index" :class="{ 'selected': ep.selected }"
                            @click="selectRow(ep)">
                            <td>{{ ep.index }}</td>
                            <td style="width: 400px;">{{ ep.title }}</td>
                            <td>{{ ep.duration }}</td>
                            <td>{{ "null" }}</td>
                            <td>{{ "null" }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
        <div class="bott-bar">
            <!-- 勾选全选 -->
            <!-- 勾选框 -->
            <div>
                <input type="checkbox" id="scales" name="scales" checked />
                <label for="scales">Scales</label>
            </div>
            <div class="search-bar">
                <input type="text" placeholder="Search" />
                <button class="search-button">搜索</button>
            </div>
            <button class="download-button" style="right: 3%; position: fixed;">下载</button>
        </div>
    </main>
</template>

<style scoped>
.loading-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(255, 255, 255, 0.9);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.loading-spinner {
    border: 5px solid #f3f3f3;
    border-top: 5px solid #3498db;
    border-radius: 50%;
    width: 50px;
    height: 50px;
    animation: spin 2s linear infinite;
}

@keyframes spin {
    0% {
        transform: rotate(0deg);
    }

    100% {
        transform: rotate(360deg);
    }
}

.main-container {
    height: 95%;
    width: 95%;
    font-family: 'Arial', sans-serif;
    margin: auto;
    background: #FFF;
    padding: 10px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.video-info {
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;
    height: 30%;
}

.video-info img {
    object-fit: cover;
}

.article {
    margin: 0 10px;
    display: flex;
    flex-direction: column;
}

.article-title {
    font-size: 15px;
    font-weight: 550;
    margin: 0 0 10px 0;
}

.article-meta {
    font-size: 12px;
    color: #666;
    margin: 0 0 10px 0;
    display: flex;
    justify-content: space-between;
    flex-direction: column;
}

.article-dirc {
    text-wrap: wrap;
    font-size: 12px;
    max-width: 100%;
    /* 行间距 */
    line-height: 1.5;
}

.face {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.face img {
    width: 50px;
    height: 50px;
    border-radius: 50%;
}

.face p {
    margin-top: 10px;
    font-size: 10px;
    text-wrap: nowrap;
}

.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.search-bar {
    width: 300px;
    border-radius: 4px;
    position: fixed;
    display: flex;
    right: 30%;
}

.search-button {
    margin-left: 10px;
}

.selected {
    background-color: #0d79d7;
}

.file-table {
    /* width: 100%; */
    margin-top: 10px;
    border: 1px solid #ccc;
}

.file-table table {
    /* 高度 */
    width: 100%;
    border-collapse: collapse;
    table-layout: fixed;
    border: black;
}

.tbody-container::-webkit-scrollbar {
    display: none;
    /* 隐藏WebKit浏览器的滚动条 */
}


.tbody-container {
    max-height: 200px;
    /* 这是示例高度，您可以根据需要调整 */
    overflow-y: visible;
    /* 隐藏滚动条 */
    overflow-x: hidden;
}

thead {
    background-color: rgb(215, 220, 217);
}

th {
    text-align: left;
    border: 1px solid #ffffff;
    padding: 10px;
    font-size: 14px;
}

td {
    text-align: left;
    border: none;
    padding: 10px;
    /* height: 10px */
    font-size: 12px;
}

/* 鼠标悬停时，背景颜色变亮 */
tr:hover {
    background-color: #ddd;
    /* 鼠标样式 */
    cursor: pointer;
}


.search-button:hover,
.download-button:hover {
    background-color: #1C86EE;
}

.bott-bar {
    display: flex;
    margin-top: 10px;
    width: 100%;
}

.search-button,
.download-button {
    border: none;
    border-radius: 1px;
    height: 30px;
    width: 50px;
    cursor: pointer;
    background-color: #1E90FF;
    color: white;
}
</style>
