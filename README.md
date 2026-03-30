# whros-cli

个人事务管理 CLI 工具，基于 lark-cli 架构设计。

## 功能特性

- **任务管理** - 添加、查看、完成、删除任务
- **日历管理** - 添加、查看、删除日程事件
- **笔记管理** - 添加、查看、搜索、删除笔记
- **配置管理** - 查看和修改配置

## 安装

### 一键安装 (推荐)

```bash
curl -fsSL https://gitee.com/tuoju/whros-cli/raw/main/install.sh | bash
```

### Go install

```bash
go install github.com/tuoju/whros-cli@latest
```

### Homebrew (macOS)

```bash
brew tap tuoju/whros
brew install whros
```

## 构建

```bash
# 安装依赖
go mod tidy

# 构建
go build -o bin/whros .

# 或使用 Makefile
make build
```

## 使用

### 任务管理

```bash
# 添加任务
./bin/whros task add "完成任务" --priority high
./bin/whros task add "日常任务" --priority medium --tag "工作"

# 查看任务列表
./bin/whros task list

# 查看所有任务（包括已完成）
./bin/whros task list --all

# 标记任务完成
./bin/whros task done <task_id>

# 删除任务
./bin/whros task delete <task_id>
```

**参数说明：**
- `--priority, -p` - 优先级 (high/medium/low)
- `--tag, -t` - 标签

### 日历管理

```bash
# 添加日程
./bin/whros calendar add "会议" --time "2024-01-15 14:00" --duration 60

# 查看所有日程
./bin/whros calendar list

# 查看指定日期的日程
./bin/whros calendar list --date 2024-01-15

# 删除日程
./bin/whros calendar delete <event_id>
```

**参数说明：**
- `--time, -t` - 时间 (格式: YYYY-MM-DD HH:MM)
- `--duration, -d` - 持续时间（分钟）
- `--date` - 筛选特定日期

### 笔记管理

```bash
# 添加笔记
./bin/whros note add "想法" --content "这是笔记内容" --tag "灵感"

# 查看所有笔记
./bin/whros note list

# 搜索笔记
./bin/whros note search "关键词"

# 删除笔记
./bin/whros note delete <note_id>
```

**参数说明：**
- `--content, -c` - 笔记内容
- `--tag, -t` - 标签

### 配置管理

```bash
# 查看当前配置
./bin/whros config show
```

## 数据存储

数据默认存储在 `.whros/` 目录：

```
.whros/
├── tasks.json    # 任务数据
├── calendar.json # 日历数据
├── notes.json    # 笔记数据
└── config.yaml   # 配置文件
```

## 项目结构

```
whros-cli/
├── main.go              # 程序入口
├── Makefile             # 构建脚本
├── go.mod               # Go 模块定义
├── cmd/                 # CLI 命令层
│   ├── root.go          # 根命令
│   ├── task.go         # 任务子命令
│   ├── calendar.go      # 日历子命令
│   ├── note.go         # 笔记子命令
│   └── config.go       # 配置子命令
└── internal/            # 核心业务模块
    ├── task/            # 任务模块
    ├── calendar/        # 日历模块
    ├── note/            # 笔记模块
    └── config/          # 配置模块
```

## 开发

```bash
# 运行测试
go test ./...

# 清理构建
make clean

# 完整构建
make all
```

## Gitee Access Token 获取

1. 登录 Gitee：https://gitee.com
2. 进入个人设置 → 私人令牌 → 获取私人令牌
3. 点击 "生成新令牌"，填写备注信息
4. 勾选权限：
   - `projects` - 项目操作
   - `releases` - 发布版本操作
5. 点击提交，复制生成的 Token

## Token 配置

Token 文件路径：`/opt/settings/whros-local.properties`

内容格式：
```properties
gitee.access-token=你的Gitee_ACCESS_TOKEN
```