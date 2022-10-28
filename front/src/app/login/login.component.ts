import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ConnectionServiceService } from '../connection-service.service';
import Swal from 'sweetalert2'



@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent  {

  email: string;
  password: string;

  constructor(public conection: ConnectionServiceService,private router:Router) {}

  login() {
    //console.log(this.email);
    //console.log(this.password);
    const user = {email: this.email, pass: this.password};
    console.log(JSON.stringify(user));
    this.conection.login(JSON.stringify(user)).subscribe( data => {
      
      if (data['email'] ) {
        console.log(data);
        localStorage.setItem('id', data['id']);
        this.router.navigate(['/memsim']);
        //Swal.fire('ALV', '', 'success')
        mensajeExito("Succes","Bienvenido a MetaOS "+data['email']);
      }else{
        mensajeError("Error","Usuario o contraseña incorrectos");
      }
    });
  }

  register() {
    this.router.navigate(['/register']);
  }
  /*async login(){
    const user = {email: this.email, pass: this.password};
    console.log(JSON.stringify(user));
    const response = await fetch('http://localhost:8080/login', {
      method : "POST",
      mode: 'no-cors',
      headers : {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        
      },
      body : JSON.stringify(user)
    });

    const data = await response.text();

    console.log(data);
  }*/
}

export function mensajeError(titulo:string, mensaje:string) : void{
  Swal.fire({
    position: 'top',
    icon: 'error',
    title: titulo,
    text : mensaje,
    showConfirmButton: false,
    timer: 2500
  })
}

export function mensajeExito(titulo:string, mensaje:string) :void {
  Swal.fire({
    position: 'top',
    icon: 'success',
    title: titulo,
    text : mensaje,
    showConfirmButton: false,
    timer: 2500
  })
}
