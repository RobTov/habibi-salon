import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, firstValueFrom } from 'rxjs';
import { RestService } from '../rest.service';
import {
  ICreateServicesModel,
  ServicesModel,
} from '../../models/services.model';

@Injectable({
  providedIn: 'root',
})
export class ServicesResource {
  private readonly services$: BehaviorSubject<ServicesModel[]> =
    new BehaviorSubject<ServicesModel[]>([]);
  public services = this.services$.asObservable();

  private readonly serviceDetails$: BehaviorSubject<ServicesModel | null> =
    new BehaviorSubject<ServicesModel | null>(null);
  public serviceDetails = this.serviceDetails$.asObservable();

  private readonly restService: RestService = inject(RestService);

  async get() {
    const url = `services/`;
    const services = await firstValueFrom(
      this.restService.get<ServicesModel[]>(url)
    );
    this.services$.next(services);
  }

  async getDetails(serviceID: string) {
    const service = await firstValueFrom(
      this.restService.get<ServicesModel>(`services/${serviceID}/`)
    );
    this.serviceDetails$.next(service);
  }

  async create(serviceData: ICreateServicesModel): Promise<void> {
    console.log('Creating service', serviceData);

    const s = (await firstValueFrom(
      this.restService.post('services/', {
        name: serviceData.name,
        description: serviceData.description,
        price: serviceData.price,
        is_active: serviceData.is_active,
      })
    )) as ServicesModel;
    const services = this.services$.value;
    this.services$.next([...services, s]);
  }

  async edit(request: ICreateServicesModel, serviceID: number): Promise<void> {
    const response = await firstValueFrom(
      this.restService.put(`services/${serviceID}/`, {
        name: request.name,
        description: request.description,
        price: request.price,
        is_active: request.is_active,
      })
    );

    const currentServices = this.services$.getValue();
    const updatedService = currentServices.map((s) => {
      if (s.id === serviceID) {
        console.log(request.name);
      }
      return s.id === serviceID ? { ...s, request } : s;
    });

    console.log(updatedService);
    this.services$.next(updatedService);
  }

  async delete(request: ServicesModel): Promise<void> {
    const response = await firstValueFrom(
      this.restService.delete(`services/${request.id}/`)
    );

    const currentServices = this.services$.getValue();
    const updatedService = currentServices.filter((service) => {
      return service.id !== request.id;
    });

    console.log(updatedService);
    this.services$.next(updatedService);
  }
}
