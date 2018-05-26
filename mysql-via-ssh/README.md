# mysql-via-ssh
公開鍵認証でssh接続しmysqlに接続する

## How to setup
- .envファイルに環境情報をセット

```
$ cp .env.sample .env
$ vi .env
```
## Ref
### BaseCode
- [vinzenz/dial-mysql-via-ssh.go](https://gist.github.com/vinzenz/d8e6834d9e25bbd422c14326f357cce0)
- [Go言語で認証鍵を使ってSSHの接続を行う](https://saitodev.co/article/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E8%AA%8D%E8%A8%BC%E9%8D%B5%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6SSH%E3%81%AE%E6%8E%A5%E7%B6%9A%E3%82%92%E8%A1%8C%E3%81%86)

### Debugging
- [GoLang ssh : Getting “Must specify HosKeyCallback” error despite setting it to nil
](https://stackoverflow.com/questions/44269142/golang-ssh-getting-must-specify-hoskeycallback-error-despite-setting-it-to-n)
- [error connecting to database with mysqldriver
](https://stackoverflow.com/questions/25244089/error-connecting-to-database-with-mysqldriver)