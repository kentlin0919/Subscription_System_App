# GEMINI Analysis: Subscription System App

## Project Overview

This project is a Go-based backend service for a subscription system. It is built using Go version 1.21 and the Gin web framework for handling HTTP requests.

The application follows a **Clean Architecture** pattern, clearly separating concerns into distinct layers:

*   **`main.go`**: The application's entry point. It initializes and wires together all the components (dependency injection).
*   **`internal/domain`**: Contains the core business models, such as the `User` struct. This layer has no external dependencies.
*   **`internal/usecase`**: Implements the core business logic. For example, `UserCreator` handles the process of creating a new user. It depends on interfaces defined within this layer.
*   **`internal/adapter`**: Connects the use cases to external agencies. `internal/adapter/http` contains the Gin router and HTTP handlers (e.g., `UserHandler`) that translate HTTP requests into calls to the use case layer.
*   **`internal/infra`**: Contains implementations for external concerns like databases and logging. `internal/infra/db` provides a PostgreSQL-based implementation of the user repository.

The service exposes a REST API, starting with an endpoint to create users (`POST /users`). The database logic for PostgreSQL is defined but contains `TODO` markers, indicating it is not fully implemented.

## Building and Running

Standard Go toolchain commands are used for managing the project.

*   **Run the application (for development):**
    ```bash
    go run main.go
    ```
    The service will start and listen on port `:8080`.

*   **Build the project:**
    ```bash
    go build ./...
    ```

*   **Run tests:**
    ```bash
    go test ./...
    ```

*   **Format code:**
    Before committing, ensure code is formatted according to Go standards.
    ```bash
    go fmt ./...
    ```

## Development Conventions

*   **Architecture**: Adhere to the existing Clean Architecture structure. New features should be implemented by adding to the appropriate layers (domain, usecase, adapter, infra).
*   **Dependency Injection**: Dependencies are explicitly passed into components via their constructors (e.g., `NewUserHandler` receives a `UserCreator`). This is all wired together in `main.go`.
*   **Interfaces**: Decoupling between layers is achieved with interfaces. For example, the `UserCreator` use case depends on a `UserRepository` interface, not a concrete database implementation.
*   **Routing**: HTTP routes are defined in `internal/adapter/http/router.go` using the Gin framework. Handlers are implemented in the same directory (e.g., `user_handler.go`).
*   **Configuration**: Currently, configuration is hardcoded (e.g., the server port). This could be extracted into a separate configuration management system.
*   **Error Handling**: The `README.md` suggests that proper error handling and data validation are yet to be fully implemented and should be added.


## 基本檔案結構

yourapp/          // 組裝應用、載入環境與啟動伺服器
├─ internal/
│  ├─ domain/               // 實體與領域介面（純 Go，無外部依賴）
│  │  └─ user.go
│  ├─ usecase/              // Use Cases（業務規則），依賴 domain 抽象
│  │  └─ user_create.go
│  ├─ adapter/              // 介面層：HTTP handlers、gRPC、CLI、PubSub 等
│  │  └─ http/
│  │     ├─ router.go
│  │     └─ user_handler.go
│  └─ infra/                // 基礎設施：DB、外部 API、Cache 的實作
│     ├─ db/
│     │  ├─ postgres.go
│     │  └─ user_repo_pg.go
│     └─ log/
│        └─ logger.go
├─ pkg/                     // 可重用套件（非業務），或跨專案共用工具
└─ go.mod
└─ main.go


## 回覆語言
- 繁體中文

## README 書寫條件
- 要有詳細的版本資訊
- 詳細的檔案結構
- 針對每個功能畫出時序圖
- 針對每個系統的環境安裝的詳細流程以及指令
- 每次執行時都對 README.md 進行更新確保符合現在專案


## 提交與 Pull Request 規範
- Commit：簡短、命令式、標明範疇（如：`fix(chart): clamp pan at edges`）
- 語言 ： 請以繁體中文書寫
- 請將相關變更分組，避免不相關的重構
- PR 必須包含：
  - 變更摘要與原因
  - UI 變更請附截圖/GIF
  - 手動測試步驟與影響模組/路徑
  - 若適用請附上相關 issue 或任務 ID
- 分支命名：`feature/<name>`、`fix/<name>`、`chore/<name>`
- 每個 PR 需至少一位審核者
- 請保持 PR 規模小且聚焦（盡量少於 300 行）
- 請將所有commit 經過逐行分析後放進/docs中的commit_summary_2025_{當月月份}.md 的文件中
- commit_summary_2025_{當月月份}.md 請以時間新到舊排序

## 使用套件的版本
- 確保環境安裝腳本有確認是否安裝必要套件
