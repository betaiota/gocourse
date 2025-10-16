#!/bin/bash
  
gameInit() {
    export attempts=3
    read -p "Введите загадываемое слово: " -s wonderWord
    revealed=()
    for ((i=0; i<${#wonderWord}; i++)); do
        revealed+=("*")
    done
    status="Начните угадывать по буквам, или назовите слово целиком."
    clear
}

gameMainLoop() {
    while true; do
        clear
        echo $status
        echo "----------------------------------------------------------"
        if [[ $attempts -eq 0 ]]; then
            echo "У вас не осталось попыток. Загаданное слово - $wonderWord. Конец игры."
            break
        elif [[ $attempts -eq 1  ]]; then
            echo "У вас осталась $attempts попытка."
        else
            echo "У вас осталось $attempts попытки."
        fi
        echo "Слово - ${revealed[@]}"
        echo "Можно:"
        echo "1. Угадать одну букву"
        echo "2. Ввести слово целиком"
        read -p "Выберите вариант: " state
        case "$state" in
            "1")
                read -p "Введите букву: " letter
                if [[ ${#letter} -ge 2 ]]; then
                    status="Вы ввели больше одной буквы..."
                    continue
                fi
                if [[ "$wonderWord" =~ "$letter" ]]; then
                    status="В слове есть буква $letter!"
                    for ((i=0; i<${#wonderWord}; i++)); do
                        if [ "${wonderWord:$i:1}" = "$letter" ]; then
                            revealed[$i]=$letter
                        fi
                    done
                    if [[ "${revealed[@]}" != *"*"* ]]; then
                            echo "Вы угадали слово - $wonderWord! Победа!"
                            break;
                    fi
                else
                    status="Буквы $letter нет в слове, минус попытка!"
                    attempts=$(($attempts - 1))
                fi
                ;;
            "2")
                echo "Угадываем слово."
                read -p "Введите слово целиком: " word
                if [[ "$wonderWord" = "$word" ]]; then
                    echo "Вы угадали слово - $wonderWord! Победа!"
                    break
                else
                    echo "Неправильный ответ! Загаданное слово - $wonderWord. Конец игры."
                    break
                fi
                ;;
            *)
                status="Выбран некорректный вариант, повторите..."
        esac
    done
}

gameInit
gameMainLoop