import { Injectable } from '@angular/core';
import { HttpClient , HttpHeaders } from "@angular/common/http";
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ConnectionServiceService {
  constructor(private http: HttpClient) {}

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  url = environment.url;

  login(user): Observable<any> {
    return this.http.post(this.url+"/login", user);
  }


  newMemsim(data): Observable<any> {
    return this.http.post(this.url+"/memsim",data);
  }

  newUser(user): Observable<any> {
    return this.http.post(this.url+"/logup",user);
  }

}
