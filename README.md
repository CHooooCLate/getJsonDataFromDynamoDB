# getJsonDataFromDynamoDB
## AWSの構成
### Lambda


### APIGateway
#### メソッドリクエスト
なし
#### 統合リクエスト
なし
#### 統合レスポンス
200
##### ヘッダーのマッピング
なし

##### マッピングテンプレート
```
#set($inputRoot = $input.path('$'))

$inputRoot
```

#### メソッドレスポンス
