# Amazon Cognito へのユーザー移行の Lambda トリガー

---

- ユーザーによるパスワードによるサインイン時、またはユーザーがパスワードを忘れた場合のフローにいる時に、そのユーザーがユーザープールに存在しない場合、Amazon Cognito は、このトリガーを呼び出します。
- Lambda 関数が正常に値を返した後でユーザーをユーザープールに作成します。

![signin.png](assets/img/signin.png)

![password_reset.png](assets/img/password_reset.png)

