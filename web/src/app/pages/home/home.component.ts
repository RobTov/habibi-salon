import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { ServicesModel } from '../../models/services.model';
import { ServicesResource } from '../../services/api/services.resource';

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

  private servicesResource = inject(ServicesResource);

  ngOnInit(): void {
    this.servicesResource.get();
    this.servicesResource.services.subscribe(services => {
      this.services = services;
    })
    
    this.isLoading = false;
  }
 }  
