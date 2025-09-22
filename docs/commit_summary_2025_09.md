# 2025 年 09 月提交摘要

## 2025-09-22 feat(http): adopt gin router
- main.go: 改為應用進入點，注入依賴並使用 Gin 監聽埠 8080。
- internal/adapter/http: 將路由與 Handler 切換至 Gin，保留用戶建立端點。
- README.md: 更新版本資訊、檔案結構與時序圖，新增 Gin 相關說明與文件連結。
- docs: 新增 `router_flow.md` 與 `GEMINI.md` 記錄 Gin 路由流程與專案整理。

## 2025-09-20 chore(deps): install gin framework
- go.mod: 新增 Gin 及其間接相依套件需求，準備後續改用 Gin 處理 HTTP。
- go.sum: 新增相依套件的雜湊記錄以確保模組可重現安裝。
- README.md: 補充 Gin 版本資訊、同步檔案結構並說明如何維護 Gin。
