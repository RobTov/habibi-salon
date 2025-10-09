import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { NavbarComponent } from './components/navbar/navbar.component';
import { FooterComponent } from './components/footer/footer.component';
import { ServicesResource } from './services/api/services.resource';
import { ICreateServicesModel } from './models/services.model';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, NavbarComponent, FooterComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {
  private servicesResource = inject(ServicesResource);

  private ngOnInit(): void {
    // const testService: ICreateServicesModel = {
    //   name: 'Created from the frontend',
    //   description: 'Lorem ipsum dolor sit amet',
    //   price: 10.0,
    //   is_active: true,
    // };

    // this.servicesResource.create(testService);

    // this.servicesResource.get();
    // this.servicesResource.services.subscribe((services) => {
    //   console.log(services);
    // });
  }
}
