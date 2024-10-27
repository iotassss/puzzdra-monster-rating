# PUZZDRA MONSTER RATING
パズドラモンスターの評価データベース

## DB
```bash
# ローカル環境でのテスト

# 初期化
docker run --name my-mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=puzzdra_db -e MYSQL_USER=testuser -e MYSQL_PASSWORD=testpw -p 3306:3306 -d mysql:8.0

# 初期化（Volumeつき）
docker run --name my-mysql \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=puzzdra_db \
  -e MYSQL_USER=testuser \
  -e MYSQL_PASSWORD=testpw \
  -p 3306:3306 \
  -v /path/to/your/local/directory:/var/lib/mysql \
  -d mysql:8.0

# 接続
mysql -u root -p -h 127.0.0.1 -P 3306 puzzdra_db
```
