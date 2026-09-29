[中文版本](README_zh-CN.md)

<h1 align="center">
  <br>
  <a href="https://cloudreve.org/" alt="logo" ><img src="https://raw.githubusercontent.com/cloudreve/frontend/master/public/static/img/logo192.png" width="150"/></a>
  <br>
  Cloudreve
  <br>
</h1>
<h4 align="center">Cloudreve 4.18.0 fork with zero-copy offline download relocation.</h4>

> **Unofficial variant based on [Cloudreve 4.18.0](https://github.com/cloudreve/cloudreve/releases/tag/4.18.0).** When a completed offline download on the master node is transferred to local storage on the same filesystem, this fork links the file into place and removes the source after upload completion instead of copying its contents. Transfers across filesystems fall back to copying. Seeding tasks and other storage paths retain the original behavior.

<p align="center">
  <a href="https://github.com/Corezcy/cloudreve/actions/workflows/fork-release.yml">
    <img src="https://github.com/Corezcy/cloudreve/actions/workflows/fork-release.yml/badge.svg?branch=feat%2F4.18.0-zero-copy-remote-download"
         alt="Fork build">
  </a>
  <a href="https://github.com/Corezcy/cloudreve/releases">
    <img src="https://img.shields.io/github/v/release/Corezcy/cloudreve?include_prereleases" alt="Fork release" />
  </a>
</p>
<p align="center">
  <a href="https://cloudreve.org">Homepage</a> •
  <a href="https://demo.cloudreve.org">Try it</a> •
  <a href="https://github.com/cloudreve/cloudreve/discussions">Discussion</a> •
  <a href="https://docs.cloudreve.org">Documents</a> •
  <a href="https://github.com/Corezcy/cloudreve/releases">Download this fork</a> •
  <a href="https://github.com/cloudreve/cloudreve">Upstream</a> •
  <a href="https://t.me/cloudreve_official">Telegram</a> •
  <a href="https://discord.com/invite/WTpMFpZT76">Discord</a>
</p>

![Screenshot](https://raw.githubusercontent.com/cloudreve/docs/master/images/homepage.png)

## :sparkles: Features

- :cloud: Support storing files into Local, Remote node, OneDrive, S3 compatible API, Qiniu Kodo, Aliyun OSS, Tencent COS, Huawei Cloud OBS, Kingsoft Cloud KS3, Upyun.
- :outbox_tray: Upload/Download in directly transmission from client to storage providers.
- 💾 Integrate with Aria2/qBittorrent to download files in background, use multiple download nodes to share the load.
- 📚 Compress/Extract/Preview archived files, download files in batch.
- 💻 WebDAV support covering all storage providers.
- :zap:Drag&Drop to upload files or folders, with parallel resumable upload support.
- :card_file_box: Extract media metadata from files, search files by metadata or tags.
- :family_woman_girl_boy: Multi-users with multi-groups.
- :link: Create share links for files and folders with expiration date.
- :eye_speech_bubble: Preview videos, images, audios, ePub files online; edit texts, diagrams, Markdown, images, Office documents online.
- :art: Customize theme colors, dark mode, PWA application, SPA, i18n.
- :rocket: All-in-one packaging, with all features out of the box.
- 🌈 ... ...

## :hammer_and_wrench: Deploy

To deploy Cloudreve, you can refer to [Getting started](https://docs.cloudreve.org/overview/quickstart) for a quick local deployment to test.

When you're ready to deploy Cloudreve to a production environment, you can refer to [Deploy](https://docs.cloudreve.org/overview/deploy/) for a complete deployment.

## :gear: Build

Please refer to [Build](https://docs.cloudreve.org/overview/build/) for how to build Cloudreve from source code.

This fork also [builds and tests on GitHub Actions](https://github.com/Corezcy/cloudreve/actions/workflows/fork-release.yml) and publishes binaries under [Releases](https://github.com/Corezcy/cloudreve/releases). It does not publish a Docker image.

## :rocket: Contributing

If you're interested in contributing to Cloudreve, please refer to [Contributing](https://docs.cloudreve.org/api/contributing/) for how to contribute to Cloudreve.

## :alembic: Stacks

- [Go](https://golang.org/) + [Gin](https://github.com/gin-gonic/gin) + [ent](https://github.com/ent/ent)
- [React](https://github.com/facebook/react) + [Redux](https://github.com/reduxjs/redux) + [Material-UI](https://github.com/mui-org/material-ui)

## :scroll: License

GPL V3
