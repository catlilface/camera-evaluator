import { methods } from './constants'

export type EvaluationMethod = (typeof methods)[number]
export type EvaluationMethodId = EvaluationMethod['id']

<<<<<<< Updated upstream
export type EvaluateResponse = {
  status: string
  id?: string
=======
export type WsResultMessage = {
  id: string
  score?: number
  status?: string
}

export type EvaluationResult = {
  imageUrl: string
  score: number | null
  status: string
>>>>>>> Stashed changes
}
