# Оценщик качества камеры 

Работа выполняется в рамках лабораторных работ по дисциплине Техническое зрение

## Архитектура

Была выбрана микросервисная архитектура с брокером. Такая архитектура позволяет максимально эффективно обрабатывать множество запросов и оценивать качество, особенно когда оно занимает значитаельное время. Ниже представлена диаграмма последовательности по взаимодействию внутри системы.

sequenceDiagram
  Клиент->>Сервер: Сырой запрос
  Сервер-->>RabbitMQ: Запрос с параметрами
  RabbitMQ-->>Оценщик: Очередь на обработку
  Оценщик-->>RabbitMQ: Очередь ответов
  RabbitMQ-->>Сервер: Оценка
  Сервер->>Клиент: Результат оценки

## Стек

Backend: Go 

Frontend: React, TypeScript, Nginx

Evaluator: Python (FastAPI / MUSIQ)

Брокер: RabbitMQ

База данных: PostgreSQL

## Литература

1. Chen C., Mo J. IQA-PyTorch: PyTorch Toolbox for Image Quality Assessment. URL: https://github.com/chaofengc/IQA-PyTorch

2. scikit-image. skimage.metrics: Image Quality Metrics. URL: https://scikit-image.org/docs/stable/api/skimage.metrics.html

3. Calibrite. ColorChecker Classic Mini. URL: https://calibrite.com/us/product/colorchecker-classic-mini/

4. Image Quality Labs. ISO-12233 Enhanced Digital Camera Resolution Chart. URL: https://www.imagequalitylabs.com/product/iso-12233-enhanced-digital-camera-resolution-chart

5. Image Quality Labs. Checkerboard Distortion Chart for Microsoft Lync Certification. URL: https://www.imagequalitylabs.com/product/checkerboard-distortion-for-microsoft-c-lync-tm-certification

6. DPReview. Digital Camera Reviews. URL: https://www.dpreview.com

7. Imaging Resource. Camera Reviews and News. URL: https://www.imaging-resource.com

## Запуск
### Первый запуск на Windows 10
1. Проверьте, что у Вас установлен Docker Desktop и WSL

2. В настройках Docker Desktop (Settings - Resources - WSL Integration) проверьте активные галочку и переключатель напротив Ubuntu.
<img width="687" height="403" alt="image" src="https://github.com/user-attachments/assets/f219d155-2fb3-47ca-a1f9-9711cde01c99" />

3. Откройте PowerShell и последовательно выполните
```powershell
wsl
cd ~
git clone https://github.com/catlilface/camera-evaluator.git
cd camera-evaluator
make init
make build
make up
```
4. Откройте браузер на http://localhost:3000

### Повторный запуск
1. Откройте PowerShell и выполните:
```powershell
wsl
cd ~/camera-evaluator
make up
```
2. Откройте браузер на http://localhost:3000

### Остановка
```powershell
wsl
cd ~/camera-evaluator
make down
```
## Лицензия

MIT
