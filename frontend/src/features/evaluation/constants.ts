const DEFAULT_MAX_FILE_SIZE = 50 * 1024 * 1024

function getMaxFileSize() {
  const value = Number(import.meta.env.VITE_MAX_FILE_SIZE)

  if (Number.isFinite(value) && value > 0) {
    return value
  }

  return DEFAULT_MAX_FILE_SIZE
}

export const MAX_FILE_SIZE = getMaxFileSize()

export const methods = [
  {
    id: 'rr',
    title: 'RR',
    subtitle: 'Reference&#8209;based',
    description:
      'Метод с эталонным изображением для последующего сравнения.',
  },
  {
    id: 'nr',
    title: 'NR',
    subtitle: 'No&#8209;reference',
    description:
      'Метод без эталона, когда оценка строится только по загруженной картинке.',
  },
] as const

export const notes = [
  'Сейчас доступны только методы RR и NR.',
  'Выбор мониторов в этой версии не используется.',
  `Максимальный размер изображения: ${MAX_FILE_SIZE / (1024 * 1024)} МБ.`,
] as const
