export enum AppointmentStatus {
  'pending',
  'confirmed',
  'cancelled',
}

export interface AppointmentsModel {
  id: number;
  date: string;
  service_id: number;
  status: AppointmentStatus;
}
