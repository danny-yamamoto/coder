- 問
file コマンドの結果が文字列として与えられるので、それを parse して、ファイル名 → mime type という map (treemap, hashmap, etc.) を作る関数を作ってください
file コマンドの結果は、一行に ファイル名: mime type の形で、これが複数行あります
ファイル名と mime type の間は空白（スペース）とします。
ファイル名は、ファイル名として使えるものなら来る可能性があります
ファイル名に空白がありうることを想定してください。
例外として、ファイル名には改行はないものとして構いません。
mime type は、[a-z0-9-_.]+/[a-z0-9-_.]+ しか来ないと想定して構いません
ファイル名に重複はありません
go での関数の例
```go
func parse(output string) map[string]string {
    // Write your code
}
```

- note
   - `len`
