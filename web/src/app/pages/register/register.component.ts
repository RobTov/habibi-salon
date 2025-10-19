import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import {
  FakeUserRegistrationModel,
  RegisterService,
} from '../../services/api/register.service';
import { FakeServicesService } from '../../services/api/fake-services.service';

@Component({
  selector: 'app-register',
  imports: [RouterLink, FormsModule],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RegisterComponent {
  private toastrService = inject(ToastrService);
  private router = inject(Router);
  private userService = inject(RegisterService);

  public newUser: FakeUserRegistrationModel = {
    name: '',
    email: '',
    password: '',
  };

  public passwordConfirm = '';

  public registerUser(): void {
    if (this.newUser.password !== this.passwordConfirm) {
      this.toastrService.error('Las contraseñas no coinciden', 'Error');
      return;
    }

    this.userService.createUser(this.newUser);
    setTimeout(() => {
      this.toastrService.success('El usuario se ha creado con éxito!');
    }, 300);
    this.router.navigate(['/home']);
  }
}
