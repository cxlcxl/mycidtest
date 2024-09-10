package constant

type Platform int
type Media int

var (
	// BasePath 系统跟路径
	BasePath string
)

const (
	LoginKey     = "__sys_login_key__"
	RequestIdKey = "__sys_request_id_key__"

	PlatformJd  Platform = 1
	PlatformPdd Platform = 2
	PlatformTb  Platform = 3

	MediaTypeIntTT  Media = 1
	MediaTypeIntKS  Media = 2
	MediaTypeIntGDT Media = 3

	MediaTypeStringTT  = "toutiao"
	MediaTypeStringKS  = "kuaishou"
	MediaTypeStringGDT = "gdt"
)
