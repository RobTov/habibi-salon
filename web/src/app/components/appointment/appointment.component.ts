import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-appointment',
  imports: [],
  templateUrl: './appointment.component.html',
  styleUrl: './appointment.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppointmentComponent {
  private toastr = inject(ToastrService);
  private router = inject(Router);

  public reserve(): void {
    this.toastr.success('Se ha realizado la cita correctamente.');
    setTimeout(() => {
      this.router.navigate(['/home']);
    }, 300)
  }
}
