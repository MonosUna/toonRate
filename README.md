# toonrate - Оцени, обсуждай и создавай топы любимой анимации!

## Проект по FullStack разработке студента Б05-227 Леванова Валерия

### [Figma](https://www.figma.com/design/K3BcHDciuBD6XnBQ0TKZch/fullstack?node-id=0-1&t=lN44366AwhqKEAX7-1)

## Функциональность
Пользователь может:
- Указывать список просмотренных, запланированных прозведений
- Оценивать все произведения представленные на сайте
- Создавать подборки и tier-листы произведений
- Делится ревью по произведению со всеми пользователями, как на форуме
- Открыть страницу случайного произведения
- Оценивать ревью других пользователей

На сайте есть:
- Общие оценки произведений от всех пользователей
- Топы по общей оценке, разделенные по разным жанрам.
- Форумы по произведениям с ревью

Сущности:
- Произведение
- Обсуждение
- Тир лист
- Ревью
- Подборка
- Пользователь 

### Итерация 1: done

### Итерация 2: done

### Итерация 3:
Сделана главная страница, можно создать-войти в аккаунт. В разделе Топы можно посмотреть прозведения, зайти на них. Если вы в аккаунте можно оценить произведение (изменится общий рейтинг)
p.s топы пока не допилены, так как они нужны для того чтобы можно было редактировать сущность. На главной странице еще необходимо окошко с обсуждением, но чтобы его сделать надо сделать обсуждения, так что оставляем до следующей итерации

### Итерация 4:
Весь фронтенд готов. Теперь пользователь не теряется при перезагрузке страницы, имеет статистику в профиле, может выставлять описание и аватарку, создавать обсуждения, подборки, обзоры, оставлять комментарии на них и оценивать комментарии. Добавлена возможность добавлять произведения в списки посмотренных, просмативаемых, и буду смотреть. Можно выбрать случайное произведение, обзор, обсуждение, подборку.

### Итерация 5:
В скрипте backend/main.go создаются базы данных.
Добавлена admin панель

### Итерация 7: 
Все взаимодействия с json-овской базой данных перенесены на go. Сайт полностью в рабочем состоянии.