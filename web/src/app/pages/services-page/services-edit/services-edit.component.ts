import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import {
  FakeServicesService,
  FakeServicesModel,
} from '../../../services/api/fake-services.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-services-edit',
  imports: [FormsModule],
  templateUrl: './services-edit.component.html',
  styleUrl: './services-edit.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ServicesEditComponent {
  private router = inject(Router);
  private activatedRoute = inject(ActivatedRoute);
  private toastr = inject(ToastrService);
  private fakeServicesService = inject(FakeServicesService);
  private serviceID!: number;

  public serviceToEdit: FakeServicesModel = {
    id: 9,
    image: '',
    name: '',
    price: '',
  };

  constructor() {
    this.activatedRoute.params.subscribe((params) => {
      this.serviceID = Number(params['id']);
    });
  }

  ngOnInit(): void {
    const services = this.fakeServicesService.getServices();
    this.serviceToEdit = services.filter((service) => {
      return service.id === this.serviceID;
    })[0];
    console.log(this.serviceToEdit);
  }

  public edit(): void {
    this.serviceToEdit.image = '/img/12.jpeg';

    setTimeout(() => {
      this.fakeServicesService.editService(this.serviceToEdit);
      this.toastr.success('El servicio ha sido actualizado con éxito.');
      this.router.navigate(['services']);
    }, 300);
  }
}
