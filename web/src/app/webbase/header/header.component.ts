import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';
import * as data from 'template/data.json';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  template = (<any>data).english.header;
  logo: string;
  translateLink: string;
  menu: any;
  constructor(private translationService: TranslationService) {
    this.subscription = this.translationService.getLanguage().subscribe(language => {
      this.language = language.text;
      this.defineTemplate();
    });
  }
  ngOnInit() {
    this.language = this.translationService.getState().actual;
    this.defineTemplate();
  }
  ngOnDestroy() {
    this.subscription.unsubscribe();
  }
  defineTemplate() {
    if (this.language === this.translationService.getState().ch) {
      this.template = (<any>data).chinese.header;
    } else if (this.language === this.translationService.getState().en) {
      this.template = (<any>data).english.header;
    }
    this.logo = this.template.logo;
    this.translateLink = this.template.translateLink;
    this.menu = this.template.menu;
    console.log('logo: ', this.logo);
  }
  // switch language
  translate() {
    this.translationService.sendLanguage(this.translateLink);
  }
}
