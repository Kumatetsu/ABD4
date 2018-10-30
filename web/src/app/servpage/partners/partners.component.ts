import { Component, OnInit, OnDestroy } from '@angular/core';
import * as data from 'template/data.json';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';

@Component({
  selector: 'app-partners',
  templateUrl: './partners.component.html',
  styleUrls: ['./partners.component.css']
})
export class PartnersComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  template = (<any>data).english;
  // partners are link to partners' pictures, it is setted once only
  partners = (<any>data).partners;
  // partnersintro is text, it change with language
  partnersintro: any;
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
    this.partnersintro = this.template.partnersintro;
  }
}
