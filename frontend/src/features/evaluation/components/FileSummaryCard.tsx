import { FileImage, Info } from 'lucide-react'

import { formatFileSize } from '@/features/evaluation/utils'

type FileSummaryCardProps = {
  file: File | null
}

export function FileSummaryCard({ file }: FileSummaryCardProps) {
  return (
    <div className="flex h-full flex-col rounded-[28px] border border-slate-200 bg-slate-50 p-4">
      <div className="mb-3 flex items-center gap-2 text-sm font-medium text-slate-800">
        <FileImage className="h-4 w-4 text-sky-700" />
        Информация о файле
      </div>

      {file ? (
        <div className="space-y-2 text-sm text-slate-600">
          <p className="break-all">
            <span className="font-medium text-slate-800">Имя:</span> {file.name}
          </p>
          <p>
            <span className="font-medium text-slate-800">Размер:</span>{' '}
            {formatFileSize(file.size)}
          </p>
          <p>
            <span className="font-medium text-slate-800">Тип:</span>{' '}
            {file.type || 'Не определён'}
          </p>
        </div>
      ) : (
        <p className="text-sm leading-6 text-slate-500">
          Файл ещё не выбран. После загрузки здесь появится краткая информация.
        </p>
      )}

      <div className="mt-auto rounded-2xl border border-amber-200 bg-amber-50 p-3 text-sm text-amber-800">
        <div className="flex gap-2">
          <Info className="mt-0.5 h-4 w-4 shrink-0" />
          <span>
            Для метода RR позже можно будет добавить отдельную загрузку
            эталонного изображения.
          </span>
        </div>
      </div>
    </div>
  )
}
