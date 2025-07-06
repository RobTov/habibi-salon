import { Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { NavbarComponent } from "./components/navbar/navbar.component";
import { FooterComponent } from "./components/footer/footer.component";
import { ServicesResource } from './services/api/services.resource';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, NavbarComponent, FooterComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  private servicesResource = inject(ServicesResource);


  private ngOnInit(): void {
    this.servicesResource.get();
    this.servicesResource.services.subscribe(services => {
      console.log(services)
    })
  }
}
