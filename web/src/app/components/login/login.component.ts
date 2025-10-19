import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { LoginService } from '../../services/api/login.service';
import { Router, RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
import {
  FakeUserRegistrationModel,
  RegisterService,
} from '../../services/api/register.service';
@Component({
  selector: 'app-login',
  imports: [RouterModule, FormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginComponent {
  public loginService = inject(LoginService);
  private usersService = inject(RegisterService);
  private router = inject(Router);
  private toastr = inject(ToastrService);

  private users: FakeUserRegistrationModel[] = [];

  ngOnInit(): void {
    this.users = this.usersService.getUsers();
  }

  public email: string = '';
  public password: string = '';

  public login(): void {
    const currentUser = this.users.filter((u) => {
      return u.email === this.email;
    })[0];

    if (!currentUser) {
      this.toastr.error('Email Incorrecto', 'Error');
      return;
    }

    if (currentUser.password !== this.password) {
      this.toastr.error('Contraseña Incorrecta', 'Error');
      return;
    }

    if (this.email.includes('admin')) {
      this.loginService.setRole('admin');
    } else {
      this.loginService.setRole('user');
    }

    this.loginService.setLogged(true);

    this.router.navigate(['/home']);
    this.toastr.success('Inicio de sesión correcto', 'Bienvenido!');
  }
}
