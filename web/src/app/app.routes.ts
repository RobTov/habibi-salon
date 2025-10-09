import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'home',
    pathMatch: 'full',
  },
  {
    path: 'home',
    loadComponent: () =>
      import('./pages/home/home.component').then((mod) => mod.HomeComponent),
  },
  {
    path: 'login',
    loadComponent: () =>
      import('./components/login/login.component').then(
        (mod) => mod.LoginComponent
      ),
  },
  {
    path: 'appointment',
    loadComponent: () =>
      import('./components/appointment/appointment.component').then(
        (mod) => mod.AppointmentComponent
      ),
  },
  {
    path: 'appointments',
    loadComponent: () =>
      import('./pages/appointments/appointments.component').then(
        (mod) => mod.AppointmentsComponent
      ),
  },
  {
    path: 'services',
    children: [
      {
        path: '',
        loadComponent: () =>
          import(
            './pages/services-page/services-list/services-list.component'
          ).then((mod) => mod.ServicesListComponent),
      },
      {
        path: 'new',
        loadComponent: () =>
          import(
            './pages/services-page/services-create/services-create.component'
          ).then((mod) => mod.ServicesCreateComponent),
      },
      {
        path: 'edit/:id',
        loadComponent: () =>
          import(
            './pages/services-page/services-edit/services-edit.component'
          ).then((mod) => mod.ServicesEditComponent),
      },
    ],
  },
];
