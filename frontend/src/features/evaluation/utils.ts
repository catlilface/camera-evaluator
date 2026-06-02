import { MAX_FILE_SIZE } from "./constants";

const BYTE_UNITS = ["Байт", "Кб", "Мб", "Гб"] as const;

export function formatBytes(bytes: number, decimals = 2) {
  if (bytes === 0) {
    return `0 ${BYTE_UNITS[0]}`;
  }

  const base = 1024;
  const index = Math.min(
    Math.floor(Math.log(bytes) / Math.log(base)),
    BYTE_UNITS.length - 1,
  );
  const size = bytes / base ** index;

  return `${size.toFixed(decimals)} ${BYTE_UNITS[index]}`;
}

export function formatFileSize(size: number) {
  return formatBytes(size);
}

export function validateImageFile(file: File) {
  if (!file.type.startsWith("image/")) {
    return "Можно загружать только изображения.";
  }

  if (file.size > MAX_FILE_SIZE) {
    return `Файл превышает ограничение ${formatBytes(MAX_FILE_SIZE)}. Выберите изображение меньше.`;
  }

  return "";
}

export function revokePreviewUrl(url: string | null) {
  if (url) {
    URL.revokeObjectURL(url);
  }
}

export function buildWsUrl(channelId: string): string {
  const apiBase = import.meta.env.VITE_API_BASE_URL;
  if (apiBase) {
    const wsBase = apiBase.replace(/^https/, "wss").replace(/^http/, "ws");
    return `${wsBase}/api/v1/ws?channel_id=${channelId}`;
  }
  const proto = window.location.protocol === "https:" ? "wss" : "ws";
  return `${proto}://${window.location.host}/api/v1/ws?channel_id=${channelId}`;
}

export function extractScore(
  data: Record<string, unknown> | number | null,
): number | null {
  if (typeof data === "number") return data;
  if (data !== null && typeof data === "object") {
    const s = data["score"] ?? data["Score"];
    if (typeof s === "number") return s;
  }
  return null;
}

export function buildAttentionImageUrl(attentionPath: string): string {
  const filename = attentionPath.split("/").pop();
  if (!filename) return "";
  const apiBase = import.meta.env.VITE_API_BASE_URL;
  if (apiBase) {
    return `${apiBase}/api/v1/photos/${filename}`;
  }
  return `/api/v1/photos/${filename}`;
}

export function uuidv4() {
  return "10000000-1000-4000-8000-100000000000".replace(/[018]/g, (c) =>
    (
      +c ^
      (crypto.getRandomValues(new Uint8Array(1))[0] & (15 >> (+c / 4)))
    ).toString(16),
  );
}
