-- Запрос 1: Выборки всех уникальных eventType у которых более 1000 событий.
SELECT eventType
FROM events
GROUP BY eventType
HAVING COUNT(*) > 1000;

-- Запрос 2: Выборки событий которые произошли в первый день каждого месяца.
SELECT *
FROM events
WHERE toStartOfMonth(eventTime) = eventTime;

-- Запрос 3: Выборки пользователей которые совершили более 3 различных eventType.
SELECT userID
FROM events
GROUP BY userID
HAVING COUNT(DISTINCT eventType) > 3;
