#!/bin/bash

# Главный файл schema.yaml
MAIN_SCHEMA="schema.yaml"

# Проверка, существует ли главный файл schema.yaml
if [[ ! -f "$MAIN_SCHEMA" ]]; then
  echo "Файл $MAIN_SCHEMA не найден!"
  exit 1
fi

# Создание временного файла для хранения обновленного paths
TEMP_FILE=$(mktemp)

# Добавление первой части (до paths) из главного schema.yaml
awk '/paths:/ {exit} {print}' "$MAIN_SCHEMA" > "$TEMP_FILE"
echo "paths:" >> "$TEMP_FILE"

# Функция для экранирования пути в JSON Pointer
escape_path() {
  echo "$1" | sed 's/~/'~0'/g; s/\//~1/g'
}

# Поиск и добавление путей из подкаталогов
for dir in */; do
  if [[ -f "$dir/schema.yaml" ]]; then
    while IFS= read -r path_line; do
      path=$(echo "$path_line" | cut -d':' -f1)
      escaped_path=$(escape_path "$path")
      echo "  $path:" >> "$TEMP_FILE"
      echo "    \$ref: './${dir}schema.yaml#/paths/$escaped_path'" >> "$TEMP_FILE"
    done < <(awk '/paths:/,/definitions:/{if ($0 ~ /^  \//) print substr($1,1,length($1)-1);}' "$dir/schema.yaml")
  fi
done

# Замена старого schema.yaml на обновленный
mv "$TEMP_FILE" "$MAIN_SCHEMA"
echo "Объединение схем завершено!"

# Слияние схем в одну общую
java -jar swagger-codegen-cli.jar generate -l openapi-yaml -i $MAIN_SCHEMA -o . -DoutputFile=./openApi.yaml
echo "Финальный фаил готов!"

# Убираем лишнии файлы
rm README.md .swagger-codegen-ignore
rm -rf .swagger-codegen
