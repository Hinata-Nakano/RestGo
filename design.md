src/
├── domain/                    # ドメイン層（ビジネスロジックの核心）
│   ├── user/                 # ユーザードメイン
│   │   ├── user.go           # ユーザーエンティティ/値オブジェクト
│   │   ├── user_repository.go # リポジトリインターフェース（純粋な関数型）
│   │   └── user_service.go   # ドメインサービス（純粋関数）
│   ├── tweet/                # ツイートドメイン
│   │   ├── tweet.go          # ツイートエンティティ/値オブジェクト
│   │   ├── tweet_repository.go # リポジトリインターフェース
│   │   └── tweet_service.go  # ドメインサービス（純粋関数）
│   └── errors.go             # ドメインエラー
│
├── application/              # アプリケーション層（ユースケース）
│   ├── user/
│   │   └── user_usecase.go   # ユーザー関連のユースケース
│   └── tweet/
│       └── tweet_usecase.go  # ツイート関連のユースケース
│
├── infrastructure/           # インフラ層（副作用の実装）
│   ├── database/
│   │   ├── db.go             # DB接続
│   │   └── migrations/       # マイグレーション
│   ├── repository/           # リポジトリ実装
│   │   ├── user_repository_impl.go
│   │   └── tweet_repository_impl.go
│   └── logger/
│       └── logger.go
│
├── presentation/             # プレゼンテーション層
│   ├── http/                 # HTTPハンドラー
│   │   ├── router.go
│   │   ├── middleware.go
│   │   ├── user_handler.go
│   │   └── tweet_handler.go
│   └── dto/                  # データ転送オブジェクト
│       ├── user_dto.go
│       └── tweet_dto.go
│
├── shared/                   # 共通ユーティリティ
│   ├── types.go
│   └── validation.go
│
└── main.go                   # エントリーポイント