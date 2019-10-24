# go-api-by-generator
go-api created by xo, swagger

## api generate command
```shell script
$ sh gen-gin-server.sh 
```

更新時はgenerate配下のファイルは変更するのではなく、<br/>
変更diffを手作業で取り込む必要がありそう。

## model generate command
```shell script
docker-compose exec api bash
xo mysql://root:pass1234@db/public -o models
```

## generate時に関数名が２つ作られている場合

${table}.xo.go から該当の関数を消す以外なさそう
