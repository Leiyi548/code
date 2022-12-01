## 安装 TypeScript

有两种主要的方式来获取 TypeScript 工具：

- 通过 npm（Node.js 包管理器）
- 安装 Visual Studio 的 TypeScript 插件

Visual Studio 2017 和 Visual Studio 2015 Update 3 默认包含了 TypeScript。如果你的 Visual Studio 还没有安装 TypeScript，你可以下载它。

针对使用 npm 的用户：

> npm install -g typescript

## 安装 ts-node

安装它的原因是 typescript 自带的 tsc 命令并不能直接运行 typescript 代码。但值得注意的是 ts-node 并不等于 typescript 的 Node.js，仅仅封装了 typescript 的编译过程，提供直接运行 typescript 代码的能力。

> npm install -g ts-node
