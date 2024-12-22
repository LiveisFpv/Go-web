# syntax=docker/dockerfile:1

ARG PYTHON_VERSION=3.11.8
FROM python:${PYTHON_VERSION}-slim as base

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

WORKDIR /app

COPY backend-pdf ./

RUN apt-get update && apt-get install -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Устанавливаем Python-зависимости
RUN --mount=type=cache,target=/root/.cache/pip \
    --mount=type=bind,source=server_conf/requirements.txt,target=requirements.txt \
    python -m pip install -r requirements.txt


# Порт, используемый приложением
EXPOSE 9000

# Запускаем приложение
CMD ["gunicorn", "-w", "2", "-b", "0.0.0.0:9000", "app:app"]
