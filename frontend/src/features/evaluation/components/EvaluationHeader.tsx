export function EvaluationHeader() {
  return (
    <header className="mb-6 flex flex-col gap-4 rounded-[28px] border border-white/70 bg-white/80 p-5 shadow-[0_20px_60px_-45px_rgba(15,23,42,0.45)] backdrop-blur md:mb-8 md:p-6">
      <div>
        <div>
          <h1 className="text-2xl font-semibold tracking-tight text-slate-950 sm:text-3xl">
            Оценка качества изображения
          </h1>
          <p className="mt-2 max-w-2xl text-sm leading-6 text-slate-600 sm:text-base">
            Простая форма для выбора метода оценки и загрузки изображения.
            Главный фокус пока на RR и NR без выбора мониторов.
          </p>
        </div>
      </div>
    </header>
  )
}
