import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { LoginService } from '../../services/api/login.service';
import { Router, RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
@Component({
  selector: 'app-login',
  imports: [RouterModule, FormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginComponent {
  public loginService = inject(LoginService);
  private router = inject(Router);
  private toastr = inject(ToastrService)

  public email: string = '';

  public login(): void {
    if (this.email.includes('admin')) {
      this.loginService.setRole('admin');
    } else {
      this.loginService.setRole('user');
    }

    this.loginService.setLogged(true);

    this.router.navigate(['/home']);
    this.toastr.success('Inicio de sesión correcto', 'Bienvenido!')
  }
}
