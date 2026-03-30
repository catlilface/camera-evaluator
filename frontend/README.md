## 🚀 Стек технологий

| Категория | Технология | Версия |
|-----------|------------|--------|
| **Фреймворк** | React | 19.x |
| **Язык** | TypeScript | 5.9.x |
| **Сборщик** | Vite | 8.x |
| **Стилизация** | Tailwind CSS | 3.4.x |
| **UI-библиотека** | shadcn/ui (Radix UI) | latest |
| **Менеджер состояния** | Zustand | 5.x |
| **Линтинг** | ESLint + Prettier | 9.x / 3.x |

---

## 📁 Структура проекта

```
frontend/
├── public/                 # Статические файлы (favicon, robots.txt)
├── src/
│   ├── components/
│   │   └── ui/            # Компоненты shadcn/ui (button, input, card...)
│   ├── store/
│   │   └── useStore.ts    # Zustand стор (глобальное состояние)
│   ├── lib/
│   │   └── utils.ts       # Утилиты (cn, clsx)
│   ├── App.tsx            # Главный компонент приложения
│   ├── main.tsx           # Точка входа (ReactDOM)
│   └── index.css          # Глобальные стили + директивы Tailwind
├── index.html             # HTML-шаблон (корень проекта)
├── package.json           # Зависимости и скрипты
├── vite.config.ts         # Конфигурация Vite (алиасы, плагины)
├── tailwind.config.js     # Конфигурация Tailwind CSS
├── tsconfig.json          # Конфигурация TypeScript
├── eslint.config.js       # Конфигурация ESLint (Flat Config)
├── .prettierrc            # Настройки форматирования Prettier
├── .gitignore             # Игнорируемые файлы для Git
└── README.md              # Документация проекта
```

---

## ⚙️ Переменные окружения

На данный момент проект **не требует** настройки переменных окружения для запуска.

При необходимости добавления API-ключей или настроек, создайте файл `.env` в корне проекта:

---

## 🛠️ Установка и запуск

### Требования

- **Node.js** версии 18.x или выше
- **npm** версии 9.x или выше

Проверить версии:
```bash
node -v
npm -v
```

### 1. Клонирование репозитория

```bash
git clone https://github.com/catlilface/camera-evaluator.git
cd camera-evaluator/frontend
```

### 2. Установка зависимостей

```bash
npm install
```

### 3. Запуск в режиме разработки

```bash
npm run dev
```

После запуска откройте в браузере: **http://localhost:5173**


### 4. Сборка для продакшена

```bash
npm run build
```

Собранные файлы появятся в папке `dist/`.


---

## 📜 Доступные скрипты

| Команда | Описание |
|---------|----------|
| `npm run dev` | Запуск сервера разработки (Vite) |
| `npm run build` | Сборка проекта для продакшена |
| `npm run preview` | Просмотр собранной версии локально |
| `npm run lint` | Проверка кода через ESLint |
| `npm run lint:fix` | Автоматическое исправление ошибок ESLint |
| `npm run format` | Форматирование кода через Prettier |
| `npm run format:check` | Проверка форматирования без изменений |

---

### Автоматическое форматирование при сохранении

Добавьте в `.vscode/settings.json`:

```json
{
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": "explicit"
  },
  "css.lint.unknownAtRules": "ignore"
}