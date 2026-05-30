const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? ''

export class ApiError extends Error {
  status: number
  payload: unknown

  constructor(message: string, status: number, payload: unknown) {
    super(message)
    this.name = 'ApiError'
    this.status = status
    this.payload = payload
  }
}

function resolveUrl(path: string) {
  if (!API_BASE_URL) {
    return path
  }

  return `${API_BASE_URL.replace(/\/$/, '')}${path}`
}

function getErrorMessage(payload: unknown) {
  if (
    typeof payload === 'object' &&
    payload !== null &&
    'error' in payload &&
    typeof payload.error === 'string'
  ) {
    return payload.error
  }

  if (
    typeof payload === 'object' &&
    payload !== null &&
    'message' in payload &&
    typeof payload.message === 'string'
  ) {
    return payload.message
  }

  if (typeof payload === 'string' && payload) {
    return payload
  }

  return 'Request failed'
}

export async function requestFormData<T>(
  path: string,
  formData: FormData,
  options: Omit<RequestInit, 'body' | 'method'> = {},
): Promise<T> {
  const response = await fetch(resolveUrl(path), {
    ...options,
    method: 'POST',
    body: formData,
  })

  const contentType = response.headers.get('content-type') ?? ''
  const payload = contentType.includes('application/json')
    ? await response.json()
    : await response.text()

  if (!response.ok) {
    throw new ApiError(getErrorMessage(payload), response.status, payload)
  }

  return payload as T
}
