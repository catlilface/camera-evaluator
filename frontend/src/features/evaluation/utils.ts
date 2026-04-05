import { MAX_FILE_SIZE } from './constants'

const BYTE_UNITS = ['Байт', 'Кб', 'Мб', 'Гб'] as const

export function formatBytes(bytes: number, decimals = 2) {
  if (bytes === 0) {
    return `0 ${BYTE_UNITS[0]}`
  }

  const base = 1024
  const index = Math.min(
    Math.floor(Math.log(bytes) / Math.log(base)),
    BYTE_UNITS.length - 1,
  )
  const size = bytes / base ** index

  return `${size.toFixed(decimals)} ${BYTE_UNITS[index]}`
}

export function formatFileSize(size: number) {
  return formatBytes(size)
}

export function validateImageFile(file: File) {
  if (!file.type.startsWith('image/')) {
    return 'Можно загружать только изображения.'
  }

  if (file.size > MAX_FILE_SIZE) {
    return `Файл превышает ограничение ${formatBytes(MAX_FILE_SIZE)}. Выберите изображение меньше.`
  }

  return ''
}

export function revokePreviewUrl(url: string | null) {
  if (url) {
    URL.revokeObjectURL(url)
  }
}
