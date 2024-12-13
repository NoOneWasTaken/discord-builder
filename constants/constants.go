package constants

//goland:noinspection ALL
var (
	NodeJSDefaultModules      = [...]string{"_http_agent", "_http_client", "_http_common", "_http_incoming", "_http_outgoing", "_http_server", "_stream_duplex", "_stream_passthrough", "_stream_readable", "_stream_transform", "_stream_wrap", "_stream_writable", "_tls_common", "_tls_wrap", "assert", "assert/strict", "async_hooks", "buffer", "child_process", "cluster", "console", "constants", "crypto", "dgram", "diagnostics_channel", "dns", "dns/promises", "domain", "events", "fs", "fs/promises", "http", "http2", "https", "inspector", "inspector/promises", "module", "net", "os", "path", "path/posix", "path/win32", "perf_hooks", "process", "punycode", "querystring", "readline", "readline/promises", "repl", "stream", "stream/consumers", "stream/promises", "stream/web", "string_decoder", "sys", "timers", "timers/promises", "tls", "trace_events", "tty", "url", "util", "util/types", "v8", "vm", "wasi", "worker_threads", "zlib"}
	NodeJSBlacklistedKeywords = [...]string{"node_modules", "favicon.ico"}

	// Reset
	ANSIReset = "\033[0m"

	// Text styles
	ANSIBold      = "\033[1m"
	ANSIItalic    = "\033[3m"
	ANSIUnderline = "\033[4m"
	ANSIReverse   = "\033[7m"

	// Foreground colors
	ANSIBlack   = "\033[30m"
	ANSIRed     = "\033[31m"
	ANSIGreen   = "\033[32m"
	ANSIYellow  = "\033[33m"
	ANSIBlue    = "\033[34m"
	ANSIMagenta = "\033[35m"
	ANSICyan    = "\033[36m"
	ANSIWhite   = "\033[37m"

	// Bright foreground colors
	ANSIBrightBlack   = "\033[90m"
	ANSIBrightRed     = "\033[91m"
	ANSIBrightGreen   = "\033[92m"
	ANSIBrightYellow  = "\033[93m"
	ANSIBrightBlue    = "\033[94m"
	ANSIBrightMagenta = "\033[95m"
	ANSIBrightCyan    = "\033[96m"
	ANSIBrightWhite   = "\033[97m"

	// Background colors
	ANSIBackgroundBlack   = "\033[40m"
	ANSIBackgroundRed     = "\033[41m"
	ANSIBackgroundGreen   = "\033[42m"
	ANSIBackgroundYellow  = "\033[43m"
	ANSIBackgroundBlue    = "\033[44m"
	ANSIBackgroundMagenta = "\033[45m"
	ANSIBackgroundCyan    = "\033[46m"
	ANSIBackgroundWhite   = "\033[47m"

	// Bright background colors
	ANSIBackgroundBrightBlack   = "\033[100m"
	ANSIBackgroundBrightRed     = "\033[101m"
	ANSIBackgroundBrightGreen   = "\033[102m"
	ANSIBackgroundBrightYellow  = "\033[103m"
	ANSIBackgroundBrightBlue    = "\033[104m"
	ANSIBackgroundBrightMagenta = "\033[105m"
	ANSIBackgroundBrightCyan    = "\033[106m"
	ANSIBackgroundBrightWhite   = "\033[107m"
)
