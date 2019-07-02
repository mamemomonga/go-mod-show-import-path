# go-mod-show-import-path

Search go.mod and show absolute import path from relative path.
This is useful when using a package you created yourself.

# INSTALL

	$ go install github.com/mamemomonga/go-mod-show-import-path

# USAGE

	$ go mod init github.com/mamemomonga/hoge01
	$ goimportabs .
	github.com/mamemomonga/hoge01

	$ mkdir -p hoge/moge
	$ cd hoge/moge
	$ goimportabs .
	github.com/mamemomonga/hoge01/hoge/moge

	$ goimportabs ../..
	github.com/mamemomonga/hoge01

	$ goimportabs ../tada
	github.com/mamemomonga/hoge01/hoge/tada

# LICENSE

MIT

