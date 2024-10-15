# Amazon Cognito の Lambda トリガー


## ユーザ移行トリガー

- ユーザーによるパスワードによるサインイン時、そのユーザーがユーザープールに存在しない場合、Amazon Cognito は、このトリガーを呼び出します。
- その後、Lambda 関数が正常に値を返した場合、ユーザーをユーザープールに作成しようとします。
  - フロー詳細: https://docs.aws.amazon.com/ja_jp/cognito/latest/developerguide/cognito-user-pools-import-using-lambda.html
- ただし、Google などのソーシャルログインを含むフェデレーティッドサインインは、このトリガーを呼び出しません。


## サインアップ前トリガー

- 以下の場合に Amazon Cognito はこのトリガーを呼び出します。
  - ユーザーがサインアップする場合
  - ユーザーがソーシャルプロバイダーを使用してログインする際、ユーザープールに同じ認証方式のユーザーが存在しない場合
  - 管理者がユーザーをユーザープールに追加しようとしている場合
    - Cognito によるユーザ移行の際
    - Cognito SDK によるユーザーの作成の際

