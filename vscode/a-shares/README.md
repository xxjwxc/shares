<div align="center">
<img src="https://raw.githubusercontent.com/xxjwxc/shares/master/image/0.png" alt="复利备忘录" width="256"/>

# a-shares（复利备忘录）

**复利备忘录**——VSCode 里也可以看股票 & 量化 做最好用的投资插件。

Leek box - Monitor the real-time data of stock & fund & future in VSCode, Be the most excellent investment extension.

[![Marketplace](https://img.shields.io/visual-studio-marketplace/v/a-shares..a-shares.svg?label=Marketplace&style=for-the-badge&logo=visual-studio-code)](https://marketplace.visualstudio.com/items?itemName=a-shares.a-shares)
[![Installs](https://img.shields.io/visual-studio-marketplace/i/a-shares..a-shares.svg?style=for-the-badge)](https://marketplace.visualstudio.com/items?itemName=a-shares..a-shares)
[![Rating](https://img.shields.io/visual-studio-marketplace/stars/a-shares..a-shares.svg?style=for-the-badge)](https://marketplace.visualstudio.com/items?itemName=a-shares..a-shares)

投资有风险，入市需谨慎！

</div>

## Table of contents

- [a-shares（复利备忘录）](#a-shares复利备忘录)
  - [Table of contents](#table-of-contents)
  - [功能特性](#功能特性)
  - [安装使用](#安装使用)
  - [插件介绍](#插件介绍)
  - [插件设置](#插件设置)
  - [社区交流](#社区交流)
  - [感谢 PR](#感谢-pr)
  - [代码目录说明](#代码目录说明)
  - [License](#license)

> 投资其实就是一次心态修炼，稳住心态长期投资都会有收益的！！

## 功能特性

本插件具有以下特点：

- 基金实时涨跌，实时数据，支持海外基展示
- 股票实时涨跌，支持 A 股、港股、美股
- 开市自动刷新，节假日关闭轮询
- 支持升序/降序排序、基金持仓金额升序/降序
- 基金实时走势图和历史走势图
- 股市资金流向（沪深港通资金流向、北向资金、南向资金）
- 日频股东数
- 日频公募数据
- 金融大模型，复利Chat，让投资跟聊天一样简单
- 温度计
- 全链路量化，行业板块分析，直接贴图。欢迎体验
- 北向，机构数据一眼清
- 个股诊断
- 黄金上穿
- 欢迎 PR [Github 源码](https://github.com/xxjwxc/shares)

## 安装使用

安装插件：[VisualStudio - Marketplace](https://marketplace.visualstudio.com/items?itemName=a-shares..a-shares)，VSCode 最低版本要求：`^1.44.0`

## 插件介绍

- [复利备忘录使用文档](https://github.com/xxjwxc/shares)
- [VSCode 插件开发——复利备忘录（图片如果展示不了可以看知乎的文章界面功能截图）](https://blog.csdn.net/xie1xiao1jun/article/details/131491699)

<!-- https://raw.staticdn.net/ 为GitHub raw 加速地址 -->

![概览](https://github.com/xxjwxc/shares)


## 插件设置

**添加/删除股票时，建议使用新增按钮模糊搜索添加（支持名称和编码搜索）**，详细可查看 [复利备忘录使用文档](https://github.com/xxjwxc/shares)

自定义配置在 **Settings** 视图下：

## 社区交流

公众号：

<img width="200" alt="微信公众号" src="https://raw.githubusercontent.com/xxjwxc/shares/master/image/0.png">

## Core Contributors

- [a-shares](https://github.com/xxjwxc/shares)





## 代码目录说明

> 历史原因，仓库中类文件并没有以`PascalCase`规范，导致有些文件不好区分是函数方式书写还是面向对象类的写法。

```shell

src
├── data                        # 静态数据
│   └── fundSuggestData.ts      # 基金数据，执行 `node ./demo/fundSuggestList.js` 更新生成
├── explorer                    # 侧边栏核心代码
│   ├── binanceProvider.ts      # 数字货币
│   ├── binanceService.ts
│   ├── forexProvider.ts        # 外汇
│   ├── forexService.ts
│   ├── fundProvider.ts         # 基金
│   ├── fundService.ts
│   ├── leekService.ts
│   ├── newsProvider.ts         # 雪球新闻
│   ├── newsService.ts
│   ├── stockProvider.ts        # 股票
│   └── stockService.ts
├── extension.ts                # 插件初始化入口
├── globalState.ts              # 全局缓存，插件激活到销毁周期内的变量缓存
├── output                      # Terminal 视图下的OUTPUT栏输出新闻
│   └── flash-news
├── registerCommand.ts          # 注册命令
├── shared                      # 工具函数或者类
│   ├── WVMessageUtils.ts
│   ├── constant.ts
│   ├── holidayHelper.ts
│   ├── leekConfig.ts
│   ├── leekTreeItem.ts
│   ├── remindNotification.ts
│   ├── telemetry.ts
│   ├── typed.ts
│   └── utils.ts
├── statusbar                   # 状态栏
│   ├── Profit.ts
│   └── statusBar.ts
└── webview   # webview 页面

```

## 赞助支持一下 ↓↓

[paypal](https://www.paypal.me/xxjwxc)




