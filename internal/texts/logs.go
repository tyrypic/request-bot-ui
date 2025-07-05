package texts

const (
	// ENV
	LogEnvLoaded   = "✅ Переменные загружены из .env"
	LogEnvNotFound = "⚠️  .env файл не найден, полагаемся на реальные ENV"

	// DATABASE
	LogDBConnect      = "🔗 Подключаемся к БД: %s"
	LogDBConnected    = "✅ Подключение к БД установлено"
	LogDBClosed       = "👋 Соединение с БД закрыто"
	LogDBConnectError = "❌ Не удалось подключиться к БД: %v"
	LogDBUrlMissing   = "✋ Ошибка: DATABASE_URL не задана"

	// MIGRATIONS
	LogMigrationsStart = "Запуск миграций"
	LogMigrationsNew   = "миграции: не смогли создать экземпляр: %v"
	LogMigrationsUp    = "миграции: UP упали: %v"
	LogMigrationsDone  = "✅ Миграции применены"

	// ADMIN
	LogAdminIDNotInt  = "❌ ADMIN_ID не число: %v"
	LogAdminID        = "👑 ID Администратора: %v"
	LogAdminSeedError = "❌ Не удалось зарегистрировать администратора: %v"
	LogAdminSeedOk    = "✅ Администратор зарегистрирован"

	// BOT
	LogBotTokenMissing = "✋ Ошибка: TELEGRAM_APITOKEN не задан"
	LogBotInitError    = "❌ Не удалось инициализировать Telegram-бота: %v"

	// USER
	LogUserRegisterError  = "❌ Register error: %v"
	LogUserRegistered     = "👤 Пользователь зарегистрирован: %+v"
	LogProfileUpdateError = "❌ UpdateProfile error: %v"
	LogProfileUpdated     = "✏️ Профиль обновлён: %+v"

	// REQUESTS
	LogRequestSubmitError  = "❌ SubmitRequest error: %v"
	LogRequestSubmitted    = "📨 Заявка подана/обновлена: %+v"
	LogRequestListError    = "❌ ListRequests error: %v"
	LogRequestList         = "📋 Всего заявок у пользователя %d: %d"
	LogRequestApproveError = "❌ ApproveRequest error: %v"
	LogRequestApproved     = "✅ Заявка %d одобрена"

	// VERIFIED
	MsgUserCheckError  = "Ошибка при проверке верификации."
	MsgUserNotVerified = "Для того, чтобы воспользоваться функционалом бота, пройдите верификацию."
	MsgUserVerified    = "Бот готов к использованию"

	// OTHER
	LogDone = "🎉 Работа завершена"
)
