import { methods } from './constants'

export type EvaluationMethod = (typeof methods)[number]
export type EvaluationMethodId = EvaluationMethod['id']

export type EvaluateResponse = {
  status: string
  id?: string
}
