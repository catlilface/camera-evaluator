type ImagePreviewCardProps = {
  previewUrl: string | null
}

export function ImagePreviewCard({
  previewUrl,
}: ImagePreviewCardProps) {
  return (
    <div className="rounded-[28px] border border-slate-200 bg-slate-50 p-4">
      <div className="mb-3 text-sm font-medium text-slate-800">Превью</div>
      <div className="relative flex min-h-[220px] items-center justify-center overflow-hidden rounded-[24px] border border-slate-200 bg-white md:min-h-[260px]">
        {previewUrl ? (
          <img
            src={previewUrl}
            alt="Превью загруженного изображения"
            className="absolute inset-0 h-full w-full object-cover"
          />
        ) : (
          <div className="relative px-6 text-center text-sm leading-6 text-slate-500">
            Здесь появится превью после загрузки изображения.
          </div>
        )}
      </div>
    </div>
  )
}
