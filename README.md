### BookService

• Основной сервис, который отвечает за управление данными о книгах (создание, чтение, обновление и удаление книг).

• Использует PostgreSQL для хранения данных о книгах и опционально Redis для кэширования.

• Логирует все действия на трех уровнях: debug, info, error.