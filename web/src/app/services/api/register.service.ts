import { Injectable, signal } from '@angular/core';
import { FakeServicesModel } from './fake-services.service';

export type FakeUserRegistrationModel = {
  name: string;
  email: string;
  password: string;
};
@Injectable({
  providedIn: 'root',
})
export class RegisterService {
  public users = signal<FakeUserRegistrationModel[]>([]);

  public setInitialUsers(): void {
    this.users.set([
      {
        name: 'admin',
        email: 'admin@gmail.com',
        password: 'admin',
      },
      {
        name: 'user',
        email: 'user@gmail.com',
        password: 'user',
      },
    ]);

    window.localStorage.setItem('users', JSON.stringify(this.users()));
  }

  public getUsers(): FakeUserRegistrationModel[] {
    if (!window.localStorage.getItem('users')) {
      this.setInitialUsers();
    }

    this.users.set(
      JSON.parse(
        window.localStorage.getItem('users')!
      ) as FakeUserRegistrationModel[]
    );
    return this.users();
  }

  public createUser(user: FakeUserRegistrationModel): void {
    this.users.update(() => [...this.users(), user]);
    window.localStorage.setItem('users', JSON.stringify(this.users()));
  }
}
