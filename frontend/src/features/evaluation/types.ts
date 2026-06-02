import { methods } from "./constants";

export type EvaluationMethod = (typeof methods)[number];
export type EvaluationMethodId = EvaluationMethod["id"];

export type EvaluateResponse = {
  status: string;
  id?: string;
};

export type WsResultMessage = {
  id: string;
  score?: number;
  status?: string;
  attn_img?: string;
};

export type EvaluationResult = {
  imageUrl: string;
  attentionImageUrl?: string;
  score: number | null;
  status: string;
};
