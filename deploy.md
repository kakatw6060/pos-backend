# POS Backend Deployment Guide

## 🚀 快速部署路徑 (建議使用 Render.com - 免費)

1. **將代碼上傳到 GitHub**:
   - 創建一個 private 或 public repo。
   - 將 `/opt/data/pos-backend` 下的所有文件 push 上去。

2. **連接 Render**:
   - 登入 [Render.com](https://render.com) (用 GitHub 登入)。
   - 選擇 **"New"** $\rightarrow$ **"Web Service"**。
   - 選擇你啱啱創建嘅 GitHub Repo。

3. **配置設定**:
   - **Runtime**: `Docker` (Render 會自動偵測 Dockerfile)。
   - **Instance Type**: `Free`。
   - 點擊 **"Create Web Service"**。

4. **攞網址**:
   - 等約 2 分鐘部署完成，Render 會喺頁面頂部畀個 `https://pos-backend-xxx.onrender.com` 嘅網址你。

---

## ⚠️ 重要：關於數據持久化 (Persistence)

由於你目前使用 **SQLite**，而 Render 嘅 Free Tier 係 **Ephemeral File System** (每次 restart 或 deploy 完之後，`.db` 檔案會被重設/刪除)。

**解決方案 (這就是我之前加抽象層的原因)：**

1. **方案 A (簡單但唔持久)**: 直接用 SQLite，適合測試 UI/API，唔適合儲存正式訂單。
2. **方案 B (生產級 - 強烈建議)**:
   - 喺 Render 創建一個 **"Free PostgreSQL"** 數據庫。
   - 喺 Web Service 嘅 **Environment Variables** 加入：
     - `DB_TYPE=postgres`
     - `DB_URL=postgres://user:password@host:port/dbname`
   - 我可以幫你寫一個簡單嘅 `database.go` 修改版，令佢讀取呢啲 Env 變量就自動由 SQLite 換成 Postgres。

如果你決定用方案 B，話我知，我 10 秒鐘幫你改好代碼！
