import { Component, OnInit } from '@angular/core';
import { TarifService } from '../services/tarif.service';
import { HttpClient } from '@angular/common/http';
import { NotifierService } from 'angular-notifier';

@Component({
  selector: 'app-tarif',
  templateUrl: './tarif.component.html',
  styleUrls: ['./tarif.component.css']
})

export class TarifComponent {
    tarifs: any;
    new_description: any;
    new_prix: any = "";
    new_tarif: any = "";
    error_msg: boolean = false;

  constructor(private TarifService: TarifService) {
  }

  ngOnInit() {
    this.getAllTarif();
  }

  addNewTarif(){
    if (this.new_description == "" || this.new_prix == "") {
      console.log("Champs requis");

      this.error_msg = true;
      
      return;
    }

    this.new_tarif = { 
      "description" : this.new_description,
      "prix" : +this.new_prix
    };

    this.TarifService.postTarif(this.new_tarif).subscribe(
      response => console.log(response),
      error => console.log("Error : ", error),
      () => {
        this.getAllTarif()
        this.new_description = ''
        this.new_prix = ""
      }
    )
  }

  getAllTarif() {
    let load_tarifs = true
    this.TarifService.getTarif().subscribe(
      response => this.tarifs = response.data,
      error => console.log(error),
      () => load_tarifs = false
    )
  }
}
