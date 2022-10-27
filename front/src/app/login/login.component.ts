import { Component, OnInit } from '@angular/core';
import { ConnectionServiceService } from '../connection-service.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent  {

  email: string;
  password: string;

  constructor(public conection: ConnectionServiceService) {}

  /*login() {
    //console.log(this.email);
    //console.log(this.password);
    const user = {email: this.email, pass: this.password};
    console.log(JSON.stringify(user));
    this.conection.login(JSON.stringify(user)).subscribe( data => {
      console.log(data);
    });
  }*/

  async login(){
    const user = '{email: '+this.email+', pass: '+this.password+'}';
    const response = await fetch('http://localhost:8080/login', {
      method : "POST",
      mode: 'no-cors',
      headers : {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        
      },
      body : JSON.parse(JSON.stringify(user))
    });

    const data = await response.json();

    console.log(data);
  }
}
