<script setup lang=ts>
</script>

<template>
    <main>
        <div style="width: 100%; display: flex;">
            <input type="text" placeholder="保存位置" readonly style="width: 100%;" />
            <button style="text-wrap: nowrap;">选择位置</button>
        </div>
        <div style="display: flex;">
            <img src="../assets/images/bili_wails.png" alt="logo" width="250" height="226" />
            <text for="">
                <p>视频链接343454666</p>
                <p> jjjjjjjj</p>
            </text>
            <img src="../assets/images/bili_wails.png" alt="logo" width="250" height="226" />
        </div>
        <div>
            <!-- 表格 -->
            <table style="border: 1;">
                <tr>
                    <th>文件名</th>
                    <th>大小</th>
                    <th>进度</th>
                    <th>状态</th>
                    <th>操作</th>
                </tr>
                <tr>
                    <td>100</td>
                    <td>100</td>
                    <td>100</td>
                    <td>状态</td>
                    <td><button>下载</button></td>
                </tr>
                <tr>
                    <td>100</td>
                    <td><select>
                            <option>选项1</option>
                            <option>选项2</option>
                        </select></td>
                    <!-- <td>I'm your second cell.</td> -->
                    <td>I'm your third cell.</td>
                    <td>I'm your fourth cell.</td>
                    <td>I'm your fifth cell.</td>
                </tr>
            </table>
        </div>
    </main>
</template>

<style scoped lang="css">
table {
    border-collapse: collapse;
}

td,
th {
    border: 1px solid black;
    padding: 10px 20px;
}
</style>
