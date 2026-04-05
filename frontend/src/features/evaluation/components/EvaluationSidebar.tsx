import { CheckCircle2 } from 'lucide-react'

import { Badge, Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui'
import { notes } from '@/features/evaluation/constants'
import { type EvaluationMethod } from '@/features/evaluation/types'

type EvaluationSidebarProps = {
  activeMethod: EvaluationMethod
}

export function EvaluationSidebar({
  activeMethod,
}: EvaluationSidebarProps) {
  return (
    <Card className="border-white/70 bg-white/85">
      <CardHeader>
        <CardTitle>Что доступно сейчас</CardTitle>
        <CardDescription>
          Текущая версия страницы содержит только необходимое для демонстрации
          формы.
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        {notes.map((note) => (
          <div
            key={note}
            className="flex gap-3 rounded-2xl border border-slate-200 bg-slate-50 p-4"
          >
            <CheckCircle2 className="mt-0.5 h-5 w-5 shrink-0 text-emerald-600" />
            <p className="text-sm leading-6 text-slate-600">{note}</p>
          </div>
        ))}

        <div className="rounded-3xl bg-slate-950 p-5 text-white">
          <p className="text-sm uppercase tracking-[0.22em] text-sky-200">
            Активный метод
          </p>
          <div className="mt-3 flex items-end justify-between gap-4">
            <div>
              <div className="text-3xl font-semibold">{activeMethod.title}</div>
              <p className="mt-2 text-sm leading-6 text-slate-300">
                {activeMethod.description}
              </p>
            </div>
            <Badge
              className="border-white/10 bg-white/10 text-white"
              dangerouslySetInnerHTML={{ __html: activeMethod.subtitle }}
            />
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
