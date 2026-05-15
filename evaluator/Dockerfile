FROM ghcr.io/astral-sh/uv:python3.13-alpine

WORKDIR /app

COPY ./pyproject.toml ./uv.lock ./.env /app/

RUN uv sync --no-dev

COPY ./src/ /app/src/

CMD ["uv", "run", "--no-sync", "faststream", "run", "src.broker.app:app"]
