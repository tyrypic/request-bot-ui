package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"ExBot/internal/adapter/bot"
	"ExBot/internal/adapter/db"
	"ExBot/internal/appinit"
	"ExBot/internal/texts"
	"ExBot/internal/usecase"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// 1. Загрузка переменных окружения из .env (если есть)
	appinit.LoadEnv()

	// 2. Получение строки подключения к базе данных из окружения
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal(texts.LogDBUrlMissing)
	}
	log.Printf(texts.LogDBConnect, dsn)

	// 3. Создание контекста с таймаутом для инициализации БД
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 4. Создание пула соединений с базой данных
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf(texts.LogDBConnectError, err)
	}
	log.Println(texts.LogDBConnected)
	defer func() {
		pool.Close()
		log.Println(texts.LogDBClosed)
	}()

	// 5. Применение миграций к базе данных
	appinit.RunMigrations(dsn)

	// 6. Регистрация администратора (seed)
	adminIDStr := os.Getenv("ADMIN_IDS")
	adminID, err := strconv.ParseInt(adminIDStr, 10, 64)
	if err != nil {
		log.Fatalf(texts.LogAdminIDNotInt, err)
	}
	log.Printf(texts.LogAdminID, adminID)

	userRepo := db.NewUserRepo(pool)
	userSvc := usecase.NewUserService(userRepo)
	if err := userSvc.SeedAdmin(ctx, adminID); err != nil {
		log.Fatalf(texts.LogAdminSeedError, err)
	}
	log.Println(texts.LogAdminSeedOk)

	// 7. Инициализация Telegram-бота
	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Fatalf(texts.LogBotInitError, err)
	}
	// botAPI.Debug = true // включить для отладки

	// 8. Инициализация репозиториев и сервисов сообщений
	messageRepo := bot.NewMessageRepo(botAPI, nil) // временно nil, замыкаем зависимость ниже
	messageSvc := usecase.NewMessageService(userRepo, messageRepo)
	messageRepo.MessageSvc = messageSvc // связываем сервис с адаптером

	// 9. Запуск слушателя Telegram-бота
	updates := botAPI.GetUpdatesChan(tgbotapi.NewUpdate(0))
	go messageRepo.Listen(context.Background(), updates)

	// 10. Блокировка main-горутины (бот работает, пока не завершится процесс)
	select {}
}
