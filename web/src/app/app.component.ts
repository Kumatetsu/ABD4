import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import * as data from 'template/data.json';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit, OnDestroy {
  background: string;

  constructor() {

  }

  ngOnInit() {

  }

  ngOnDestroy() {

  }
  
  getBg() {
    return 'url(' + this.background + ') no-repeat center';
  }
}
