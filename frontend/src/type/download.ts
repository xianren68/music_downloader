export interface DownloadInt {
    id:string,
    name:string,
    album:string,
    isStart:boolean,
    author:string,
    progress:number,
}
export interface ProgressInt{
    id:string,
    progress:number,
    status:boolean,
}
export interface Music {
    id:string,
    album:string,
    name:string,
    author:string
}