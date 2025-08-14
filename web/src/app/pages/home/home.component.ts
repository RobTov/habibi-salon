import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { ServicesModel } from '../../models/services.model';
import { ServicesResource } from '../../services/api/services.resource';
import { ClientsModel } from '../../models/clients.model';
import { ClientsResource } from '../../services/api/clients.resource';

@Component({
  selector: 'app-home',
  imports: [],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HomeComponent {
  public isLoading: boolean = true;
  public services: ServicesModel[] = [];
  public clients: ClientsModel[] = [];

  private servicesResource = inject(ServicesResource);
  private clientsResource = inject(ClientsResource);

  ngOnInit(): void {
    this.servicesResource.get();
    this.servicesResource.services.subscribe((services) => {
      this.services = services;
    });

    this.clientsResource.get();
    this.clientsResource.clients.subscribe((clients) => {
      this.clients = clients;
    });

    this.isLoading = false;
  }
}
