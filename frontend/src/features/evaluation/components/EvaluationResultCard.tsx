import { ExternalLink, GitBranch, FileText, Brain } from "lucide-react";

import { Badge } from "@/components/ui/Badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/Card";
import type { EvaluationResult } from "@/features/evaluation/types";

type EvaluationResultCardProps = {
  result: EvaluationResult;
};

function getScoreColor(score: number): string {
  if (score >= 75) return "text-emerald-600";
  if (score >= 50) return "text-yellow-600";
  if (score >= 25) return "text-orange-600";
  return "text-red-600";
}

function getScoreBgColor(score: number): string {
  if (score >= 75) return "bg-emerald-500";
  if (score >= 50) return "bg-yellow-500";
  if (score >= 25) return "bg-orange-500";
  return "bg-red-500";
}

function getScoreLabel(score: number): string {
  if (score >= 85) return "Великолепное качество";
  if (score >= 70) return "Хорошее качество";
  if (score >= 50) return "Среднее качество";
  if (score >= 30) return "Ниже среднего";
  return "Низкое качество";
}

export function EvaluationResultCard({ result }: EvaluationResultCardProps) {
  const score = result.score ?? 0;
  const scoreColor = getScoreColor(score);
  const scoreBgColor = getScoreBgColor(score);
  const scoreLabel = getScoreLabel(score);

  return (
    <Card className="border-white/70 bg-white/90">
      <CardHeader>
        <CardTitle>Результат оценки</CardTitle>
        <CardDescription>
          Оценка качества изображения методом MUSIQ
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-6">
        <div className="relative flex min-h-[220px] items-center justify-center overflow-hidden rounded-[24px] border border-slate-200 bg-slate-50 md:min-h-[260px]">
          <img
            src={result.imageUrl}
            alt="Оцененное изображение"
            className="absolute inset-0 h-full w-full object-cover"
          />
        </div>

        <div className="space-y-3">
          <div className="flex items-baseline justify-between">
            <span className="text-sm font-medium text-slate-700">
              Оценка качества
            </span>
            <div className="flex items-baseline gap-2">
              <span className={`text-3xl font-bold ${scoreColor}`}>
                {score.toFixed(2)}
              </span>
              <span className="text-sm text-slate-500">/ 100</span>
            </div>
          </div>

          <div className="h-3 w-full overflow-hidden rounded-full bg-slate-100">
            <div
              className={`h-full transition-all duration-500 ${scoreBgColor}`}
              style={{ width: `${score}%` }}
            />
          </div>

          <div className="flex items-center">
            <Badge variant="secondary" className={scoreColor}>
              {scoreLabel}
            </Badge>
          </div>
        </div>

        <div className="rounded-[20px] border border-slate-200 bg-slate-50 p-4">
          <h4 className="mb-2 text-sm font-semibold text-slate-900">
            Метод MUSIQ
          </h4>
          <p className="mb-3 text-sm leading-6 text-slate-600">
            Multi-scale Image Quality Transformer — модель оценки качества
            изображений без эталона от Google Research. Использует multi-scale
            patch embedding и механизм Attention для анализа изображений на
            разных уровнях детализации.
          </p>
          <div className="flex flex-wrap gap-2">
            <a
              href="https://research.google/blog/musiq-assessing-image-aesthetic-and-technical-quality-with-multi-scale-transformers/"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center gap-1.5 rounded-full border border-slate-200 bg-white px-3 py-1.5 text-xs font-medium text-slate-700 transition-colors hover:bg-slate-50"
            >
              <Brain className="h-3.5 w-3.5" />
              Google Research
              <ExternalLink className="h-3 w-3" />
            </a>
            <a
              href="https://github.com/google-research/google-research/tree/master/musiq"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center gap-1.5 rounded-full border border-slate-200 bg-white px-3 py-1.5 text-xs font-medium text-slate-700 transition-colors hover:bg-slate-50"
            >
              <GitBranch className="h-3.5 w-3.5" />
              GitHub
              <ExternalLink className="h-3 w-3" />
            </a>
            <a
              href="https://arxiv.org/abs/2108.05997"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center gap-1.5 rounded-full border border-slate-200 bg-white px-3 py-1.5 text-xs font-medium text-slate-700 transition-colors hover:bg-slate-50"
            >
              <FileText className="h-3.5 w-3.5" />
              ArXiv
              <ExternalLink className="h-3 w-3" />
            </a>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
