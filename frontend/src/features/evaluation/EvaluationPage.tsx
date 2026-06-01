import { useEffect, useMemo, useState, type ChangeEvent } from "react";

import {
  EvaluationForm,
  EvaluationHeader,
  EvaluationResultCard,
  EvaluationSidebar,
} from "@/features/evaluation/components";
// import { evaluateImage } from "@/features/evaluation/api";
import { methods } from "@/features/evaluation/constants";
import {
  type EvaluationMethodId,
  type EvaluationResult,
  type WsResultMessage,
  type EvaluateResponse,
} from "@/features/evaluation/types";
import {
  buildWsUrl,
  revokePreviewUrl,
  validateImageFile,
} from "@/features/evaluation/utils";

type Phase = "idle" | "loading" | "result" | "error";

export function EvaluationPage() {
  const [selectedMethod, setSelectedMethod] =
    useState<EvaluationMethodId>("rr");
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [previewUrl, setPreviewUrl] = useState<string | null>(null);
  const [error, setError] = useState("");
  const [_apiError, setApiError] = useState("");
  const [_apiResponse, setApiResponse] = useState<EvaluateResponse | null>(
    null,
  );
  const [_isSubmitting, setIsSubmitting] = useState(false);
  const [inputKey, setInputKey] = useState(0);

  const [phase, setPhase] = useState<Phase>("idle");
  const [result, setResult] = useState<EvaluationResult | null>(null);
  const [submissionError, setSubmissionError] = useState("");

  const activeMethod = useMemo(
    () => methods.find((method) => method.id === selectedMethod) ?? methods[0],
    [selectedMethod],
  );

  useEffect(() => {
    return () => {
      revokePreviewUrl(previewUrl);
    };
  }, [previewUrl]);

  function handleFileChange(event: ChangeEvent<HTMLInputElement>) {
    const file = event.target.files?.[0];

    if (!file) {
      revokePreviewUrl(previewUrl);
      setSelectedFile(null);
      setPreviewUrl(null);
      setError("");
      setApiError("");
      setApiResponse(null);
      return;
    }

    const validationError = validateImageFile(file);

    if (validationError) {
      revokePreviewUrl(previewUrl);
      setSelectedFile(null);
      setPreviewUrl(null);
      setError(validationError);
      setApiError("");
      setApiResponse(null);
      event.target.value = "";
      return;
    }

    revokePreviewUrl(previewUrl);
    setSelectedFile(file);
    setPreviewUrl(URL.createObjectURL(file));
    setError("");
    setApiError("");
    setApiResponse(null);
  }

  // async function handleSubmit() {
  //   if (!selectedFile || isSubmitting) {
  //     return;
  //   }

  //   setIsSubmitting(true);
  //   setApiError("");
  //   setApiResponse(null);

  //   try {
  //     const response = await evaluateImage(selectedFile, selectedMethod);
  //     setApiResponse(response);
  //   } catch (error) {
  //     setApiError(
  //       error instanceof Error
  //         ? error.message
  //         : "Не удалось отправить изображение",
  //     );
  //   } finally {
  //     setIsSubmitting(false);
  //   }
  // }

  function handleReset() {
    revokePreviewUrl(previewUrl);
    setSelectedMethod("rr");
    setSelectedFile(null);
    setPreviewUrl(null);
    setError("");
    setApiError("");
    setApiResponse(null);
    setIsSubmitting(false);
    setInputKey((current) => current + 1);
    setPhase("idle");
    setResult(null);
    setSubmissionError("");
  }

  function handleEvaluate() {
    if (!selectedFile || !previewUrl) return;

    setPhase("loading");
    setSubmissionError("");

    const channelId = crypto.randomUUID();
    const capturedImageUrl = previewUrl;
    const ws = new WebSocket(buildWsUrl(channelId));
    let received = false;

    ws.onopen = async () => {
      try {
        const formData = new FormData();
        formData.append("method_id", String(activeMethod.apiId));
        formData.append("channel_id", channelId);
        formData.append("image", selectedFile!);

        const resp = await fetch("/api/v1/evaluate", {
          method: "POST",
          body: formData,
        });
        if (!resp.ok) {
          const body = (await resp.json().catch(() => ({}))) as Record<
            string,
            string
          >;
          throw new Error(body["error"] ?? `HTTP ${resp.status}`);
        }
      } catch (err) {
        ws.close();
        setSubmissionError(
          err instanceof Error ? err.message : "Ошибка загрузки",
        );
        setPhase("error");
      }
    };

    ws.onmessage = (event: MessageEvent<string>) => {
      received = true;
      try {
        const msg = JSON.parse(event.data) as WsResultMessage;
        setResult({
          imageUrl: capturedImageUrl,
          score: typeof msg.score === "number" ? msg.score : null,
          status: msg.status ?? "ok",
        });
        setPhase("result");
      } catch {
        setSubmissionError("Не удалось разобрать ответ сервера");
        setPhase("error");
      }
    };

    ws.onerror = () => {
      if (!received) {
        setSubmissionError("Ошибка WebSocket соединения");
        setPhase("error");
      }
    };

    ws.onclose = () => {
      if (!received) {
        setSubmissionError("Соединение прервано без результата");
        setPhase("error");
      }
    };
  }

  return (
    <main className="min-h-screen bg-[linear-gradient(180deg,_#f8fbff_0%,_#eff6ff_100%)] text-slate-900">
      <section className="mx-auto max-w-6xl px-4 py-6 sm:px-6 md:px-8 md:py-8 lg:py-10">
        <EvaluationHeader />

        <div className="grid gap-6 lg:grid-cols-[0.9fr_1.1fr]">
          <EvaluationSidebar activeMethod={activeMethod} />
          {phase === "result" && result ? (
            <EvaluationResultCard />
          ) : (
            <EvaluationForm
              error={error}
              inputKey={inputKey}
              isLoading={phase === "loading"}
              previewUrl={previewUrl}
              selectedFile={selectedFile}
              selectedMethod={selectedMethod}
              submissionError={submissionError}
              onFileChange={handleFileChange}
              onMethodChange={setSelectedMethod}
              onReset={handleReset}
              onSubmit={handleEvaluate}
              apiError=""
            />
          )}
        </div>
      </section>
    </main>
  );
}
