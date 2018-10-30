import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';
import * as data from 'template/data.json';

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.css']
})
export class FooterComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  template = (<any>data).english;
  menu: any;
  company: string;
  logo: string;
  aboutus: string;
  phrase: string;
  addr: string;
  companymail: string;

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
    if (this.language === this.translationService.getState().en) {
      this.template = (<any>data).english;
    } else if (this.language === this.translationService.getState().ch) {
      this.template = (<any>data).chinese;
    }
    this.company = this.template.company;
    this.menu = this.template.header.menu;
    this.logo = this.template.header.logo;
    this.aboutus = this.template.whois.title;
    this.phrase = this.template.whois.phrases[0];
    this.companymail = this.template.companymail;
    this.addr = this.template.addr;
  }
}
