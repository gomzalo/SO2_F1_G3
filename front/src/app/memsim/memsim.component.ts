import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ConnectionServiceService } from '../connection-service.service';
import Swal from 'sweetalert2'

@Component({
  selector: 'app-memsim',
  templateUrl: './memsim.component.html',
  styleUrls: ['./memsim.component.css']
})
export class MemsimComponent  {

  constructor(public conection: ConnectionServiceService,private router:Router) { }
  ciclos: string;
  unidades: string;
  memsim:any;
  id: string;

  newMemsim() {
    this.id=localStorage.getItem('id');
    const unidadesAux = '['+this.unidades+']';
    
    //const data = {ciclos: this.ciclos, unidades: this.unidades};
    const data = "{\"user_id\":\""+this.id+"\",\"ciclos\":"+this.ciclos+",\"unidades\":"+unidadesAux+"}";
    console.log(data);
    this.conection.newMemsim(data).subscribe( re => {
      //console.log(re);
      this.memsim=re;
      mensajeExito("PROCESADO","Validar en la parte inferior");
      //console.log("-------------------------");
      //console.log(this.memsim['memsim']);
    });
  }

  salir(){

    this.router.navigate(['/login']);
    localStorage.clear();
    mensajeExito("Success","Sesion cerrada");
      
  }



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