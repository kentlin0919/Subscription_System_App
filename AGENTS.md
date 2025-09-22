# Repository Guidelines

## Project Structure & Module Organization
- Core entry point lives in `main.go`. Keep the HTTP handlers, services, and data access code in well-named packages under `internal/` or `pkg/` when the application expands.
- Group shared utilities under `pkg/` and domain-specific logic under `internal/<domain>` so that dependencies flow inward.
- Place integration fixtures, sample payloads, or SQL migrations under a top-level `assets/` directory; keep test resources alongside the corresponding test files.

## Build, Test, and Development Commands
- `go run main.go` starts the subscription service locally with hot feedback during development.
- `go build ./...` produces a compiled binary and quickly surfaces compilation issues across all packages.
- `go test ./...` executes the full Go test suite; prefer running this before opening a pull request.
- `go vet ./...` catches common correctness issues; run it when introducing non-trivial changes.

## Coding Style & Naming Conventions
- Use tabs for indentation and rely on `gofmt` (or `go fmt ./...`) before committing. Import ordering should follow `goimports` conventions.
- Package names should be short, lowercase, and match their directory. Exported identifiers stay in UpperCamelCase; unexported ones in lowerCamelCase.
- Keep files focused: limit each file to one primary type or responsibility, and document exported functions with complete sentences.

## Testing Guidelines
- Write tests with Go’s standard `testing` package; name them `TestFeature_Scenario` for clarity.
- Co-locate unit tests in `<file>_test.go` next to the code. Use table-driven tests for permutations and add benchmarks (`BenchmarkX`) when performance matters.
- Target meaningful coverage of critical paths (billing cycles, plan transitions). Add regression tests whenever you fix a bug.

## Commit & Pull Request Guidelines
- Commit messages follow the conventional imperative mood (e.g., `Add customer renewal scheduler`) and focus on what the change delivers.
- Squash local work-in-progress commits before requesting review. Reference issue IDs in either the title or first line of the description.
- Pull requests should describe the change, outline testing performed (`go test ./...`, manual QA steps), and attach screenshots or logs when updating behaviour.

## Configuration & Security Tips
- Store environment-specific secrets in `.env.local` or your shell profile; never commit credentials. Document required variables in the pull request body.
- Review dependencies introduced with `go list -m all` and prefer standard-library utilities unless an external package adds clear value.


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