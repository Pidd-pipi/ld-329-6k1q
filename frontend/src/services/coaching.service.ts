import type { CoachingAppointment, CoachingRole, SkillWallEntry } from '../types/domain';

const API_BASE = '/api';

async function parseResponse<T>(response: Response, fallback: string): Promise<T> {
  if (!response.ok) {
    let message = fallback;
    try {
      const body = (await response.json()) as { message?: string };
      if (body.message) message = body.message;
    } catch {
      // 保留默认错误提示
    }
    throw new Error(message);
  }
  return response.json() as Promise<T>;
}

export async function fetchCoachingAppointments(user: string, role: CoachingRole): Promise<CoachingAppointment[]> {
  const response = await fetch(`${API_BASE}/coaching/appointments?user=${encodeURIComponent(user)}&role=${role}`);
  return parseResponse<CoachingAppointment[]>(response, '无法加载阶段陪练预约');
}

export async function addCoachingTask(
  appointmentId: number,
  payload: { user: string; title: string; requirement: string },
): Promise<CoachingAppointment> {
  const response = await fetch(`${API_BASE}/coaching/appointments/${appointmentId}/tasks`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
  return parseResponse<CoachingAppointment>(response, '追加任务失败');
}

export async function submitCoachingTask(
  taskId: number,
  payload: { user: string; note: string; link: string },
): Promise<CoachingAppointment> {
  const response = await fetch(`${API_BASE}/coaching/tasks/${taskId}/submit`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
  return parseResponse<CoachingAppointment>(response, '提交失败');
}

export async function reviewCoachingTask(
  taskId: number,
  payload: { user: string; approve: boolean; reason: string },
): Promise<CoachingAppointment> {
  const response = await fetch(`${API_BASE}/coaching/tasks/${taskId}/review`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  });
  return parseResponse<CoachingAppointment>(response, '处理失败');
}

export async function fetchSkillWall(user: string): Promise<SkillWallEntry[]> {
  const response = await fetch(`${API_BASE}/coaching/skillwall?user=${encodeURIComponent(user)}`);
  return parseResponse<SkillWallEntry[]>(response, '无法加载技能墙');
}
