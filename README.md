# Candy-server
Конфетный бизнес оказался под угрозой из-за взлома сервера и ухода разработчика. Нужно было восстановить сервер и добавить защиту от атаки "Man-in-the-middle". 

Для этого:

 * Был написан сервер на Go, который обрабатывает заказы на покупку конфет по определенным правилам (проверка суммы денег, выбор типа конфет и их количества).
   Код сгенерирован из swagger спецификации.
 * Реализована аутентификация с помощью сертификатов для сервера и тестового клиента, использующего самоподписанный сертификат и локальный центр сертификации.
 * Добавлена функцию, которая при покупке конфет выводит ASCII-изображение коровы с помощью функции, написанной на C, без изменения ее кода.
