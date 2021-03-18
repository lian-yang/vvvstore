module.exports = {
  devServer: {
    // 配置跨域
    /* proxy: {
      "/api": {
        target: "http://127.0.0.1:9595",
        changeOrign: true,
        ws: false,
        pathRewrite: {
          "^/api": ''
        }
      }
    }, */
    // 配置vue自带的遮罩提示层
    overlay: {
      warnings: false,
      errors: false,
    }
  },
  // 关闭eslint检查
  lintOnSave: false,
}