package analyzer

import (
	"encoding/json"
	"go/ast"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

// Analyzer - основной анализатор, который проверяет лог-сообщения по правилам.
var Analyzer = &analysis.Analyzer{
	Name: "loglinter",
	Doc:  "checks log messages rules",
	Run:  run,
}

// -------------- Config - настройки чувствительный слов --------------
type Config struct {
	SensitiveWords []string `json:"sensitive_words"`
}
// Заводской конфиг: "password", "api_key", "token", "secret"
var defaultConfig = Config{
	SensitiveWords: []string{"password", "api_key", "token", "secret"},
}
// loadConfig читает конфиг из файла config_analyzer.json,
func loadConfig() Config {
	// Определяем путь к текущему файлу (analyzer.go) через runtime.Caller.
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return defaultConfig
	}

	// Папка, где лежит analyzer.go.
	dir := filepath.Dir(filename)

	// Файл конфига в этой папке.
	path := filepath.Join(dir, "config_analyzer.json")

	data, err := os.ReadFile(path)
	if err != nil {
		// Файла нет или не читается — используем дефолт.
		return defaultConfig
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		// Ошибка разбора — тоже откат к дефолту.
		return defaultConfig
	}

	if len(cfg.SensitiveWords) == 0 {
		return defaultConfig
	}

	return cfg
}
// --------------------------------------------------------------------

// run — точка входа анализа для одного пакета.
func run(pass *analysis.Pass) (any, error) {
	cfg := loadConfig()

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			checkLogCall(pass, call, cfg)
			return true
		})
	}

	return nil, nil
}

// checkLogCall - проверяет, что вызов похож на лог-запись, и если да,
func checkLogCall(pass *analysis.Pass, call *ast.CallExpr, cfg Config) {
	// Ожидаем X.Y(...): логгер.метод(...)
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	// Имя метода: Info / Error / Warn / Debug.
	method := sel.Sel.Name
	if method != "Info" && method != "Error" && method != "Warn" && method != "Debug" {
		return
	}

	// Достаём строку сообщения.
	msg := getMsg(call.Args)
	if msg == "" {
		return
	}

	// Проверяем все правила из ТЗ, с учётом конфига.
	checkRules(pass, sel.Pos(), msg, cfg)
}

// getMsg - достаём строковый литерал
func getMsg(args []ast.Expr) string {
	if len(args) == 0 {
		return ""
	}

	lit, ok := args[0].(*ast.BasicLit)
	if !ok {
		return ""
	}
	if lit.Kind != token.STRING {
		return ""
	}
	// Убираем кавычки.
	return strings.Trim(lit.Value, `"`)
}

// checkRules - применяет все правила к одному сообщению лога.
func checkRules(pass *analysis.Pass, pos token.Pos, msg string, cfg Config) {
	if msg == "" {
		return
	}

	// 1. Начинается со строчной буквы.
	r, _ := utf8.DecodeRuneInString(msg)
	if unicode.IsLetter(r) && unicode.IsUpper(r) {
		pass.Reportf(pos, "❌ log message must start with lowercase letter: %q", msg)
	}

	// 2. Только АНГЛ (запрещаем РУ).
	for _, c := range msg {
		if c >= 0x0400 && c <= 0x04FF {
			pass.Reportf(pos, "❌ log message must be in English only (no Cyrillic): %q", msg)
			break
		}
	}

	// 3. Нет спецсимволов/эмодзи. Вообще тут можно по айдишникам сделать,
	// но к сожалению времени разбираться в силу учёбы в тот момент не было, сделаю где-то 15числа,
	// как раз нужно будет ставить Линукс, ибо на Винде с плагинами беда
	badRunes := []rune{'!', '?', '🚀'}
	for _, br := range badRunes {
		if strings.ContainsRune(msg, br) {
			pass.Reportf(pos, "❌ log message must not contain special chars or emojis: %q", msg)
			break
		}
	}

	// 4. Нет чувствительных слов - используем cfg.SensitiveWords из конфига.
	msgLow := strings.ToLower(msg)
	for _, w := range cfg.SensitiveWords {
		if strings.Contains(msgLow, strings.ToLower(w)) {
			pass.Reportf(pos, "❌ log message must not contain sensitive data (%s): %q", w, msg)
			break
		}
	}
}
