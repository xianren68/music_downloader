<template>
    <div class="search">
        <div class="searchbox">
                <input type="text" placeholder="请输入您要搜索的内容..." v-model="searchText">
                <button @click="search">搜索</button>
        </div>
        <div class="result">
            <div class="list">
                <div class="head" v-if="ResultList!=null">
                    <div class="name">名称</div>
                    <div class="author">作者</div>
                    <div class="albums">专辑</div>
                    <div class="download"></div>
                </div>
                <div class="item" v-for="(i,index) in ResultList" :key="index">
                    <div class="name">{{i.name}}</div>
                    <div class="author">{{showSignersOrAlbums(i.singers)}}</div>
                    <div class="albums">{{showSignersOrAlbums(i.albums)}}</div>
                    <div class="download"><button @click="dialogVisible=true;selected=index">下载</button></div>
                </div>
            </div>
            <Pagenation :total="total" :current="pageNo" @changePage="changePage" v-if="ResultList!=null"></Pagenation>
        </div>
        <el-dialog
        v-model="dialogVisible"
        width="500"
      >
        <span>音质选择</span>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="dialogVisible = false" type="success">无损</el-button>
            <el-button type="primary" @click="dialogVisible = false">
              高品
            </el-button>
          </div>
        </template>
      </el-dialog>
    </div>
</template>

<script setup lang="ts">
import {Result, SearchResult,Singers,Albums } from '@/type'
import { Search } from '@/../wailsjs/go/main/App'
import {ref} from 'vue'
import { ElMessage } from 'element-plus'
import Pagenation from '@/components/Pagenation.vue'
const searchText = ref('')
const ResultList = ref<Array<Result>|null>(null)
const total = ref(0)
const pageNo = ref(1)
// 选中要下载歌曲的索引
const selected = ref(0)
// 发送请求
const search = async () => {
    if(searchText.value === ""){
        return
    }
    // 前端发送有跨域问题，到后端去请求
    const res = await Search(pageNo.value,searchText.value)
    if(res === ""){
        ElMessage.error("请求出错")
        return
    }
    const result:SearchResult = JSON.parse(res)
    if(result.code !== "000000"){
        ElMessage.error(result.info)
        return
    }
    if(result.songResultData.result.length === 0){
        ElMessage.info("没有找到相关歌曲")
    }
    total.value = +result.songResultData.totalCount
    ResultList.value = result.songResultData.result
}
// 展示每个歌手
const showSignersOrAlbums = (value:Array<Singers|Albums>):string => {
    if(value.length === 0){
        return value[0].name
    }
    return value.reduce((pre,cur) => pre+cur.name,"")
}
// 切换页码
const changePage = (value:number) => {
    if(value < 1 || value > Math.ceil(total.value / 20)){
        return
    }
    pageNo.value = value
    search()
}

// 音质选择弹框
const dialogVisible = ref(false)

</script>

<style scoped lang="scss">
.search {
    height: 100%;
    width:calc(100% - 60px);
    padding-top:1px;
    display: flex;
    flex-direction: column;
    .searchbox{
        margin-top: 50px;
        margin-left:90px;
        display: flex;
        align-items: center;
        input{
            box-sizing: border-box;
            width: 400px;
            height: 30px;
            border: 1px solid #ccc;
            border-right: none;
            border-radius: 5px 0 0 5px;
            padding-left: 10px;
        }
        button{
            height: 30px;
            box-sizing: border-box;
            width:80px;
            text-align: center;
            border:1px solid #ccc;
            background-color:rgba(234,233,235,.6);
            border-radius:0 5px 5px 0;
            filter: blur(.3px);
        }

    }
    .result{
        width: 100%;
        flex-grow: 1;
        display: flex;
        align-items: center;
        .list {
            margin-left: 60px;
            margin-top: 20px;
            width: 60%;
            &>div{
                display: flex;
                height: 25px;
                line-height: 25px;
                font-size: 14px;
                color: #cfc;
                &>div{
                    padding-left: 30px;
                    box-sizing: border-box;
                    white-space: nowrap;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    width: 25%;
                    button {
                        background-color: rgba(0,0,0,0);
                        font-size: 12px;
                        margin-top: 2px;
                        color:#fff;
                        border: none;
                        border-bottom: 1px solid #fcf;
                        
                    }
                }
            }
            .head{
                color: #fff;
            }
        }
    }
}
</style>