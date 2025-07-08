import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, firstValueFrom } from 'rxjs';
import { RestService } from '../rest.service';
import {
  AppointmentsModel,
  ICreateAppointmentsModel,
} from '../../models/appointments.model';

@Injectable({
  providedIn: 'root',
})
export class AppointmentsResource {
  private readonly appointments$: BehaviorSubject<AppointmentsModel[]> =
    new BehaviorSubject<AppointmentsModel[]>([]);
  public appointments = this.appointments$.asObservable();

  private readonly appointmentDetails$: BehaviorSubject<AppointmentsModel | null> =
    new BehaviorSubject<AppointmentsModel | null>(null);
  public appointmentDetails = this.appointmentDetails$.asObservable();

  private readonly restService: RestService = inject(RestService);

  async get() {
    const url = `appointments/`;
    const appointments = await firstValueFrom(
      this.restService.get<AppointmentsModel[]>(url)
    );
    this.appointments$.next(appointments);
  }

  async getDetails(appointmentID: string) {
    const service = await firstValueFrom(
      this.restService.get<AppointmentsModel>(`appointments/${appointmentID}/`)
    );
    this.appointmentDetails$.next(service);
  }

  async create(appointmentData: ICreateAppointmentsModel): Promise<void> {
    console.log('Creating appointment', appointmentData);

    const s = (await firstValueFrom(
      this.restService.post('appointments/', {
        date: appointmentData.date,
        service_id: appointmentData.service_id,
        status: appointmentData.status,
      })
    )) as AppointmentsModel;
    const appointments = this.appointments$.value;
    this.appointments$.next([...appointments, s]);
  }

  async edit(
    request: ICreateAppointmentsModel,
    appointmentID: number
  ): Promise<void> {
    const response = await firstValueFrom(
      this.restService.put(`appointments/${appointmentID}/`, {
        date: request.date,
        service_id: request.service_id,
        status: request.status,
      })
    );

    const currentAppointments = this.appointments$.getValue();
    const updatedAppointment = currentAppointments.map((s) => {
      if (s.id === appointmentID) {
        console.log(`${request.date} - date:${request.service_id}`);
      }
      return s.id === appointmentID ? { ...s, request } : s;
    });

    console.log(updatedAppointment);
    this.appointments$.next(updatedAppointment);
  }

  async delete(request: AppointmentsModel): Promise<void> {
    const response = await firstValueFrom(
      this.restService.delete(`appointments/${request.id}/`)
    );

    const currentAppointments = this.appointments$.getValue();
    const updatedAppointment = currentAppointments.filter((service) => {
      return service.id !== request.id;
    });

    console.log(updatedAppointment);
    this.appointments$.next(updatedAppointment);
  }
}
