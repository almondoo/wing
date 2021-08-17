# タスク管理システム OSS  

## 開発環境手順  
1. make start-up
2. make bash
3. npm i  
4. npm run dev

## 技術スタック  
React 17.0.2  
Next 11.0.1  
SWR 0.5.6  
ホスティング  
vercel  

Go 1.5  
echo 4系  

## ディレクトリ構造  

```
├── src
│   ├── pages ルーティング
│   │   └── [PageName]
│   │           ├── index.tsx
│   │           └── [PageName]
│   ├── components
│   │   ├── identities ページ単位で必要なパーツ
│   │   │   └── [Page Name]
│   │   │       └──[Component Name]
│   │   │          ├── index.tsx
│   │   │          └── style.tsx
│   │   │
│   │   ├── originals 特定の用途でしか使われない
│   │   ├── templates レイアウト
│   │   └── uiParts 共通UIパーツ
│   │       └── [Component Name]
│   │           ├── index.tsx
│   │           └── style.tsx
│   │
│   ├── models 型定義
│   ├── pageStyles pagesに直接使う なくなる予定
│   ├── styles 全体のスタイル
│   └── utils 他に必要なもの
```
