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
    // proxy: "https://testadmin.yinsfinance.com" //测试 http://121.199.167.227:5002/
    proxy: "http://121.199.167.227:5002/" //测试 http://121.199.167.227:5002/

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
