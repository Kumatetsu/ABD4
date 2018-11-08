import { Component, OnInit, Input } from '@angular/core';
import * as data from 'template/data.json';

@Component({
  selector: 'app-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.css']
})
export class LayoutComponent implements OnInit {
  background: string;
  constructor() { }

  ngOnInit() {
    this.background =  (<any>data).english.background_layout ;
  }

  getBg() {
    return 'url(' + this.background + ') no-repeat center';
  }
}
