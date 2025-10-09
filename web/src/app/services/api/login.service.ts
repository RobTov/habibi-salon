import { Injectable, signal } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class LoginService {
  public setRole(role: string): void {
    window.localStorage.setItem('role', role);
  }

  public setLogged(isLogged: boolean): void {
    window.localStorage.setItem('isLogged', String(isLogged));
  }

  public getRole(): string {
    return window.localStorage.getItem('role') || '';
  }

  public getIsLogged(): boolean {
    return window.localStorage.getItem('isLogged') === 'true';
  }
}
