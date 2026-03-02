# Предварительная подготовка

В начале необходимо выполнить команды в консоли:
```diff
+ go mod init [Folder]
```
Далее получим необходимые файлы проекта:
```diff
+ go get gorm.io/gorm
+ go get gorm.io/driver/postgres
```
---
# **🙂‍↔️Написание самого кода и что к чему**
### Создание структуры пользователя:
```go
type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Age       int       `json:"id"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	CreatedAt time.Time `json:"created_at"`
}
```
> В структуре используется только 1 импорт для контроля времени ***time***

