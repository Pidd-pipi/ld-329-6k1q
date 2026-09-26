export interface Skill {
  id: number;
  owner: string;
  title: string;
  category: string;
  level: number;
  campus: string;
  description: string;
  timeSlots: string[];
  rewards: string[];
  portfolio: string;
}

export interface Need {
  id: number;
  requester: string;
  title: string;
  category: string;
  campus: string;
  expectTime: string;
  budgetType: string;
  description: string;
  responses: number;
}

export interface Match {
  id: number;
  provider: string;
  learner: string;
  offerSkill: string;
  wantedSkill: string;
  score: number;
  commonSlots: string[];
  recommendation: string;
}

export interface Appointment {
  id: number;
  pair: string;
  time: string;
  place: string;
  status: string;
  agenda: string;
}

export interface Review {
  id: number;
  from: string;
  to: string;
  rating: number;
  content: string;
}

export interface Conversation {
  id: number;
  withUser: string;
  unread: number;
  messages: string[];
}

export interface Profile {
  name: string;
  major: string;
  creditScore: number;
  creditLevel: string;
  skillWall: Skill[];
  radar: Record<string, number>;
  history: string[];
  reviews: Review[];
}

export interface Overview {
  service: string;
  categories: string[];
  metrics: Record<string, number>;
  skills: Skill[];
  needs: Need[];
  matches: Match[];
  appointments: Appointment[];
  reviews: Review[];
  messages: Conversation[];
  profile: Profile;
}

export type CoachingRole = 'mentor' | 'learner';

export type CoachingTaskStatus = 'todo' | 'submitted' | 'approved' | 'rejected';

export interface TaskSubmission {
  note: string;
  link: string;
  submittedAt: string;
}

export interface CoachingTask {
  id: number;
  appointmentId: number;
  title: string;
  requirement: string;
  status: CoachingTaskStatus;
  submission?: TaskSubmission;
  rejectReason?: string;
}

export interface CoachingAppointment {
  id: number;
  mentor: string;
  learner: string;
  skill: string;
  time: string;
  place: string;
  closed: boolean;
  closedAt?: string;
  tasks: CoachingTask[];
}

export interface SkillWallEntry {
  user: string;
  skill: string;
  partner: string;
  note: string;
  closedAt: string;
}
