# go-rqh-webserver

* HTTP Request Headerを表示するシンプルなウェブサーバです。
* dep, packr, html/templateの習作

# 実行例

	$ ./rqh-webserver-darwin-amd64 -listen 0.0.0.0:8000

# 開発

## ビルド環境

事前に必要なもの

* make
* go

導入されてない場合導入されるもの
		
* dep
* packr

### 必要なパッケージの導入

* dep eusure が実行されます。dep, packrが導入されていない場合は導入されます。
* GOPATHが正しく設定されている必要があります。

コマンド

	$ make deps

### 実行

コマンド

	$ make run

### ビルド

コマンド

	$ make

### リリース向けビルド

* 公開用のバイナリが生成されます。
* Dockerが必要です。

コマンド

	$ make release


