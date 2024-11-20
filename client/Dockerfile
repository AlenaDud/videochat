FROM nginx:latest

# Объявляем переменные окружения для клиента
ENV SIGNAL_SERVER_WS=ws://0.0.0.0:3002/ws

# Копируем все файлы клиента в директорию Nginx
COPY ./ /usr/share/nginx/html/

# Добавляем обработку переменных окружения
RUN apt-get update && apt-get install -y gettext-base && apt-get clean

# Подготавливаем файл JavaScript для подстановки переменных окружения
RUN envsubst < /usr/share/nginx/html/app.js > /usr/share/nginx/html/app.js.generated && \
    mv /usr/share/nginx/html/app.js.generated /usr/share/nginx/html/app.js

# Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]

EXPOSE 80
