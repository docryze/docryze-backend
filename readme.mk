docryze-backend/
├── cmd/                  # 应用入口
│   └── main.go           # 主启动文件
├── internal/             # 私有应用代码（禁止外部导入）
│   ├── config/           # 配置管理
│   ├── controllers/      # HTTP 控制器层
│   ├── models/           # 数据模型定义
│   ├── repositories/     # 数据访问层（GORM 操作）
│   ├── services/         # 业务逻辑层
│   ├── routes/           # 路由定义
│   ├── middleware/       # Gin 中间件
│   └── utils/            # 工具函数
├── pkg/                  # 公共库（可对外暴露）
├── api/                  # API 协议定义（Swagger/Protobuf）
├── database/             # 数据库脚本
│   └── migrations/       # 迁移文件
├── scripts/              # 部署/运维脚本
├── configs/              # 配置文件
│   └── config.yaml       # 主配置文件
├── tests/                # 测试代码
├── go.mod
└── go.sum