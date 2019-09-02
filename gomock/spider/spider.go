package spider

//duan@DESKTOP-G4D1C1K MINGW64 /c/Go/projects/GoStuff/gomock/spider (master)
//$ mockgen -source=spider.go > mock/mock_spider.go
//2019/08/30 14:02:54 Loading input failed: loading package failed

type Spider interface {
	GetVersion() string
}
