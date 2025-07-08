import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, firstValueFrom } from 'rxjs';
import { RestService } from '../rest.service';
import { ICreateClientsModel, ClientsModel } from '../../models/clients.model';

@Injectable({
  providedIn: 'root',
})
export class ClientsResource {
  private readonly clients$: BehaviorSubject<ClientsModel[]> =
    new BehaviorSubject<ClientsModel[]>([]);
  public clients = this.clients$.asObservable();

  private readonly clientDetails$: BehaviorSubject<ClientsModel | null> =
    new BehaviorSubject<ClientsModel | null>(null);
  public clientDetails = this.clientDetails$.asObservable();

  private readonly restService: RestService = inject(RestService);

  async get() {
    const url = `clients/`;
    const clients = await firstValueFrom(
      this.restService.get<ClientsModel[]>(url)
    );
    this.clients$.next(clients);
  }

  async getDetails(clientID: string) {
    const service = await firstValueFrom(
      this.restService.get<ClientsModel>(`clients/${clientID}/`)
    );
    this.clientDetails$.next(service);
  }

  async create(clientData: ICreateClientsModel): Promise<void> {
    console.log('Creating client', clientData);

    const s = (await firstValueFrom(
      this.restService.post('clients/', {
        name: clientData.name,
        email: clientData.email,
        phone: clientData.phone,
        address: clientData.address,
      })
    )) as ClientsModel;
    const clients = this.clients$.value;
    this.clients$.next([...clients, s]);
  }

  async edit(request: ICreateClientsModel, clientID: number): Promise<void> {
    const response = await firstValueFrom(
      this.restService.put(`clients/${clientID}/`, {
        name: request.name,
        email: request.email,
        phone: request.phone,
        address: request.address,
      })
    );

    const currentClients = this.clients$.getValue();
    const updatedClient = currentClients.map((s) => {
      if (s.id === clientID) {
        console.log(request.name);
      }
      return s.id === clientID ? { ...s, request } : s;
    });

    console.log(updatedClient);
    this.clients$.next(updatedClient);
  }

  async delete(request: ClientsModel): Promise<void> {
    const response = await firstValueFrom(
      this.restService.delete(`clients/${request.id}/`)
    );

    const currentClients = this.clients$.getValue();
    const updatedClient = currentClients.filter((service) => {
      return service.id !== request.id;
    });

    console.log(updatedClient);
    this.clients$.next(updatedClient);
  }
}
