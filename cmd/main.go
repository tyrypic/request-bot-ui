package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"ExBot/internal/adapter/db"
	"ExBot/internal/usecase"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// loadEnv загружает .env (если есть) и проверяет ошибки
func loadEnv() {
	// Попытка загрузить .env из текущей директории
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env файл не найден, полагаемся на реальные ENV")
	} else {
		log.Println("✅ Переменные загружены из .env")
	}
}

func runMigrations(dsn string) {
	fmt.Println(dsn)
	// путь "file://../migrations" относительный к рабочей директории
	m, err := migrate.New("file://../migrations", dsn)
	if err != nil {
		log.Fatalf("миграции: не смогли создать экземпляр: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("миграции: UP упали: %v", err)
	}
	log.Println("✅ Миграции применены")
}

func main() {
	// 0. Загружаем .env до всего остального
	loadEnv()

	// 1. Читаем настройки из окружения
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("✋ Ошибка: DATABASE_URL не задана")
	}
	log.Printf("🔗 Подключаемся к БД: %s", dsn)

	// 2. Контекст с таймаутом для инициализации БД
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. Создаём пул соединений
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("❌ Не удалось подключиться к БД: %v", err)
	}
	log.Println("✅ Подключение к БД установлено")
	defer func() {
		pool.Close()
		log.Println("👋 Соединение с БД закрыто")
	}()

	// 5. Запуск миграций
	runMigrations(dsn)

	// 4. Сеем администраторов
	adminID := os.Getenv("ADMIN_IDS")
	if adminID == "" {
		log.Fatalf("❌ Не обнаружен ADMIN_ID")
		return
	}
	log.Printf("👑 ID Администратора: %v", adminID)

	userRepo := db.NewUserRepo(pool)
	userSvc := usecase.NewUserService(userRepo)
	if err := userSvc.SeedAdmin(ctx, adminID); err != nil {
		log.Fatalf("❌ Не удалось зарегистрировать администратора: %v", err)
	}
	log.Println("✅ Администратор зарегистрирован")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
	// === Дальнейшие примеры использования ===

	// 6. Пример регистрации/обновления Telegram-пользователя
	// me := &domain.User{
	// 	TelegramID: 12345678,
	// 	Username:   "my_bot_user",
	// 	FirstName:  "Иван",
	// 	LastName:   "Иванов",
	// }
	// if err := userSvc.Register(ctx, me); err != nil {
	// 	log.Fatalf("❌ Register error: %v", err)
	// }
	// log.Printf("👤 Пользователь зарегистрирован: %+v", me)

	// // 7. Пример обновления профиля
	// me.RealName = "Иван Иванов"
	// me.Email = "ivan@example.com"
	// me.Age = 28
	// me.City = "Москва"
	// if err := userSvc.UpdateProfile(ctx, me); err != nil {
	// 	log.Fatalf("❌ UpdateProfile error: %v", err)
	// }
	// log.Printf("✏️ Профиль обновлён: %+v", me)

	// // 8. Работа с заявками
	// reqRepo := db.NewRequestRepo(pool)
	// reqSvc := usecase.NewRequestService(reqRepo)

	// // 8.1 Подать или обновить заявку (UPSERT)
	// req, err := reqSvc.Submit(ctx, me.ID)
	// if err != nil {
	// 	log.Fatalf("❌ SubmitRequest error: %v", err)
	// }
	// log.Printf("📨 Заявка подана/обновлена: %+v", req)

	// // 8.2 Список всех заявок пользователя
	// list, err := reqSvc.List(ctx, me.ID)
	// if err != nil {
	// 	log.Fatalf("❌ ListRequests error: %v", err)
	// }
	// log.Printf("📋 Всего заявок у пользователя %d: %d", me.ID, len(list))

	// // 8.3 Одобрить первую заявку (пример)
	// if len(list) > 0 {
	// 	first := list[0]
	// 	if err := reqSvc.Approve(ctx, first.ID); err != nil {
	// 		log.Fatalf("❌ ApproveRequest error: %v", err)
	// 	}
	// 	log.Printf("✅ Заявка %d одобрена", first.ID)
	// }

	// log.Println("🎉 Работа завершена")
}
