/*
 * @Author: qiubenyang qiubenyang@mycaiwen.com
 * @Date: 2022-06-05 13:54:25
 * @LastEditors: qiubenyang qiubenyang@mycaiwen.com
 * @LastEditTime: 2022-06-12 18:15:37
 * @FilePath: /OnlineTeaching/vue.config.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
const { defineConfig } = require("@vue/cli-service");
const path = require("path");

const IS_PRODUCTION = process.env.NODE_ENV === "production";

function resolve(dir) {
  return path.join(__dirname, dir);
}

module.exports = defineConfig({
  publicPath: "./",
  // see https://cli.vuejs.org/zh/config/#transpiledependencies
  transpileDependencies: false,
  devServer: {
    static: {
      directory: path.join(__dirname, "static")
    },
    // host: "127.0.0.1",
    port: 8090,
    // proxy: "http://121.199.167.227:5002/"
    proxy: "http://192.168.3.17:5002/"

    // devtool: "cheap-source-map"
  },
  css: {
    sourceMap: !IS_PRODUCTION,
    loaderOptions: {
      less: {
        lessOptions: {
          paths: [path.resolve(__dirname, "node_modules")],
          modifyVars: {
            hack: `true; @import "${resolve("./src/styles/theme")}";`
          }
        }
      }
    }
  },
  configureWebpack: config => {
    // see https://cli.vuejs.org/zh/guide/webpack.html
    if (process.env.NODE_ENV === "production") {
      // 为生产环境修改配置...
    } else {
      // 为开发环境修改配置...
    }
  },
  chainWebpack: config => {}
});
