FROM ghcr.io/astral-sh/uv:python3.13-bookworm

WORKDIR /app

COPY ./pyproject.toml ./uv.lock /app/

RUN apt-get update && apt-get install -y \
    libgl1-mesa-glx \
    libglib2.0-0 \
    && rm -rf /var/lib/apt/lists/*

RUN uv sync --no-dev

COPY ./src/ /app/src/

CMD ["uv", "run", "--no-sync", "faststream", "run", "src.broker.app:app"]
