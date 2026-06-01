import { AlertCircle, CheckCircle2, Clock3 } from "lucide-react";

import type { EvaluateResponse } from "@/features/evaluation/types";

type EvaluationResultCardProps = {
  error: string;
  isSubmitting: boolean;
  response: EvaluateResponse | null;
};

export function EvaluationResultCard({
  error,
  isSubmitting,
  response,
}: EvaluationResultCardProps) {
  if (!isSubmitting && !response && !error) {
    return null;
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
    );
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
    );
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
          <dd className="font-medium">{response?.status || "OK"}</dd>
        </div>
        <div>
          <dt className="text-emerald-700">ID</dt>
          <dd className="break-all font-medium">
            {response?.id ?? "Не вернулся"}
          </dd>
        </div>
      </dl>
      <pre className="mt-4 overflow-auto rounded-2xl bg-white/80 p-3 text-xs text-slate-800">
        {JSON.stringify(response, null, 2)}
      </pre>
    </div>
  );
}
