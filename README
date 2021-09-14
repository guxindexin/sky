<p align="center">
  <img src="https://www.fdevops.com/wp-content/uploads/2020/09/1599039924-ferry_log.png">
</p>


<p align="center">
  <a href="https://github.com/lanyulei/sky">
    <img src="https://www.fdevops.com/wp-content/uploads/2020/07/1595067271-badge.png">
  </a>
  <a href="https://github.com/lanyulei/sky">
    <img src="https://www.fdevops.com/wp-content/uploads/2020/07/1595067272-apistatus.png" alt="license">
  </a>
  <a href="https://github.com/lanyulei/sky">
    <img src="https://www.fdevops.com/wp-content/uploads/2020/07/1595067269-donate.png" alt="donate">
  </a>
</p>

本系统使用最新的 Vue3 及相关技术栈，实现了菜单及页面按钮的权限管控，基于 Casbin 实现了后端 API 接口的管控，不进行过度的封装，代码简洁易懂，方便二次开发及当成后端管理平台脚手架使用，数据库暂时仅支持 postgres 。

演示站点：[http://fdevops.com:8088](http://fdevops.com:8088)

文档地址：[https://www.fdevops.com/docs/sky-30965](https://www.fdevops.com/docs/sky-30965)

## 技术栈

### 后端

* Go
* Gin
* Gorm

### 前端

* Vue3
* CompositionAPI
* typescript
* vite
* element plus

### 数据库

* postgresql

### 其他

* Casbin
* 等等，其他的开源 Go 依赖

## 使用

**前端**

```
# 克隆代码
git clone https://github.com/lanyulei/sky-ui.git

# 安装依赖
cd sky-ui
npm install

# 启动
npm run dev

# 编译
npm run build
```

**后端**

```
# 克隆代码
git clone https://github.com/lanyulei/sky

# 同步数据
go run main.go migrate -c config/settings.yml

# 启动服务
go run main.go server -c config/settings.yml 
```

## License

[MIT License](https://github.com/lanyulei/sky/blob/master/LICENSE)
