import { requestFormData } from '@/lib/http'
import { getMethodApiId } from '@/features/evaluation/constants'
import type {
  EvaluateResponse,
  EvaluationMethodId,
} from '@/features/evaluation/types'

export function evaluateImage(file: File, method: EvaluationMethodId) {
  const formData = new FormData()

  formData.append('method_id', String(getMethodApiId(method)))
  formData.append('image', file)

  return requestFormData<EvaluateResponse>('/api/v1/evaluate', formData)
}
