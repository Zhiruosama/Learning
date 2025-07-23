// 自己编写声明文件
declare module 'express' {
    interface Router {
        get(path: string, cb: (req: any, res: any) => void): void
    }
    interface App {
        use(path: string, router: any): void
        listen(port: number, cb?: () => void)
    }
    interface Express {
        (): App
        Router(): Router
    }
    const express: Express

    export default express
}

// 还可以用declare扩充全局变量
// 还有函数也是可以的 且引入以后是有代码提示的
declare var a:number