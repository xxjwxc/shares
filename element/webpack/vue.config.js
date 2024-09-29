const { defineConfig } = require('@vue/cli-service')
const timeStamp = new Date().getTime()
module.exports = defineConfig({
  transpileDependencies: true,


  configureWebpack: { // 输出重构 打包编译后的js文件名称,添加时间戳.
    output: {
      filename: `js/[name].js?v+${timeStamp}` ,
     chunkFilename: `js/chunk.[id].js?v+${timeStamp}`
     }
  },
  css: {
    extract: { // 打包后css文件名称添加时间戳
      filename: `css/[name].css?v+${timeStamp}`,
      chunkFilename: `css/chunk.[id].css?v+${timeStamp}`
    }
  },
  publicPath: '/pc',
  lintOnSave: false,
  devServer: {
    host: 'localhost',
    port: 8080,  // 启动端口号
    proxy: {
      '/shares': { // 请求接口中要替换的标识
        // target: 'http://localhost:82', // 代理地址
        // target: 'http://192.155.1.28:82', // 代理地址
        target: 'https://www..cn', // 代理地址
        ChangeOrigin: true, // 是否允许跨域
        secure: true,
        "logLevel": "debug"
      },
      '/open_api': { // 请求接口中要替换的标识
        // target: 'http://localhost:82', // 代理地址
        // target: 'http://192.155.1.28:82', // 代理地址
        target: 'https://api.coze.cn', // 代理地址
        ChangeOrigin: true, // 是否允许跨域
        secure: true,
        "logLevel": "debug"
      },
    }
  }
})
