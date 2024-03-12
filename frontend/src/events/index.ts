import { ConfigStore } from "@/store"
import { ElMessage } from "element-plus"
import { EventsEmit, EventsOn } from "@/../wailsjs/runtime"
import { DownStore } from "@/store"
import { DownloadInt, ProgressInt, Music } from "@/type"
import { SaveConfig} from "@/../wailsjs/go/main/App"
import { toRaw } from "vue"
// 消息弹窗事件
EventsOn("info", (message: string) => {
    ElMessage.info(message)
})
// 警告弹窗事件
EventsOn("warning", (message: string) => {
    ElMessage.warning(message)
})
// 成功弹窗事件
EventsOn("success", (message: string) => {
    ElMessage.success(message)
})
// 错误弹窗事件
EventsOn("error", (message: string) => {
    ElMessage.error(message)
})

// 继续下载队列
EventsOn("continueDownload", (list: Array<Music>) => {
    for (let i of list) {
        const downItem: DownloadInt = { ...i, isStart: false, progress: 0 }
        DownStore().downList.set(i.id, downItem)
        DownStore().size = DownStore().downList.size
    }
})

// 下载进度事件
EventsOn("downloadProgress", (progressInfo: ProgressInt) => {
    const downStore = DownStore()
    const value = downStore.downList.get(progressInfo.id)
    console.log(value?.name)
    if (progressInfo.status) {
        // 下载完成
        ElMessage.success(`${value?.name}下载完成`)
        downStore.downList.delete(progressInfo.id)
        downStore.size--
        return 
    }
    // 修改下载进度
    value!.progress = progressInfo.progress
    value!.isStart = true
    downStore.downList.set(progressInfo.id, value!)
})

// 收尾工作
EventsOn("ending", () => {
    const configStore = ConfigStore()
    // 更新后端配置
    SaveConfig(JSON.stringify(toRaw(configStore.config)))
    // 触发后端收尾
    EventsEmit("end",null)
    
})

// 需要让这个文件中的代码执行
export default null