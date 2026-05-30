<<<<<<< Updated upstream
import { AlertCircle, CheckCircle2, Clock3 } from 'lucide-react'

import type { EvaluateResponse } from '@/features/evaluation/types'

type EvaluationResultCardProps = {
  error: string
  isSubmitting: boolean
  response: EvaluateResponse | null
}

export function EvaluationResultCard({
  error,
  isSubmitting,
  response,
}: EvaluationResultCardProps) {
  if (!isSubmitting && !response && !error) {
    return null
  }

  if (isSubmitting) {
    return (
      <div className="rounded-[24px] border border-sky-200 bg-sky-50 p-4 text-sm text-sky-900">
        <div className="flex items-center gap-2 font-medium">
          <Clock3 className="h-4 w-4" />
          Отправляем изображение
        </div>
        <p className="mt-2 leading-6 text-sky-800">
          Файл загружается в API оценки. Дождитесь ответа сервера.
        </p>
      </div>
    )
  }

  if (error) {
    return (
      <div className="rounded-[24px] border border-red-200 bg-red-50 p-4 text-sm text-red-800">
        <div className="flex items-center gap-2 font-medium">
          <AlertCircle className="h-4 w-4" />
          Ошибка API
        </div>
        <p className="mt-2 break-words leading-6">{error}</p>
      </div>
    )
  }

  return (
    <div className="rounded-[24px] border border-emerald-200 bg-emerald-50 p-4 text-sm text-emerald-900">
      <div className="flex items-center gap-2 font-medium">
        <CheckCircle2 className="h-4 w-4" />
        Ответ API получен
      </div>
      <dl className="mt-3 grid gap-2 sm:grid-cols-2">
        <div>
          <dt className="text-emerald-700">Status</dt>
          <dd className="font-medium">{response?.status || 'OK'}</dd>
        </div>
        <div>
          <dt className="text-emerald-700">ID</dt>
          <dd className="break-all font-medium">
            {response?.id ?? 'Не вернулся'}
          </dd>
        </div>
      </dl>
      <pre className="mt-4 overflow-auto rounded-2xl bg-white/80 p-3 text-xs text-slate-800">
        {JSON.stringify(response, null, 2)}
      </pre>
    </div>
=======
import {
  Button,
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui'
import { type EvaluationResult } from '@/features/evaluation/types'

type EvaluationResultCardProps = EvaluationResult & {
  onReset: () => void
}

export function EvaluationResultCard({
  imageUrl,
  score,
  status,
  onReset,
}: EvaluationResultCardProps) {
  const displayScore = score !== null ? Math.round(score * 10) / 10 : null

  return (
    <Card className="border-white/70 bg-white/90">
      <CardHeader>
        <CardTitle>Результат оценки</CardTitle>
        <CardDescription>Статус: {status || 'завершено'}</CardDescription>
      </CardHeader>
      <CardContent className="space-y-6">
        <div className="overflow-hidden rounded-xl border border-slate-200 bg-slate-50">
          <img
            src={imageUrl}
            alt="Загруженное изображение"
            className="mx-auto max-h-64 w-full object-contain"
          />
        </div>

        <div className="flex items-center justify-center rounded-xl border border-blue-200 bg-blue-50 p-6">
          {displayScore !== null ? (
            <div className="text-center">
              <div className="text-5xl font-bold tabular-nums text-blue-700">
                {displayScore}
              </div>
              <div className="mt-1 text-sm text-blue-500">из 100</div>
            </div>
          ) : (
            <p className="text-sm text-slate-500">Оценка недоступна</p>
          )}
        </div>

        <div className="flex justify-end">
          <Button type="button" variant="outline" size="lg" onClick={onReset}>
            Новая оценка
          </Button>
        </div>
      </CardContent>
    </Card>
>>>>>>> Stashed changes
  )
}
