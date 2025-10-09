import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import {
  FakeServicesModel,
  FakeServicesService,
} from '../../../services/api/fake-services.service';
import { ToastrService } from 'ngx-toastr';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-services-create',
  imports: [FormsModule],
  templateUrl: './services-create.component.html',
  styleUrl: './services-create.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ServicesCreateComponent {
  private router = inject(Router);
  private toastr = inject(ToastrService);
  private fakeServicesService = inject(FakeServicesService);

  public newService: FakeServicesModel = {
    id: 9,
    image: '',
    name: '',
    price: '',
  };

  public create(): void {
    this.newService.image = '/img/12.jpeg';

    console.log(this.newService);
    setTimeout(() => {
      this.fakeServicesService.createService(this.newService);
      this.toastr.success('El servicio ha sido creado con éxito.');
      this.router.navigate(['services']);
    }, 300);
  }
}
