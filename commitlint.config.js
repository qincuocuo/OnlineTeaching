/*
 * feat：新功能
 * fix：bug 修复
 * refactor：重构代码(既没有新增功能，也没有修复 bug)
 * docs：文档更新
 * chore：其他修改, 比如改变构建流程、或者增加依赖库、工具等
 * style：不影响程序逻辑的代码修改(修改空白字符，格式缩进)
 * build：主要目的是修改项目构建系统(例如 glup，webpack，rollup 的配置等)的提交
 * perf：性能, 体验优化
 */
module.exports = {
  extends: ["@commitlint/config-conventional"],
  rules: {
    "type-enum": [
      2,
      "always",
      ["feat", "fix", "refactor", "docs", "chore", "style", "build", "perf"]
    ]
  }
};
