import { useEffect, useMemo, useState, type ChangeEvent } from 'react'

import {
  EvaluationForm,
  EvaluationHeader,
  EvaluationSidebar,
} from '@/features/evaluation/components'
import { methods } from '@/features/evaluation/constants'
import { type EvaluationMethodId } from '@/features/evaluation/types'
import {
  revokePreviewUrl,
  validateImageFile,
} from '@/features/evaluation/utils'

export function EvaluationPage() {
  const [selectedMethod, setSelectedMethod] = useState<EvaluationMethodId>('rr')
  const [selectedFile, setSelectedFile] = useState<File | null>(null)
  const [previewUrl, setPreviewUrl] = useState<string | null>(null)
  const [error, setError] = useState('')
  const [inputKey, setInputKey] = useState(0)

  const activeMethod = useMemo(
    () => methods.find((method) => method.id === selectedMethod) ?? methods[0],
    [selectedMethod],
  )

  useEffect(() => {
    return () => {
      revokePreviewUrl(previewUrl)
    }
  }, [previewUrl])

  function handleFileChange(event: ChangeEvent<HTMLInputElement>) {
    const file = event.target.files?.[0]

    if (!file) {
      revokePreviewUrl(previewUrl)
      setSelectedFile(null)
      setPreviewUrl(null)
      setError('')
      return
    }

    const validationError = validateImageFile(file)

    if (validationError) {
      revokePreviewUrl(previewUrl)
      setSelectedFile(null)
      setPreviewUrl(null)
      setError(validationError)
      event.target.value = ''
      return
    }

    revokePreviewUrl(previewUrl)
    setSelectedFile(file)
    setPreviewUrl(URL.createObjectURL(file))
    setError('')
  }

  function handleReset() {
    revokePreviewUrl(previewUrl)
    setSelectedMethod('rr')
    setSelectedFile(null)
    setPreviewUrl(null)
    setError('')
    setInputKey((current) => current + 1)
  }

  return (
    <main className="min-h-screen bg-[linear-gradient(180deg,_#f8fbff_0%,_#eff6ff_100%)] text-slate-900">
      <section className="mx-auto max-w-6xl px-4 py-6 sm:px-6 md:px-8 md:py-8 lg:py-10">
        <EvaluationHeader />

        <div className="grid gap-6 lg:grid-cols-[0.9fr_1.1fr]">
          <EvaluationSidebar activeMethod={activeMethod} />
          <EvaluationForm
            error={error}
            inputKey={inputKey}
            previewUrl={previewUrl}
            selectedFile={selectedFile}
            selectedMethod={selectedMethod}
            onFileChange={handleFileChange}
            onMethodChange={setSelectedMethod}
            onReset={handleReset}
          />
        </div>
      </section>
    </main>
  )
}
