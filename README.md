# Практическая работа №3 Свидовский М.А ЭФМО-01-25  
## 1. Структура проекта:	
<img width="280" height="336" alt="image" src="https://github.com/user-attachments/assets/4da8e4ca-40fe-4b2c-9cae-bfb0df702c7b" />


## 2. GET /health — возвращает {"status":"ok"}
<img width="1497" height="510" alt="image" src="https://github.com/user-attachments/assets/3536aff7-ac56-404f-8eb8-c4e47d3dce7a" />


## 3. GET /tasks (вывод всех тасков) 
<img width="776" height="558" alt="image" src="https://github.com/user-attachments/assets/3e75f345-7fae-4376-a281-c3b0617a0fee" />

## 3.1 GET /tasks?q (поиск таска по названию)
<img width="776" height="558" alt="image" src="https://github.com/user-attachments/assets/01e8d585-027f-4195-829a-1b8b687d89de" />

## 4. POST /task (создниае нового таска)
<img width="885" height="140" alt="image" src="https://github.com/user-attachments/assets/e4b49a47-929a-4cda-97d1-326398a0b041" />

## 5. GET /tasks/{id} (поиск таска по id)
<img width="779" height="495" alt="image" src="https://github.com/user-attachments/assets/474c9b64-b0be-40ca-b935-41be76bc1103" />

## Доп плюшки
## DELETE /tasks/{id} 

### Проверяем наличие такого таска

<img width="779" height="495" alt="image" src="https://github.com/user-attachments/assets/b247cf4b-7c95-4f04-aeeb-324619628271" />

### Удаляем таск

<img width="779" height="495" alt="image" src="https://github.com/user-attachments/assets/d2e7f1a4-433e-4151-a42d-5e62d981d04f" />

### Проверяем после удаления есть такой или нет

<img width="779" height="495" alt="image" src="https://github.com/user-attachments/assets/01f0494b-e5fa-4cc2-b066-9f0e945c23c0" />

## Валидация: title
<img width="884" height="176" alt="image" src="https://github.com/user-attachments/assets/54b2d117-662f-4b3a-b898-2eb262abb4ba" />

## PATCH /tasks/{id}

### Создаем новый таск для проверки (не получается, валидация по титле :( )
<img width="1487" height="391" alt="image" src="https://github.com/user-attachments/assets/3ee1ddeb-0aed-4e48-a5ae-702570c66507" />

<img width="1488" height="629" alt="image" src="https://github.com/user-attachments/assets/88016a22-c4b0-4707-bc91-99aee51a4408" />


Меняем done на true, а title на другой 
<img width="1492" height="672" alt="image" src="https://github.com/user-attachments/assets/cb3a732d-25c9-4524-999d-3c607c959789" />

<img width="1503" height="706" alt="image" src="https://github.com/user-attachments/assets/2f284434-4d40-4be4-897f-aae1fe46ef05" />

## CORS 
<img width="1494" height="517" alt="image" src="https://github.com/user-attachments/assets/746612d4-cd4d-4b95-be22-1979ccb4785b" />

## Graceful shutdown
<img width="551" height="219" alt="image" src="https://github.com/user-attachments/assets/0e28e534-6ac9-4e54-b71f-f6113049d9ab" />

## Тестовые запросы
<img width="1407" height="700" alt="image" src="https://github.com/user-attachments/assets/f30995b1-b9e3-4f48-a70d-fffdc7ab72da" />











