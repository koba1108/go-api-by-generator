### デプロイ手順

1. コンテナイメージを生成
```
docker-compose build --no-cache swagger-ui
```
2. タグ付け
```
docker tag go-api-by-generator_swagger-ui gcr.io/ykoba-dev/swagger-ui:0.0.1
```
3. コンテナレジストリへアップロード
```
docker push gcr.io/ykoba-dev/swagger-ui:0.0.1
```
4. アップロードしたGCPプロジェクトのCloudRun画面からサービスを選択
https://console.cloud.google.com/run?project=ykoba-dev

5. 新しいバージョンをデプロイをクリックし、アップロードしたイメージを選択後に「デプロイ」をクリック

※ swagger-mockはこっち
```
docker-compose build --no-cache swagger-mock
docker tag go-api-by-generator_swagger-mock gcr.io/ykoba-dev/swagger-mock:0.0.1
docker push gcr.io/ykoba-dev/swagger-mock:0.0.1
```
