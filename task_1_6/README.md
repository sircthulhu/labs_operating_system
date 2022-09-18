# Защита ОС (лаб 1 задание 6)
## Задача
Дан текст длиной n символов и ключ длиной k символов. 
Осуществите блочное шифрование текста по следующему алгоритму:
- Разделить текст на блоки длиной k символов.
Если n не кратно k, то допустимо, чтобы длина последнего блока была меньше k.
- Внутри каждого блока выполнить перестановку символов так,
чтобы первый символ занял место последнего, второй – предпоследнего и т.д. 
Последний символ должен оказаться на месте первого символа блока.
- Применить ключ к каждому блоку. Шифрованный i-й символ блока должен быть получен, 
как результат исключающего или между i-м исходным символом блока и i-м символом ключа.