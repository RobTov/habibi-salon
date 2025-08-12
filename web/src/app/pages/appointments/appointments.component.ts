import {
  ChangeDetectionStrategy,
  Component,
  inject,
  signal,
} from '@angular/core';
import { AppointmentsModel } from '../../models/appointments.model';
import { AppointmentsResource } from '../../services/api/appointments.resource';

@Component({
  selector: 'app-appointments',
  imports: [],
  templateUrl: './appointments.component.html',
  styleUrl: './appointments.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppointmentsComponent {
  public fakeAppointments = [
    {
      name: 'Mariana López',
      email: 'mariana.lopez@email.com',
      phone: '55991213',
      status: 'active',
    },
    {
      name: 'Valeria Gómez',
      email: 'valeria.gomez@email.com',
      phone: '58215422',
      status: 'active',
    },
    {
      name: 'Camila Rodríguez',
      email: 'camila.rodriguez@email.com',
      phone: '51179877',
      status: 'active',
    },
    {
      name: 'Sofía Fernández',
      phone: '55190023',
      email: 'sofia.fernandez@email.com',
      status: 'inactive',
    },
    {
      name: 'Isabella Ramírez',
      email: 'isabella.ramirez@email.com',
      phone: '53905005',
      status: 'active',
    },
    {
      name: 'Gabriela Torres',
      email: 'gabriela.torres@email.com',
      phone: '58990167',
      status: 'inactive',
    },
    {
      name: 'Luciana Vargas',
      email: 'luciana.vargas@email.com',
      phone: '54228899',
      status: 'inactive',
    },
    {
      name: 'Daniela Castillo',
      email: 'daniela.castillo@email.com',
      phone: '52152963',
      status: 'active',
    },
    {
      name: 'Alejandra Pérez',
      email: 'alejandra.perez@email.com',
      phone: '55059128',
      status: 'active',
    },
    {
      name: 'Natalia Sánchez',
      email: 'natalia.sanchez@email.com',
      phone: '59993367',
      status: 'inactive',
    },
  ];

  public appointments: AppointmentsModel[] = [];
  public isLoading = signal<boolean>(true);
  private appointmentsResource = inject(AppointmentsResource);

  ngOnInit(): void {
    this.appointmentsResource.get();
    this.appointmentsResource.appointments.subscribe((appointments) => {
      this.appointments = appointments;
      this.isLoading.set(false);
    });
  }
}
