package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}

// | Тег              | Описание                                                                  |
// | ---------------- | ------------------------------------------------------------------------- |
// | `primaryKey`     | Обозначает поле как первичный ключ (`PRIMARY KEY`)                        |
// | `type`           | Явное указание типа столбца в базе (`varchar(100)`, `uuid`, `int` и т.д.) |
// | `size`           | Размер поля (например, для строк)                                         |
// | `unique`         | Уникальное ограничение на поле                                            |
// | `uniqueIndex`    | Создаёт уникальный индекс                                                 |
// | `index`          | Создаёт индекс                                                            |
// | `not null`       | Поле не может быть `NULL`                                                 |
// | `default`        | Значение по умолчанию                                                     |
// | `autoIncrement`  | Автоматическая инкрементация (например, для числовых ID)                  |
// | `column`         | Указать имя колонки в таблице                                             |
// | `embedded`       | Вложенная структура (встраивание)                                         |
// | `embeddedPrefix` | Префикс для полей вложенной структуры                                     |
// | `serializer`     | Кастомный сериализатор (например, JSON для сложных полей)                 |
// | `-`              | Игнорировать поле (не сохранять в базу)                                   |
// | `foreignKey`     | Поле, являющееся внешним ключом (используется в связях)                   |
// | `references`     | Указать связанное поле в другой таблице (для ассоциаций)                  |
// | `constraint`     | Управление ограничениями, например, `OnDelete:CASCADE`                    |
