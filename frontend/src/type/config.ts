type SaveConfig = {
    Path:string, //下载路径
    isSortByAuthor:boolean, //是否按作者分类
}
// 所有配置
export type allConfig = {
    Save:saveConfig, //下载相关的配置
}