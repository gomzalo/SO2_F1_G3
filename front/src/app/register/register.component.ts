import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ConnectionServiceService } from '../connection-service.service';
import Swal from 'sweetalert2'

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent  {

  email: string;
  password: string;
  confirmPassword: string;

  constructor(public conection: ConnectionServiceService,private router:Router) {}

  register() {
    if (this.password == this.confirmPassword) {
      const data = {email: this.email, pass: this.password};
      console.log(JSON.stringify(data));
      this.conection.newUser(JSON.stringify(data)).subscribe( data => {
        console.log(data);
        if (data['message']== 'Cuenta creada satisfactoriamente') {
          this.router.navigate(['/login']);
          mensajeExito("Succes","Cuenta creada satisfactoriamente");
        }else{
          mensajeError("Error",data['message']);
        }
      })
    }
  }

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