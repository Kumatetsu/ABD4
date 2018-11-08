import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import * as data from 'template/data.json';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit, OnDestroy {
  template = (<any>data).english.header;
  menu: any;

  constructor() {
  }
  ngOnInit() {
    this.menu = this.template.menu;
  }
  ngOnDestroy() {
  }
}
