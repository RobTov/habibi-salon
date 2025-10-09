import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import {
  FakeServicesModel,
  FakeServicesService,
} from '../../../services/api/fake-services.service';
import { RouterModule } from '@angular/router';

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

  ngOnInit(): void {
    this.services = this.fakeServicesService.getServices();
    console.log(this.services);
  }

  public delete(id: number): void {}
}
