export interface TaskSubmission {
  note: string;
  link: string;
  submittedAt: string;
}

export interface TaskEvent {
  action: string;
  actor: string;
  detail: string;
  at: string;
}

export interface PracticeTask {
  id: number;
  title: string;
  requirement: string;
  status: string;
  submission?: TaskSubmission;
  rejectReason?: string;
  history: TaskEvent[];
}

export interface CoachingSession {
  id: number;
  mentor: string;
  learner: string;
  skill: string;
  time: string;
  place: string;
  status: string;
  tasks: PracticeTask[];
  closedAt?: string;
}

export interface SkillWallEntry {
  id: number;
  user: string;
  kind: string;
  skill: string;
  title: string;
  detail: string;
  sessionId: number;
  createdAt: string;
}

export type Perspective = 'mentor' | 'learner';

export interface TodoItem {
  key: string;
  text: string;
  sessionId: number;
}
