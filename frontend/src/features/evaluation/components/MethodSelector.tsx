import { CheckCircle2 } from 'lucide-react'

import { methods } from '@/features/evaluation/constants'
import { type EvaluationMethodId } from '@/features/evaluation/types'

type MethodSelectorProps = {
  selectedMethod: EvaluationMethodId
  onSelect: (method: EvaluationMethodId) => void
}

export function MethodSelector({
  selectedMethod,
  onSelect,
}: MethodSelectorProps) {
  return (
    <div className="space-y-3">
      <label className="text-sm font-medium text-slate-800">Метод оценки</label>
      <div className="grid gap-3 sm:grid-cols-2">
        {methods.map((method) => {
          const isActive = method.id === selectedMethod

          return (
            <button
              key={method.id}
              type="button"
              onClick={() => onSelect(method.id)}
              className={`rounded-3xl border p-4 text-left transition-all duration-200 ${
                isActive
                  ? 'border-sky-400 bg-sky-50 shadow-sm'
                  : 'border-slate-200 bg-white hover:border-slate-300 hover:bg-slate-50'
              }`}
            >
              <div className="flex items-start justify-between gap-3">
                <div>
                  <div className="text-lg font-semibold text-slate-950">
                    {method.title}
                  </div>
                  <div
                    className="text-sm text-sky-700"
                    dangerouslySetInnerHTML={{ __html: method.subtitle }}
                  />
                </div>
                {isActive ? (
                  <CheckCircle2 className="h-5 w-5 text-sky-700" />
                ) : null}
              </div>
              <p className="mt-3 text-sm leading-6 text-slate-600">
                {method.description}
              </p>
            </button>
          )
        })}
      </div>
    </div>
  )
}
