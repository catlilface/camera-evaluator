import { type ChangeEvent } from 'react'

import {
  Button,
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui'
import {
  EvaluationResultCard,
  FileSummaryCard,
  ImagePreviewCard,
  ImageUploadField,
  MethodSelector,
} from '@/features/evaluation/components'
import {
  type EvaluateResponse,
  type EvaluationMethodId,
} from '@/features/evaluation/types'

type EvaluationFormProps = {
  apiError: string
  apiResponse: EvaluateResponse | null
  error: string
  inputKey: number
  isSubmitting: boolean
  previewUrl: string | null
  selectedFile: File | null
  selectedMethod: EvaluationMethodId
  onFileChange: (event: ChangeEvent<HTMLInputElement>) => void
  onMethodChange: (method: EvaluationMethodId) => void
  onReset: () => void
  onSubmit: () => void
}

export function EvaluationForm({
  apiError,
  apiResponse,
  error,
  inputKey,
  isSubmitting,
  previewUrl,
  selectedFile,
  selectedMethod,
  onFileChange,
  onMethodChange,
  onReset,
  onSubmit,
}: EvaluationFormProps) {
  return (
    <Card className="border-white/70 bg-white/90">
      <CardHeader>
        <CardTitle>Форма оценки системы</CardTitle>
        <CardDescription>
          Выберите метод и загрузите изображение для дальнейшей обработки.
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-6">
        <MethodSelector
          selectedMethod={selectedMethod}
          onSelect={onMethodChange}
        />

        <ImageUploadField
          error={error}
          inputKey={inputKey}
          onChange={onFileChange}
        />

        <div className="grid gap-4 md:grid-cols-[0.95fr_1.05fr]">
          <FileSummaryCard file={selectedFile} />
          <ImagePreviewCard previewUrl={previewUrl} />
        </div>

        <EvaluationResultCard
          error={apiError}
          isSubmitting={isSubmitting}
          response={apiResponse}
        />

        <div className="flex flex-col gap-3 sm:flex-row sm:justify-end">
          <Button
            type="button"
            variant="outline"
            size="lg"
            onClick={onReset}
            disabled={isSubmitting}
          >
            Сбросить
          </Button>
          <Button
            type="button"
            size="lg"
            disabled={!selectedFile || isSubmitting}
            onClick={onSubmit}
          >
            Оценить изображение
          </Button>
        </div>
      </CardContent>
    </Card>
  )
}
