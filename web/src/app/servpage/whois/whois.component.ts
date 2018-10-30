import { Component, OnInit, OnDestroy } from '@angular/core';
import * as data from 'template/data.json';
import { Subscription } from 'rxjs';
import { TranslationService } from '../../services/translation.service';

@Component({
  selector: 'app-whois',
  templateUrl: './whois.component.html',
  styleUrls: ['./whois.component.css']
})
export class WhoisComponent implements OnInit, OnDestroy {
  subscription: Subscription;
  language: string;
  template = (<any>data).english.whois;
  title: string;
  image: string;
  phrases: any;
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
      this.template = (<any>data).english.whois;
    } else if (this.language === this.translationService.getState().ch) {
      this.template = (<any>data).chinese.whois;
    }
    this.title = this.template.title;
    this.phrases = this.template.phrases;
    this.image = this.template.image;
  }
}
