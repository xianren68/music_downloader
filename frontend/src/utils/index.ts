// 节流
export function throttle(func: Function, wait: number) {
    let timer:any
    return function (...args: any[]) {
        const context= this
        if (!timer) {
            timer = setTimeout(() => {
                func.apply(context, args)
                timer = undefined
            },wait)
        }
    }
}