import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import {
  FakeServicesModel,
  FakeServicesService,
} from '../../../services/api/fake-services.service';
import { RouterModule } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-services-list',
  imports: [RouterModule],
  templateUrl: './services-list.component.html',
  styleUrl: './services-list.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ServicesListComponent {
  private fakeServicesService = inject(FakeServicesService);
  public services: FakeServicesModel[] = [];
  private toastr = inject(ToastrService);

  ngOnInit(): void {
    this.services = this.fakeServicesService.getServices();
    console.log(this.services);
  }

  public delete(id: number): void {
    const confirm = window.confirm(
      'Está seguro que desea eliminar el servicio?'
    );
    if (confirm) {
      setTimeout(() => {
        this.services = this.services.filter((s) => s.id !== id);
        this.toastr.success('El servicio se ha eliminado correctamente.');
      }, 300);
    }

    return;
  }
}
