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
<<<<<<< Updated upstream
  isSubmitting: boolean
=======
  isLoading: boolean
>>>>>>> Stashed changes
  previewUrl: string | null
  selectedFile: File | null
  selectedMethod: EvaluationMethodId
  submissionError: string
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
<<<<<<< Updated upstream
  isSubmitting,
=======
  isLoading,
>>>>>>> Stashed changes
  previewUrl,
  selectedFile,
  selectedMethod,
  submissionError,
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

<<<<<<< Updated upstream
        <EvaluationResultCard
          error={apiError}
          isSubmitting={isSubmitting}
          response={apiResponse}
        />
=======
        {submissionError && (
          <p className="text-sm text-red-600">{submissionError}</p>
        )}
>>>>>>> Stashed changes

        <div className="flex flex-col gap-3 sm:flex-row sm:justify-end">
          <Button
            type="button"
            variant="outline"
            size="lg"
<<<<<<< Updated upstream
            onClick={onReset}
            disabled={isSubmitting}
=======
            disabled={isLoading}
            onClick={onReset}
>>>>>>> Stashed changes
          >
            Сбросить
          </Button>
          <Button
            type="button"
            size="lg"
<<<<<<< Updated upstream
            disabled={!selectedFile || isSubmitting}
            onClick={onSubmit}
          >
            Оценить изображение
=======
            disabled={!selectedFile || isLoading}
            onClick={onSubmit}
          >
            {isLoading ? 'Обработка...' : 'Оценить изображение'}
>>>>>>> Stashed changes
          </Button>
        </div>
      </CardContent>
    </Card>
  )
}
