import { API_BASE } from '../constants/app.constants';
import type { CoachingSession, SkillWallEntry } from '../types/coaching';

interface BusinessErrorBody {
  code?: string;
  message?: string;
}

async function request<T>(path: string, options?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE}${path}`, {
    headers: { 'Content-Type': 'application/json' },
    ...options,
  });
  if (!response.ok) {
    let message = '请求失败，请稍后重试';
    try {
      const body = (await response.json()) as BusinessErrorBody;
      if (body.message) message = body.message;
    } catch {
      // 保留默认错误提示
    }
    throw new Error(message);
  }
  return response.json() as Promise<T>;
}

export function fetchCoachingSessions(): Promise<CoachingSession[]> {
  return request<CoachingSession[]>('/coaching/sessions');
}

export function fetchSkillWall(user: string): Promise<SkillWallEntry[]> {
  return request<SkillWallEntry[]>(`/coaching/skillwall?user=${encodeURIComponent(user)}`);
}

export function addCoachingTask(sessionId: number, payload: { mentor: string; title: string; requirement: string }) {
  return request<CoachingSession>(`/coaching/sessions/${sessionId}/tasks`, {
    method: 'POST',
    body: JSON.stringify(payload),
  });
}

export function submitCoachingTask(sessionId: number, taskId: number, payload: { learner: string; note: string; link: string }) {
  return request<CoachingSession>(`/coaching/sessions/${sessionId}/tasks/${taskId}/submit`, {
    method: 'POST',
    body: JSON.stringify(payload),
  });
}

export function reviewCoachingTask(sessionId: number, taskId: number, payload: { mentor: string; action: string; reason: string }) {
  return request<CoachingSession>(`/coaching/sessions/${sessionId}/tasks/${taskId}/review`, {
    method: 'POST',
    body: JSON.stringify(payload),
  });
}
