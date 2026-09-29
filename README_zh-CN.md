[English Version](README.md)

<h1 align="center">
  <br>
  <a href="https://cloudreve.org/" alt="logo" ><img src="https://raw.githubusercontent.com/cloudreve/frontend/master/public/static/img/logo192.png" width="150"/></a>
  <br>
  Cloudreve
  <br>
</h1>

<h4 align="center">基于 Cloudreve 4.18.0 的离线下载免拷贝移动版</h4>

> **这是基于 [Cloudreve 4.18.0](https://github.com/cloudreve/cloudreve/releases/tag/4.18.0) 的非官方修改版。** 主节点的离线下载完成后，若目标为本地存储且源文件与目标目录位于同一文件系统，会先建立硬链接，待上传完成后再删除源文件，避免再次拷贝文件内容。跨文件系统时仍回退到拷贝；做种任务及其他存储路径保持原有行为。

<p align="center">
  <a href="https://github.com/Corezcy/cloudreve/actions/workflows/fork-release.yml">
    <img src="https://github.com/Corezcy/cloudreve/actions/workflows/fork-release.yml/badge.svg?branch=feat%2F4.18.0-zero-copy-remote-download"
         alt="修改版构建">
  </a>
  <a href="https://github.com/Corezcy/cloudreve/releases">
    <img src="https://img.shields.io/github/v/release/Corezcy/cloudreve?include_prereleases" alt="修改版发布" />
  </a>
</p>
<p align="center">
  <a href="https://cloudreve.org">主页</a> •
  <a href="https://demo.cloudreve.org">演示</a> •
  <a href="https://github.com/cloudreve/cloudreve/discussions">讨论</a> •
  <a href="https://docs.cloudreve.org">文档</a> •
  <a href="https://github.com/Corezcy/cloudreve/releases">下载修改版</a> •
  <a href="https://github.com/cloudreve/cloudreve">上游项目</a> •
  <a href="https://t.me/cloudreve_official">Telegram</a> •
  <a href="https://discord.com/invite/WTpMFpZT76">Discord</a>
</p>

![Screenshot](https://raw.githubusercontent.com/cloudreve/docs/master/images/homepage.png)

## :sparkles: 特性

- :cloud: 支持本机、从机、七牛 Kodo、阿里云 OSS、腾讯云 COS、华为云 OBS、金山云 KS3、又拍云、OneDrive (包括世纪互联版) 、S3 兼容协议 作为存储端
- :outbox_tray: 上传/下载 支持客户端直传，支持下载限速
- 💾 可对接 Aria2/qBittorrent 离线下载，可使用多个从机节点分担下载任务
- 📚 在线 压缩/解压缩/压缩包预览、多文件打包下载
- 💻 覆盖全部存储策略的 WebDAV 协议支持
- :zap: 拖拽上传、目录上传、并行分片上传
- :card_file_box: 提取媒体元数据，通过元数据或标签搜索文件
- :family_woman_girl_boy: 多用户、用户组、多存储策略
- :link: 创建文件、目录的分享链接，可设定自动过期
- :eye_speech_bubble: 视频、图像、音频、 ePub 在线预览，文本、Office 文档在线编辑
- :art: 自定义配色、黑暗模式、PWA 应用、全站单页应用、国际化支持
- :rocket: All-in-One 打包，开箱即用
- 🌈 ... ...

## :hammer_and_wrench: 部署

你可以参考 [快速开始](https://docs.cloudreve.org/overview/quickstart) 启动一个本地实例进行体验、测试。

当你准备好将 Cloudreve 部署到生产环境时，可以参考 [部署](https://docs.cloudreve.org/overview/deploy/) 进行完整部署。

## :gear: 构建

你可以参考 [构建](https://docs.cloudreve.org/overview/build/) 从源代码构建 Cloudreve。

此修改版也通过 [GitHub Actions](https://github.com/Corezcy/cloudreve/actions/workflows/fork-release.yml) 构建、测试，并在 [Releases](https://github.com/Corezcy/cloudreve/releases) 提供二进制下载；此 fork 不发布 Docker 镜像。

## :rocket: 贡献

如果你有兴趣为 Cloudreve 贡献代码，请参考 [贡献](https://docs.cloudreve.org/api/contributing/) 了解如何贡献。

## :alembic: 技术栈

- [Go](https://golang.org/) + [Gin](https://github.com/gin-gonic/gin) + [ent](https://github.com/ent/ent)
- [React](https://github.com/facebook/react) + [Redux](https://github.com/reduxjs/redux) + [Material-UI](https://github.com/mui-org/material-ui)

## :scroll: 许可证

GPL V3
